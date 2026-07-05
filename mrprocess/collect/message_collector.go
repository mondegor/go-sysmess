package collect

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mondegor/go-core/mrprocess"
)

const (
	// defaultCaption - название сервиса по умолчанию.
	defaultCaption = "MessageCollector"

	// defaultReadyTimeout - таймаут готовности сервиса по умолчанию.
	defaultReadyTimeout = 30 * time.Second

	// defaultHandlerTimeout - таймаут выполнения обработчика пакета.
	defaultHandlerTimeout = 30 * time.Second

	// defaultBatchSize - размер пакета по умолчанию.
	defaultBatchSize = 100

	// defaultWorkersCount - количество воркеров-обработчиков по умолчанию.
	defaultWorkersCount = 1

	// keyWorkerID - название ключа ID воркера, добавляемого в контекст.
	keyWorkerID = "worker_id"

	// keyTaskID - название ключа ID задачи, добавляемого в контекст.
	keyTaskID = "task_id"
)

// defaultFlushPeriodStrategy - период принудительной отправки накопленного пакета.
var defaultFlushPeriodStrategy = mrprocess.NewStaticPeriodStrategy(60 * time.Second) //nolint:gochecknoglobals

type (
	// MessageCollector - многопоточный сервис сбора и пакетной обработки сообщений (PUSH-модель).
	//
	// Принцип работы:
	//  1. Внешний код отправляет сообщения через PushMessage();
	//  2. Сообщения накапливаются во внутренней очереди;
	//  3. При достижении batchSize или flushPeriodStrategy пакет отправляется обработчику;
	//  4. Обработка выполняется в отдельных воркерах;
	//
	// Тип T - тип обрабатываемых сообщений.
	MessageCollector[T any] struct {
		caption             string
		readyTimeout        time.Duration
		flushPeriodStrategy mrprocess.PeriodStrategy
		handlerTimeout      time.Duration
		batchSize           int
		workersCount        int

		handler      mrprocess.MessageBatchHandler[T]
		errorHandler errorHandler
		logger       logger
		traceManager traceManager

		wg            sync.WaitGroup
		isSendStopped atomic.Bool
		closeOnce     sync.Once
		messageQueue  chan T
		workersQueue  chan func(ctx context.Context)
		done          chan struct{}
	}

	errorHandler interface {
		Handle(ctx context.Context, err error)
	}

	logger interface {
		Debug(ctx context.Context, msg string, args ...any)
		Info(ctx context.Context, msg string, args ...any)
	}

	traceManager interface {
		WithGeneratedProcessID(ctx context.Context, key string) context.Context
	}
)

var (
	// ErrInternalCollectorWorkersAreStopped - воркеры коллектора сообщений остановлены (attrs: collector_name).
	ErrInternalCollectorWorkersAreStopped = errors.New("the message collector workers has been stopped")

	// ErrInternalMessageReceptionStopped - приём сообщений коллектором остановлен (attrs: collector_name).
	ErrInternalMessageReceptionStopped = errors.New("message reception in the message collector has been stopped")
)

// NewMessageCollector - создаёт сервис пакетной обработки сообщений (PUSH-модель).
func NewMessageCollector[T any](
	handler mrprocess.MessageBatchHandler[T],
	errorHandler errorHandler,
	logger logger,
	traceManager traceManager,
	opts ...Option[T],
) *MessageCollector[T] {
	o := options[T]{
		collector: &MessageCollector[T]{
			caption:             defaultCaption,
			readyTimeout:        defaultReadyTimeout,
			flushPeriodStrategy: defaultFlushPeriodStrategy,
			handlerTimeout:      defaultHandlerTimeout,

			handler:      handler,
			errorHandler: errorHandler,
			logger:       logger,
			traceManager: traceManager,

			wg:           sync.WaitGroup{},
			closeOnce:    sync.Once{},
			messageQueue: make(chan T),
			workersQueue: make(chan func(ctx context.Context)),
			done:         make(chan struct{}),
		},
	}

	for _, opt := range opts {
		opt(&o)
	}

	if o.captionPrefix != "" {
		o.collector.caption = o.captionPrefix + o.collector.caption
	}

	if o.collector.batchSize < 1 {
		o.collector.batchSize = defaultBatchSize
	}

	if o.collector.workersCount < 1 {
		o.collector.workersCount = defaultWorkersCount
	}

	if o.collector.handlerTimeout <= 0 {
		o.collector.handlerTimeout = defaultHandlerTimeout
	}

	return o.collector
}

// Caption - возвращает название сервиса обработки сообщений в свободной форме.
func (p *MessageCollector[T]) Caption() string {
	return p.caption
}

// ReadyTimeout - возвращает максимальное время, за которое должен быть запущен сервис.
func (p *MessageCollector[T]) ReadyTimeout() time.Duration {
	return p.readyTimeout
}

// Start - запуск сервиса обработки сообщений.
//
// Процесс работы:
//  1. Запускает N воркеров для обработки сообщений;
//  2. Накопляет сообщения из messageQueue до batchSize;
//  3. Отправляет пакет в workersQueue для обработки;
//  4. flushPeriodStrategy отвечает за период отправки накопленных сообщений (не достигших batchSize);
//  5. При отмене контекста очищает очередь и завершается;
//
// Важно:
//   - Отмена внешнего контекста приведёт к аварийному завершению (очистка очереди);
//   - Для корректной остановки используйте Shutdown;
//   - Повторный запуск того же объекта не поддерживается.
func (p *MessageCollector[T]) Start(ctx context.Context, ready func()) error {
	p.wg.Add(1)
	defer p.wg.Done()

	p.logger.Debug(ctx, "Starting the message collector...", "collector_name", p.caption)
	defer p.logger.Debug(ctx, "The message collector has been stopped")

	wgWorkers := sync.WaitGroup{}
	workersStopped := make(chan struct{})
	ticker := time.NewTicker(p.flushPeriodStrategy.Period())

	p.startWorkers(ctx, &wgWorkers)

	go func() {
		wgWorkers.Wait()
		close(workersStopped)
	}()

	defer func() {
		ticker.Stop()
		close(p.workersQueue)
		<-workersStopped
	}()

	messageBatch := make([]T, 0, p.batchSize)

	if ready != nil {
		ready()
	}

	for {
		select {
		case <-p.done:
			for {
				// в этом месте приёма новых данных уже нет,
				// но в очереди ещё могут оставаться данные, которые нужно обработать
				select {
				case message := <-p.messageQueue:
					messageBatch = append(messageBatch, message)

					if len(messageBatch) < p.batchSize {
						continue
					}
				default:
				}

				break
			}
		case <-ctx.Done():
			p.logger.Debug(ctx, "The message collector detected context 'Done'", "error", ctx.Err())

			// предварительно завершается приём данных
			p.isSendStopped.Store(true)

			// контекст отменён, поэтому накопленный пакет и оставшиеся в очереди сообщения
			// обрабатываются синхронно с отсоединённым от отмены контекстом,
			// чтобы данные не были потеряны
			flushCtx := context.WithoutCancel(ctx)

			for {
				select {
				case message := <-p.messageQueue:
					messageBatch = append(messageBatch, message)

					if len(messageBatch) < p.batchSize {
						continue
					}
				default:
				}

				if len(messageBatch) > 0 {
					p.flushBatch(flushCtx, messageBatch)
					messageBatch = messageBatch[:0]

					continue
				}

				return nil
			}
		case message := <-p.messageQueue:
			messageBatch = append(messageBatch, message)

			if len(messageBatch) < p.batchSize {
				continue
			}
		case <-ticker.C:
			p.logger.Debug(ctx, "The message collector ticker.C")
		}

		ticker.Reset(p.flushPeriodStrategy.Period())

		if len(messageBatch) == 0 {
			if p.isSendStopped.Load() {
				return nil // если данных нет и их приём остановлен, то процесс завершается
			}

			continue
		}

		p.logger.Info(ctx, "Got message batch in the message collector...", "message_batch", len(messageBatch))

		select {
		case <-workersStopped:
			return fmt.Errorf("%w [collector_name=%s]", ErrInternalCollectorWorkersAreStopped, p.caption)
		case p.workersQueue <- p.workerFunc(messageBatch):
			messageBatch = make([]T, 0, p.batchSize)
		}
	}
}

// PushMessage - отправляет сообщение в очередь для обработки.
// Блокируется до освобождения места в очереди или отмены контекста.
func (p *MessageCollector[T]) PushMessage(ctx context.Context, message T) error {
	if p.isSendStopped.Load() {
		return fmt.Errorf("%w [collector_name=%s]", ErrInternalMessageReceptionStopped, p.caption)
	}

	select {
	case p.messageQueue <- message:
		return nil
	case <-p.done:
		return fmt.Errorf("%w [collector_name=%s]", ErrInternalMessageReceptionStopped, p.caption)
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Shutdown - корректная остановка сервиса обработки сообщений.
// Останавливает приём новых сообщений и ожидает завершения всех операций.
func (p *MessageCollector[T]) Shutdown(ctx context.Context) error {
	p.logger.Debug(ctx, "Shutting down the message collector...")

	p.closeOnce.Do(func() {
		p.isSendStopped.Store(true) // завершается приём данных
		close(p.done)
	})

	p.wg.Wait()
	p.logger.Debug(ctx, "The message collector has been shut down")

	return nil
}

func (p *MessageCollector[T]) startWorkers(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i < p.workersCount; i++ {
		wg.Add(1)

		go func(ctx context.Context) {
			defer wg.Done()

			ctx = p.traceManager.WithGeneratedProcessID(ctx, keyWorkerID)

			for fn := range p.workersQueue {
				p.execWorkerFunc(ctx, fn)
			}

			p.logger.Debug(ctx, "The worker has been stopped")
		}(ctx)
	}
}

// execWorkerFunc - выполняет функцию обработки пакета с перехватом паники,
// чтобы паника при обработке одного пакета не завершала воркер.
func (p *MessageCollector[T]) execWorkerFunc(ctx context.Context, fn func(ctx context.Context)) {
	defer func() {
		if rvr := recover(); rvr != nil {
			p.errorHandler.Handle(
				ctx,
				mrprocess.NewCaughtPanicError("message collector: "+p.caption, rvr),
			)
		}
	}()

	fn(ctx)
}

func (p *MessageCollector[T]) workerFunc(messages []T) func(ctx context.Context) {
	return func(ctx context.Context) {
		p.flushBatch(ctx, messages)
	}
}

// flushBatch - синхронно передаёт пакет сообщений обработчику с таймаутом handlerTimeout.
func (p *MessageCollector[T]) flushBatch(ctx context.Context, messages []T) {
	handlerCtx, cancel := context.WithTimeout(p.traceManager.WithGeneratedProcessID(ctx, keyTaskID), p.handlerTimeout)
	defer cancel()

	p.logger.Debug(ctx, "Flushing message batch...", "message_batch", len(messages))

	if err := p.handler.Execute(handlerCtx, messages); err != nil {
		p.errorHandler.Handle(ctx, err)

		return
	}

	p.logger.Debug(ctx, "The handler has been successfully executed")
}
