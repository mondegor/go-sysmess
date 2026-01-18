package errors

import "context"

//go:generate mockgen -source=error_handler.go -destination=./mock/error_handler.go

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
