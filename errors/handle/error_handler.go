package handle

import "context"

//go:generate mockgen -source=error_handler.go -destination=./mock/error_handler.go

type (
	// ErrorHandler - интерфейс обработчика ошибок.
	ErrorHandler interface {
		Handle(ctx context.Context, err error)
	}

	// ErrorHandlerFunc - обработчик ошибок в виде функции.
	ErrorHandlerFunc func(ctx context.Context, err error)
)

// Handle - реализация интерфейса ErrorHandler в виде функции для обработки ошибок.
func (f ErrorHandlerFunc) Handle(ctx context.Context, err error) {
	f(ctx, err)
}
