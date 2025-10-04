package mrinfra

import (
	"context"
	"time"

	"github.com/mondegor/go-sysmess/mrevent"
	"github.com/mondegor/go-sysmess/mrtrace"
)

const (
	defaultReceiveTimeout = 5 * time.Second
)

type (
	// EventReceiver - отвечает за отправку событий накладывая ограничения получателей.
	EventReceiver struct {
		traceManager   mrtrace.ContextManager
		receiveTimeout time.Duration
		receivers      []mrevent.Receiver
	}
)

// NewEventReceiver - создаёт объект EventReceiver.
func NewEventReceiver(traceManager mrtrace.ContextManager, receiveTimeout time.Duration, receivers []mrevent.Receiver) *EventReceiver {
	if receiveTimeout == 0 {
		receiveTimeout = defaultReceiveTimeout
	}

	return &EventReceiver{
		traceManager:   traceManager,
		receiveTimeout: receiveTimeout,
		receivers:      receivers,
	}
}

// Receive - отправляет указанное событие.
func (er *EventReceiver) Receive(ctx context.Context, eventName string, args ...any) {
	// WARNING: используется новый контекст со скопированными ID процессами из основного контекста,
	// чтобы гарантировать завершение работы получателей при отмене основного контекста
	ctx = er.traceManager.NewContextWithIDs(ctx)

	for _, r := range er.receivers {
		go func(ctx context.Context) {
			// устанавливается индивидуальный таймаут, чтобы ограничить работу получателей
			ctx, cancel := context.WithTimeout(
				er.traceManager.WithGeneratedProcessID(ctx, mrtrace.KeyTaskID),
				er.receiveTimeout,
			)
			defer cancel()

			r.Receive(ctx, eventName, args...)
		}(ctx)
	}
}
