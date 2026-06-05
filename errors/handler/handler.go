package handler

import (
	"context"
)

type (
	// Handler - интерфейс обработчика ошибок.
	// Используется для централизованной маршрутизации ошибок (логирование, трассировка).
	Handler interface {
		Handle(ctx context.Context, err error)
	}

	// Func - адаптер, позволяющий использовать обычную функцию как Handler.
	Func func(ctx context.Context, err error)
)

// Handle - реализует интерфейс Handler, вызывая саму функцию f.
// Позволяет использовать обычную функцию как обработчик ошибок.
func (f Func) Handle(ctx context.Context, err error) {
	f(ctx, err)
}

type (
	// nop - заглушка, реализующая интерфейс Handler.
	// Игнорирует все обрабатываемые ошибки.
	nop struct{}
)

// Nop создаёт Handler-заглушку.
func Nop() Handler {
	return nop{}
}

// Handle - игнорирует ошибку (заглушка для интерфейса Handler).
func (t nop) Handle(_ context.Context, _ error) {}
