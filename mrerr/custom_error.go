package mrerr

import "fmt"

const (
	customErrorCodePrefix = "mrerr_"
)

type (
	// CustomError - ошибка, которую допустимо отображать пользователю.
	CustomError struct {
		code string
		err  *AppError
	}
)

// NewCustomError - создаётся объект CustomError.
func NewCustomError(code string, err error) *CustomError {
	if err == nil {
		return NewCustomErrorMessage(code, "err is nil")
	}

	appArr, ok := err.(*AppError)
	if !ok {
		appArr = New(
			customErrorCodePrefix+code,
			err.Error(),
		)
	}

	return &CustomError{
		code: code,
		err:  appArr,
	}
}

// NewCustomErrorAppError - создаётся объект CustomError на основе AppError.
func NewCustomErrorAppError(code string, err *AppError) *CustomError {
	if err == nil {
		return NewCustomErrorMessage(code, "appErr is nil")
	}

	return &CustomError{
		code: code,
		err:  err,
	}
}

// NewCustomErrorMessage - создаётся объект CustomError на основе message.
func NewCustomErrorMessage(code, message string) *CustomError {
	return &CustomError{
		code: code,
		err: New(
			customErrorCodePrefix+code,
			message,
		),
	}
}

// Code - возвращает код ошибки.
func (e *CustomError) Code() string {
	return e.code
}

// AppError - возвращает вложенную ошибку, которая породила текущая ошибка.
func (e *CustomError) AppError() *AppError {
	return e.err
}

// Error - возвращает ошибку в виде строки.
func (e *CustomError) Error() string {
	return fmt.Sprintf("error code=%s: {%s}", e.code, e.err.Error())
}
