package context

import (
	"context"
)

type (
	ctxRequestKey struct{}
)

// WithRequestID - возвращает новый контекст с установленным ID запроса.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ctxRequestKey{}, requestID)
}

// RequestID - извлекает ID запроса из контекста.
// Возвращает пустую строку, если ID не установлен.
func RequestID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxRequestKey{}).(string); ok {
		return value
	}

	return ""
}
