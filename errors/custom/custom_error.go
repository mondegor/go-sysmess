package custom

import (
	"errors"
	"strings"

	"github.com/mondegor/go-sysmess/errors/kind"
)

const (
	codeSeparator = "/"
	missingCode   = "!MISSINGCUSTOMCODE"
)

type (
	// Error - пользовательская ошибка с уточнённым кодом ошибки (customCode).
	// Применяется в слое представления, перед выводом ошибки пользователю.
	// Позволяет добавить контекст к базовой ошибке, указав конкретное поле.
	// Пример кода: "EmailAlreadyExists/userEmail", где EmailAlreadyExists - базовая ошибка,
	// а userEmail - поле, в котором произошла ошибка.
	Error interface {
		error

		IsKindUser() bool
		CustomCode() string
		Unwrap() error
	}

	customError struct {
		customCode string
		isKindUser bool
		err        error
		causeError error
	}
)

var (
	// ErrHasNilError - не указана пользовательская ошибка.
	ErrHasNilError = errors.New("custom error has a nil wrapped error")

	// ErrHasInternalError - пользовательская ошибка содержит внутреннюю ошибку.
	ErrHasInternalError = errors.New("custom error has an internal wrapped error")

	// ErrHasSystemError - пользовательская ошибка содержит системную ошибку.
	ErrHasSystemError = errors.New("custom error has an system wrapped error")

	// ErrHasUnexpectedError - пользовательская ошибка содержит необработанную ошибку.
	ErrHasUnexpectedError = errors.New("custom error has unexpected wrapped error")
)

// New - создаёт пользовательскую ошибку с уточнённым кодом.
// Если err == nil или err имеет тип kind.Internal или kind.System,
// ошибка помечается как невалидная с соответствующим causeError.
// Задача разработчика - обрабатывать ошибки ранее, чтобы не допускать такие ситуации.
func New(err error, customCode string) Error {
	if customCode == "" {
		customCode = missingCode
	}

	if err == nil {
		return &customError{
			customCode: customCode,
			causeError: ErrHasNilError,
		}
	}

	// в первую очередь ожидается пользовательская ошибка с её кодом
	if e, ok := err.(interface {
		Kind() kind.Enum
		Code() string
	}); ok && e.Kind() == kind.User {
		return &customError{
			customCode: e.Code() + codeSeparator + customCode,
			isKindUser: true,
			err:        err,
		}
	}

	switch kind.Extract(err) {
	case kind.Internal:
		return &customError{
			customCode: customCode,
			err:        err,
			causeError: ErrHasInternalError,
		}
	case kind.System:
		return &customError{
			customCode: customCode,
			err:        err,
			causeError: ErrHasSystemError,
		}
	default:
		return &customError{
			customCode: customCode,
			err:        err,
			causeError: ErrHasUnexpectedError,
		}
	}
}

// IsKindUser - сообщает, имеет ли обёрнутая ошибка тип kind.User.
func (e *customError) IsKindUser() bool {
	return e.isKindUser
}

// CustomCode - возвращает уточнённый код ошибки включая базовый код.
func (e *customError) CustomCode() string {
	return e.customCode
}

// Error - возвращает строковое представление ошибки.
func (e *customError) Error() string {
	var buf strings.Builder

	buf.Grow(len(e.customCode) + 132)
	buf.WriteString("#")
	buf.WriteString(e.customCode)
	buf.WriteString(" - ")

	if e.causeError != nil {
		buf.WriteString(e.causeError.Error())
		buf.WriteString(": ")
	}

	if e.err != nil {
		buf.WriteString(e.err.Error())
	} else {
		buf.WriteString("nil")
	}

	return buf.String()
}

// Unwrap - реализует интерфейс errors.Unwrap.
// Возвращает вложенную ошибку или ErrHasNilError, если исходная ошибка была nil.
func (e *customError) Unwrap() error {
	if e.err != nil {
		return e.err
	}

	return ErrHasNilError
}
