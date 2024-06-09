package mrerr

import "fmt"

type (
	// CustomError - пользовательская ошибка с персональным кодом.
	CustomError struct {
		customCode string
		err        *AppError
	}
)

var (
	errCustomErrorHasNil           = NewProto("errCustomErrorHasNil", ErrorKindUser, "custom error has nil").New()
	errCustomErrorHasExternalError = NewProto("errCustomErrorHasExternalError", ErrorKindUser, "custom error has an external error")
)

// NewCustomError - создаёт объект CustomError.
func NewCustomError(customCode string, err error) *CustomError {
	if err == nil {
		return &CustomError{
			customCode: customCode,
			err:        errCustomErrorHasNil,
		}
	}

	if e, ok := err.(*AppError); ok { //nolint:errorlint
		return &CustomError{
			customCode: customCode,
			err:        e,
		}
	}

	if e, ok := err.(*ProtoAppError); ok { //nolint:errorlint
		return &CustomError{
			customCode: customCode,
			err:        e.New(),
		}
	}

	return &CustomError{
		customCode: customCode,
		err:        errCustomErrorHasExternalError.Wrap(err),
	}
}

// CustomCode - возвращает персональный код ошибки.
func (e *CustomError) CustomCode() string {
	return e.customCode
}

// Err - возвращает вложенную ошибку привязанную к текущей ошибке.
func (e *CustomError) Err() *AppError {
	return e.err
}

// Error - возвращает ошибку в виде строки.
func (e *CustomError) Error() string {
	return fmt.Sprintf("error customCode=%s: {%s}", e.customCode, e.err.Error())
}
