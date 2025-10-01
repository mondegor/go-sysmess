package logtracer

import (
	"context"

	"github.com/mondegor/go-sysmess/mrlog"
)

type (
	// Tracer - трейсинг на основе логгера.
	Tracer struct {
		logger  mrlog.Logger
		enabled bool
	}
)

// NewTracer - создаёт объект Tracer.
func NewTracer(logger mrlog.Logger, enabledLevel mrlog.Level) *Tracer {
	return &Tracer{
		logger:  logger,
		enabled: logger.Enabled(enabledLevel),
	}
}

// NewDebugTracer - создаёт объект Tracer работающий в отладочном режиме.
func NewDebugTracer(logger mrlog.Logger) *Tracer {
	return NewTracer(logger, mrlog.LevelDebug)
}

// Enabled - сообщает, включен ли трейсер.
func (e *Tracer) Enabled() bool {
	return e.enabled
}

// Trace - если трейсер включён, то он логирует данные в режиме LevelTrace.
func (e *Tracer) Trace(ctx context.Context, args ...any) {
	if !e.enabled {
		return
	}

	e.logger.Log(ctx, mrlog.LevelTrace, "---", args...)
}
