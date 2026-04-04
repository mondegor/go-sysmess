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
	// Error - пользовательская ошибка с уточнённым кодом ошибки.
	// Используется в слое представления, непосредственно перед выводом ошибки пользователю.
	// Например, код может выглядеть следующим образом: EmailAlreadyExists/userEmail
	// Где EmailAlreadyExists - пользовательская ошибка, userEmail - поле, в которой произошла ошибка.
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

// New - создаёт объект Error.
// Если аргумент err содержит любую ошибку, у которой тип отличается от kind.User,
// то созданная ошибка будет считаться невалидной и при отображении будет выдана системная ошибка.
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

// IsKindUser - возвращает true, если внутри содержится пользовательская ошибка,
// все остальные ошибки считаются невалидными,
// программисту необходимо позаботиться их обернуть в пользовательский вид ошибки.
func (e *customError) IsKindUser() bool {
	return e.isKindUser
}

// CustomCode - возвращает персональный код ошибки.
func (e *customError) CustomCode() string {
	return e.customCode
}

// Error - возвращает ошибку в виде строки.
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

// Unwrap - возвращает вложенную ошибку.
func (e *customError) Unwrap() error {
	if e.err != nil {
		return e.err
	}

	return ErrHasNilError
}
