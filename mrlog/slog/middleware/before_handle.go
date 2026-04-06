package middleware

import (
	"context"
	"log/slog"
)

// BeforeHandle - создаёт middleware-функцию для модификации slog.Record
// перед его обработкой в slog.Handler.
// Параметр fn вызывается для каждой записи лога и может изменять атрибуты, сообщение или контекст.
func BeforeHandle(fn func(ctx context.Context, record slog.Record) slog.Record) func(next slog.Handler) slog.Handler {
	return func(next slog.Handler) slog.Handler {
		return &decorator{
			handler:      next,
			beforeHandle: fn,
		}
	}
}
