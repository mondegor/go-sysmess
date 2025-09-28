package context

import (
	"context"
)

type (
	ctxTraceKey struct{}
)

// WithTraceID - возвращает контекст с установленным в него указанным идентификатором трейса.
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, ctxTraceKey{}, traceID)
}

// TraceID - возвращает из контекста указанные идентификатор текущего трейса.
func TraceID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxTraceKey{}).(string); ok {
		return value
	}

	return ""
}
