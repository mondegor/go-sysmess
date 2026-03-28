package errors

import (
	"context"
)

type (
	// Handler - интерфейс обработчика ошибок.
	Handler interface {
		Handle(ctx context.Context, err error)
	}

	// HandlerFunc - обработчик ошибок в виде функции.
	HandlerFunc func(ctx context.Context, err error)
)

// Handle - реализация интерфейса Handler в виде функции для обработки ошибок.
func (f HandlerFunc) Handle(ctx context.Context, err error) {
	f(ctx, err)
}

type (
	nopHandler struct{}
)

// NopHandler - создаёт объект Handler, который ничего не делает.
func NopHandler() Handler {
	return nopHandler{}
}

// Handle - имитирует обработку ошибки.
func (t nopHandler) Handle(_ context.Context, _ error) {}
