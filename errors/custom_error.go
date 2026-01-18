package errors

import (
	"github.com/mondegor/go-sysmess/errors/custom"
)

type (
	// CustomError - пользовательская ошибка с уточнённым кодом ошибки.
	CustomError = custom.Error

	// CustomListError - список пользовательских ошибок.
	CustomListError = custom.ListError
)

// WithCustomCode - создаёт CustomError для оборачивания ошибок типа kind.User
// с добавлением дополнительного кода ошибки.
func WithCustomCode(err error, customCode string) CustomError {
	return custom.New(err, customCode)
}
