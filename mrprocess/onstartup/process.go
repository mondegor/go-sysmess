package onstartup

import (
	"context"
	"sync"
	"time"

	"github.com/mondegor/go-sysmess/mrprocess"
)

const (
	// defaultCaption - название сервиса по умолчанию.
	defaultCaption = "OnStartup"

	// defaultReadyTimeout - таймаут готовности сервиса по умолчанию.
	defaultReadyTimeout = 30 * time.Second

	// keyTaskID - название ключа ID задачи, добавляемого в контекст.
	keyTaskID = "task_id"
)

type (
	// Process - сервис выполнения работы при старте приложения.
	//
	// Используется когда нужно выполнить работу после гарантированного запуска
	// остальных процессов (например: инициализация данных, миграции, прогрев кэша).
	//
	// Особенность: после выполнения job процесс ожидает сигнала завершения (done или ctx.Done).
	Process struct {
		caption      string
		readyTimeout time.Duration
		job          mrprocess.Job
		logger       logger
		traceManager traceManager
		wg           sync.WaitGroup
		closeOnce    sync.Once
		done         chan struct{}
	}

	logger interface {
		Debug(ctx context.Context, msg string, args ...any)
	}

	traceManager interface {
		WithGeneratedProcessID(ctx context.Context, key string) context.Context
	}
)

// NewProcess - создаёт сервис выполнения работы при старте.
func NewProcess(
	job mrprocess.Job,
	logger logger,
	traceManager traceManager,
	opts ...Option,
) *Process {
	o := options{
		process: &Process{
			caption:      defaultCaption,
			readyTimeout: defaultReadyTimeout,
			job:          job,
			logger:       logger,
			traceManager: traceManager,
			wg:           sync.WaitGroup{},
			closeOnce:    sync.Once{},
			done:         make(chan struct{}),
		},
	}

	for _, opt := range opts {
		opt(&o)
	}

	return o.process
}

// Caption - возвращает название сервиса в свободной форме.
func (p *Process) Caption() string {
	return p.caption
}

// ReadyTimeout - возвращает максимальное время, за которое должен быть запущен сервис.
func (p *Process) ReadyTimeout() time.Duration {
	return p.readyTimeout
}

// Start - запуск сервиса выполнения работы при старте приложения.
//
// Процесс работы:
//  1. Выполняет job.Do(ctx) с генерацией task_id;
//  2. После завершения job вызывает функцию ready();
//  3. Ожидает сигнала завершения (done или отмена контекста);
//
// Важно:
//   - Отмена внешнего контекста приведёт к завершению процесса;
//   - Для корректной остановки используйте Shutdown;
//   - Повторный запуск того же объекта не поддерживается.
func (p *Process) Start(ctx context.Context, ready func()) error {
	p.wg.Add(1)
	defer p.wg.Done()

	p.logger.Debug(ctx, "Starting the startup process...")
	defer p.logger.Debug(ctx, "The startup process has been stopped")

	if err := p.execJob(ctx); err != nil {
		return err
	}

	p.logger.Debug(ctx, "The job of the process is completed")

	if ready != nil {
		ready()
	}

	select {
	case <-p.done:
	case <-ctx.Done():
		p.logger.Debug(ctx, "The startup process detected context 'Done'", "error", ctx.Err())
	}

	return nil
}

// Shutdown - корректная остановка сервиса выполнения работы.
// Завершает ожидание и останавливает процесс.
func (p *Process) Shutdown(ctx context.Context) error {
	p.logger.Debug(ctx, "Shutting down the startup process...")

	p.closeOnce.Do(func() {
		close(p.done)
	})

	p.wg.Wait()
	p.logger.Debug(ctx, "The startup process has been shut down")

	return nil
}

func (p *Process) execJob(ctx context.Context) (err error) {
	ctx = p.traceManager.WithGeneratedProcessID(ctx, keyTaskID)
	p.logger.Debug(ctx, "Execute the job", "job_name", p.Caption())

	// паника job не должна ронять приложение - перехватывается и превращается в ошибку
	defer func() {
		if rvr := recover(); rvr != nil {
			err = mrprocess.NewCaughtPanicError("startup process: "+p.Caption(), rvr)
		}
	}()

	if err = p.job.Do(ctx); err != nil {
		return err
	}

	p.logger.Debug(ctx, "The job is completed", "job_name", p.Caption())

	return nil
}
