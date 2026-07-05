package context

import (
	"context"
)

type (
	ctxWorkerKey struct{}
)

// WithWorkerID - возвращает новый контекст с установленным ID воркера.
func WithWorkerID(ctx context.Context, workerID string) context.Context {
	return context.WithValue(ctx, ctxWorkerKey{}, workerID)
}

// WorkerID - извлекает ID воркера из контекста.
// Возвращает пустую строку, если ID не установлен.
func WorkerID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxWorkerKey{}).(string); ok {
		return value
	}

	return ""
}
