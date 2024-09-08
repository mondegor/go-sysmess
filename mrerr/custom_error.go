package mrerr

import (
	"fmt"
)

type (
	// CustomError - пользовательская ошибка с персональным кодом.
	CustomError struct {
		customCode string
		err        *AppError
	}
)

var (
	// ErrCustomErrorHasInternalError - пользовательская ошибка содержит внутреннюю или системную ошибку.
	ErrCustomErrorHasInternalError = NewProto(
		"errCustomErrorHasInternalError", ErrorKindInternal, "custom error has an internal error")

	// ErrCustomErrorHasNoWrappedError - пользовательская ошибка содержит необработанную ошибку.
	ErrCustomErrorHasNoWrappedError = NewProto(
		"errCustomErrorHasNoWrappedError", ErrorKindInternal, "custom error has no wrapped error")
)

// NewCustomError - создаёт объект CustomError, аргумент err должен содержать ошибку.
func NewCustomError(customCode string, err error) *CustomError {
	newError := func(customCode string, err *AppError) *CustomError {
		return &CustomError{
			customCode: customCode,
			err:        err,
		}
	}

	if err == nil {
		return newError(customCode, ErrErrorIsNilPointer.New())
	}

	if e, ok := err.(*AppError); ok { //nolint:errorlint
		if e.Kind() == ErrorKindUser {
			return newError(customCode, e)
		}

		return newError(customCode, ErrCustomErrorHasInternalError.Wrap(err))
	}

	if e, ok := err.(*ProtoAppError); ok { //nolint:errorlint
		if e.Kind() == ErrorKindUser {
			return newError(customCode, e.New())
		}

		return newError(customCode, ErrCustomErrorHasInternalError.Wrap(err))
	}

	return newError(customCode, ErrCustomErrorHasNoWrappedError.Wrap(err))
}

// CustomCode - возвращает персональный код ошибки.
func (e *CustomError) CustomCode() string {
	return e.customCode
}

// IsValid - возвращает true, если внутри содержится
// пользовательская ошибка, все остальные ошибки считаются невалидными,
// программисту необходимо позаботиться их обернуть в пользовательский вид ошибки.
func (e *CustomError) IsValid() bool {
	return e.err.Kind() == ErrorKindUser
}

// Err - возвращает вложенную ошибку, привязанную к текущей ошибке.
func (e *CustomError) Err() *AppError {
	return e.err
}

// Error - возвращает ошибку в виде строки.
func (e *CustomError) Error() string {
	return fmt.Sprintf("error customCode=%s: {%s}", e.customCode, e.err.Error())
}
