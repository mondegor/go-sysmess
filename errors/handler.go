package errors

import (
	"github.com/mondegor/go-sysmess/errors/handler"
)

type (
	// Handler - интерфейс обработчика ошибок.
	// Используется для централизованной маршрутизации ошибок (логирование, трассировка).
	Handler = handler.Handler

	// HandlerFunc - адаптер, позволяющий использовать обычную функцию как Handler.
	HandlerFunc = handler.Func
)

// NopHandler создаёт Handler-заглушку.
func NopHandler() Handler {
	return handler.Nop()
}
