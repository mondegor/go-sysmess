package context

import (
	"context"
)

type (
	ctxCorrelationKey struct{}
)

// WithCorrelationID - возвращает новый контекст с установленным ID корреляции запроса.
// ID корреляции используется для связывания цепочки запросов между сервисами.
func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, ctxCorrelationKey{}, correlationID)
}

// CorrelationID - извлекает ID корреляции из контекста.
// Возвращает пустую строку, если ID не установлен.
func CorrelationID(ctx context.Context) string {
	if value, ok := ctx.Value(ctxCorrelationKey{}).(string); ok {
		return value
	}

	return ""
}
