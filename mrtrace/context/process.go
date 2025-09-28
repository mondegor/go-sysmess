package context

import (
	"context"
)

type (
	ctxProcessKey struct{}
)

// WithProcessID - возвращает контекст с установленным в него указанным идентификатором процесса.
func WithProcessID(ctx context.Context, processID string) context.Context {
	return context.WithValue(ctx, ctxProcessKey{}, processID)
}

// ProcessID - возвращает из контекста указанный идентификатор текущего процесса.
func ProcessID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxProcessKey{}).(string); ok {
		return value
	}

	return ""
}
