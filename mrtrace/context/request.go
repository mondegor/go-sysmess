package context

import (
	"context"
)

type (
	ctxRequestKey struct{}
)

// WithRequestID - возвращает контекст с установленным в него указанным идентификатором запроса.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ctxRequestKey{}, requestID)
}

// RequestID - возвращает из контекста указанный идентификатор текущего запроса.
func RequestID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxRequestKey{}).(string); ok {
		return value
	}

	return ""
}
