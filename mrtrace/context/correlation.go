package context

import (
	"context"
)

type (
	ctxCorrelationKey struct{}
)

// WithCorrelationID - возвращает контекст с установленным в него указанным идентификатором корреляции запроса.
func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, ctxCorrelationKey{}, correlationID)
}

// CorrelationID - возвращает из контекста указанный идентификатор корреляции текущего запроса.
func CorrelationID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxCorrelationKey{}).(string); ok {
		return value
	}

	return ""
}
