package middleware

import (
	"context"
	"log/slog"
)

// BeforeHandle - вспомогательная функция возвращающая функцию в формате middleware с целью модификации
// объекта slog.Record непосредственно перед вызовом метода Handle в интерфейсе slog.Handler.
func BeforeHandle(fn func(ctx context.Context, record slog.Record) slog.Record) func(next slog.Handler) slog.Handler {
	return func(next slog.Handler) slog.Handler {
		return &decorator{
			handler:      next,
			beforeHandle: fn,
		}
	}
}
