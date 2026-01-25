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
