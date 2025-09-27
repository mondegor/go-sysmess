package mrlog

import (
	"context"
)

type (
	// Tracer - comment struct.
	Tracer struct {
		logger Logger
	}
)

// NewTracer - создаёт объект Tracer.
func NewTracer(logger Logger) *Tracer {
	return &Tracer{
		logger: logger,
	}
}

// Enabled - comment method.
func (e *Tracer) Enabled() bool {
	return e.logger != nil
}

// Trace - comment method.
func (e *Tracer) Trace(ctx context.Context, args ...any) {
	if e.logger != nil {
		e.logger.Log(ctx, LevelTrace, "---", args...)
	}
}
