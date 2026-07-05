package context

import (
	"context"
)

type (
	ctxProcessKey struct{}
)

// WithProcessID - возвращает новый контекст с установленным ID процесса.
func WithProcessID(ctx context.Context, processID string) context.Context {
	return context.WithValue(ctx, ctxProcessKey{}, processID)
}

// ProcessID - извлекает ID процесса из контекста.
// Возвращает пустую строку, если ID не установлен.
func ProcessID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxProcessKey{}).(string); ok {
		return value
	}

	return ""
}
