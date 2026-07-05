package mrlog

import (
	"context"
)

type (
	// nopLogger - заглушка, реализующая интерфейс Logger.
	// Игнорирует все логируемые сообщения.
	nopLogger struct{}
)

// NopLogger - создаёт объект Logger, который ничего не делает.
func NopLogger() Logger {
	return nopLogger{}
}

// Debug - имитирует логирование.
func (l nopLogger) Debug(_ context.Context, _ string, _ ...any) {}

// DebugFunc - имитирует логирование.
func (l nopLogger) DebugFunc(_ context.Context, _ func() string, _ ...any) {}

// Info - имитирует логирование.
func (l nopLogger) Info(_ context.Context, _ string, _ ...any) {}

// Warn - имитирует логирование.
func (l nopLogger) Warn(_ context.Context, _ string, _ ...any) {}

// Error - имитирует логирование.
func (l nopLogger) Error(_ context.Context, _ string, _ ...any) {}
