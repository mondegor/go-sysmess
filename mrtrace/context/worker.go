package context

import (
	"context"
)

type (
	ctxWorkerKey struct{}
)

// WithWorkerID - возвращает контекст с установленным в него указанным идентификатором воркера.
func WithWorkerID(ctx context.Context, workerID string) context.Context {
	return context.WithValue(ctx, ctxWorkerKey{}, workerID)
}

// WorkerID - возвращает из контекста указанный идентификатор текущего воркера.
func WorkerID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxWorkerKey{}).(string); ok {
		return value
	}

	return ""
}
