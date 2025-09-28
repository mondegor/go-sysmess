package context

import (
	"context"
)

type (
	ctxTaskKey struct{}
)

// WithTaskID - возвращает контекст с установленным в него указанным идентификатором задачи.
func WithTaskID(ctx context.Context, taskID string) context.Context {
	return context.WithValue(ctx, ctxTaskKey{}, taskID)
}

// TaskID - возвращает из контекста указанные идентификатор текущей задачи.
func TaskID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxTaskKey{}).(string); ok {
		return value
	}

	return ""
}
