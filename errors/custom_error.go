package errors

import (
	"github.com/mondegor/go-core/errors/custom"
)

type (
	// CustomError - пользовательская ошибка с уточнённым кодом.
	CustomError = custom.Error

	// CustomListError - слайс пользовательских ошибок с уточнённым кодом.
	CustomListError = custom.ListError
)

// WithCustomCode - создаёт CustomError, оборачивающую ошибку типа kind.User
// с дополнительным уточняющим кодом (например: имя поля: "userEmail").
// Если err не является kind.User, ошибка помечается как невалидная.
func WithCustomCode(err error, customCode string) CustomError {
	return custom.New(err, customCode)
}
