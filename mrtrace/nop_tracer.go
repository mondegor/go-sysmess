package mrtrace

import (
	"context"
)

type (
	nopTracer struct{}
)

// NopTracer - создаёт объект Tracer, который ничего не делает.
func NopTracer() Tracer {
	return nopTracer{}
}

// Trace - имитирует запись данных.
func (t nopTracer) Trace(_ context.Context, _ ...any) {}
