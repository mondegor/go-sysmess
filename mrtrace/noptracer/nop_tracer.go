package noptracer

import (
	"context"
)

type (
	// Tracer - заглушка реализующая интерфейс трейсера.
	Tracer struct{}
)

// New - создаёт объект Tracer.
func New() *Tracer {
	return &Tracer{}
}

// Enabled - всегда возвращает false.
func (e *Tracer) Enabled() bool {
	return false
}

// Trace - имитирует запись данных.
func (e *Tracer) Trace(_ context.Context, _ ...any) {}
