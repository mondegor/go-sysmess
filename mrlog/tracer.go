package mrlog

import (
	"context"
)

type (
	// Tracer - трейсинг на основе логгера.
	Tracer struct {
		logger  Logger
		enabled bool
	}
)

// NewTracer - создаёт объект Tracer.
func NewTracer(logger Logger, enabledLevel Level) *Tracer {
	return &Tracer{
		logger:  logger,
		enabled: logger.Enabled(enabledLevel),
	}
}

// NewDebugTracer - создаёт объект Tracer работающий в отладочном режиме.
func NewDebugTracer(logger Logger) *Tracer {
	return NewTracer(logger, LevelDebug)
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

	e.logger.Log(ctx, LevelTrace, "---", args...)
}
