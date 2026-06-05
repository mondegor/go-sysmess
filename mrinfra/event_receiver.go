package mrinfra

import (
	"context"
	"time"

	"github.com/mondegor/go-sysmess/mrevent"
	tracectx "github.com/mondegor/go-sysmess/mrtrace/context"
	"github.com/mondegor/go-sysmess/mrtrace/process"
)

const (
	// defaultReceiveTimeout - таймаут обработки события получателем по умолчанию.
	defaultReceiveTimeout = 5 * time.Second
)

type (
	// EventReceiver - получатель событий, который асинхронно распределяет события
	// между зарегистрированными получателями с индивидуальным таймаутом для каждого.
	// Каждый получатель выполняется в отдельной горутине.
	EventReceiver struct {
		traceManager   process.ContextManager
		receiveTimeout time.Duration
		receivers      []mrevent.Receiver
	}
)

// NewEventReceiver - создаёт новый EventReceiver для асинхронной обработки событий.
// Параметры:
//   - traceManager - менеджер трейсинга для добавления ID процесса в контекст;
//   - receiveTimeout - таймаут на обработку события каждым получателем (если 0, используется defaultReceiveTimeout);
//   - receivers - список получателей, которым будут отправляться события.
func NewEventReceiver(traceManager process.ContextManager, receiveTimeout time.Duration, receivers []mrevent.Receiver) *EventReceiver {
	if receiveTimeout == 0 {
		receiveTimeout = defaultReceiveTimeout
	}

	return &EventReceiver{
		traceManager:   traceManager,
		receiveTimeout: receiveTimeout,
		receivers:      receivers,
	}
}

// Receive - асинхронно распределяет событие всем зарегистрированным получателям.
// Каждый получатель выполняется в отдельной горутине с индивидуальным таймаутом.
// В контекст каждого получателя добавляется уникальный ID процесса (ProcessID) для трейсинга.
// Метод не блокируется и не ожидает завершения обработки получателями.
func (er *EventReceiver) Receive(ctx context.Context, eventName string, args ...any) {
	for _, r := range er.receivers {
		go func(ctx context.Context) {
			// устанавливается индивидуальный таймаут, чтобы ограничить работу получателей
			ctx, cancel := context.WithTimeout(
				er.traceManager.WithGeneratedProcessID(ctx, tracectx.KeyTaskID),
				er.receiveTimeout,
			)
			defer cancel()

			r.Receive(ctx, eventName, args...)
		}(ctx)
	}
}
