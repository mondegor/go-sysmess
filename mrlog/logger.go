package mrlog

import (
	"context"
)

type (
	// Logger - интерфейс логирования сообщений и ошибок с использованием контекста.
	Logger interface {
		WithAttrs(args ...any) Logger
		Enabled(level Level) bool

		Debug(ctx context.Context, msg string, args ...any)
		DebugFunc(ctx context.Context, createMsg func() string, args ...any)
		Info(ctx context.Context, msg string, args ...any)
		Warn(ctx context.Context, msg string, args ...any)
		Error(ctx context.Context, msg string, args ...any)

		Log(ctx context.Context, level Level, message string, args ...any)
	}

	// LiteLogger - упрощённый интерфейс логирования сообщений и ошибок.
	LiteLogger interface {
		Debug(msg string, args ...any)
		DebugFunc(createMsg func() string, args ...any)
		Info(msg string, args ...any)
		Warn(msg string, args ...any)
		Error(msg string, args ...any)

		ContextLogger() Logger
	}
)
