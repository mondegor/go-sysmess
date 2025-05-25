package mrerr

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerrors"
)

type (
	// CustomError - пользовательская ошибка с уточнённым кодом ошибки.
	// Например, код может выглядеть следующим образом: EmailAlreadyExists/userEmail
	// Где EmailAlreadyExists - пользовательская ошибка, userEmail - поле, в которой произошла ошибка.
	CustomError struct {
		customCode string
		err        *mrerrors.InstantError
	}
)

var (
	// ErrCustomErrorHasNilError - пользовательская ошибка содержит nil.
	ErrCustomErrorHasNilError = NewKindInternal("custom error has an nil error")

	// ErrCustomErrorHasInternalError - пользовательская ошибка содержит внутреннюю ошибку.
	ErrCustomErrorHasInternalError = NewKindInternal("custom error has an internal error")

	// ErrCustomErrorHasSystemError - пользовательская ошибка содержит системную ошибку.
	ErrCustomErrorHasSystemError = NewKindSystem("custom error has an system error")

	// ErrCustomErrorHasNoWrappedError - пользовательская ошибка содержит необработанную ошибку.
	ErrCustomErrorHasNoWrappedError = NewKindInternal("custom error has no wrapped error")
)

// NewCustomError - создаёт объект CustomError.
// Если аргумент err содержит любую ошибку, которая не соответствует типу ErrorKindUser
// то вся эта ошибка будет считаться невалидной.
func NewCustomError(customCode string, err error) *CustomError {
	if err == nil {
		return newCustomError(customCode, ErrCustomErrorHasNilError.New())
	}

	var appErr *mrerrors.InstantError

	switch e := err.(type) { //nolint:errorlint
	case *mrerrors.InstantError:
		appErr = e
	case *mrerrors.ProtoError:
		appErr = mrerrors.CastProto(e)
	}

	if appErr == nil {
		return newCustomError(customCode, ErrCustomErrorHasNoWrappedError.Wrap(err))
	}

	if appErr.Kind() == ErrorKindUser {
		return newCustomError(customCode, appErr)
	}

	if appErr.Kind() == ErrorKindSystem {
		return newCustomError(customCode, ErrCustomErrorHasSystemError.Wrap(appErr))
	}

	return newCustomError(customCode, ErrCustomErrorHasInternalError.Wrap(appErr))
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

// Err - возвращает вложенную ошибку.
func (e *CustomError) Err() *mrerrors.InstantError {
	return e.err
}

// Error - возвращает ошибку в виде строки.
func (e *CustomError) Error() string {
	return fmt.Sprintf("customCode=%s: {%s}", e.customCode, e.err.Error())
}

func newCustomError(customCode string, err *mrerrors.InstantError) *CustomError {
	return &CustomError{
		customCode: customCode,
		err:        err,
	}
}
