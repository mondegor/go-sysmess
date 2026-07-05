package context

import (
	"context"
)

type (
	ctxTaskKey struct{}
)

// WithTaskID - возвращает новый контекст с установленным ID задачи.
func WithTaskID(ctx context.Context, taskID string) context.Context {
	return context.WithValue(ctx, ctxTaskKey{}, taskID)
}

// TaskID - извлекает ID задачи из контекста.
// Возвращает пустую строку, если ID не установлен.
func TaskID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxTaskKey{}).(string); ok {
		return value
	}

	return ""
}
