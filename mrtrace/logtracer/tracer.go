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

// New - создаёт объект Tracer.
func New(logger mrlog.Logger) *Tracer {
	return &Tracer{
		logger:  logger,
		enabled: logger.Enabled(mrlog.LevelError),
	}
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
