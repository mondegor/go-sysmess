package errors

import (
	"context"
)

type (
	// Handler - интерфейс обработчика ошибок.
	// Используется для централизованной маршрутизации ошибок (логирование, трассировка).
	Handler interface {
		Handle(ctx context.Context, err error)
	}

	// HandlerFunc - адаптер, позволяющий использовать обычную функцию как Handler.
	HandlerFunc func(ctx context.Context, err error)
)

// Handle - реализует интерфейс Handler, вызывая саму функцию f.
// Позволяет использовать обычную функцию как обработчик ошибок.
func (f HandlerFunc) Handle(ctx context.Context, err error) {
	f(ctx, err)
}

type (
	// nopHandler - заглушка, реализующая интерфейс Handler.
	// Игнорирует все обрабатываемые ошибки.
	nopHandler struct{}
)

// NopHandler создаёт Handler-заглушку.
func NopHandler() Handler {
	return nopHandler{}
}

// Handle - игнорирует ошибку (заглушка для интерфейса Handler).
func (t nopHandler) Handle(_ context.Context, _ error) {}
