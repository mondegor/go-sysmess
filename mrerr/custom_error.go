package mrerr

import "fmt"

const (
	customErrorCodePrefix = "mrerr_"
)

type (
	CustomError struct {
		code string
		err  *AppError
	}
)

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

func NewCustomErrorAppError(code string, err *AppError) *CustomError {
	if err == nil {
		return NewCustomErrorMessage(code, "appErr is nil")
	}

	return &CustomError{
		code: code,
		err:  err,
	}
}

func NewCustomErrorMessage(code, message string) *CustomError {
	return &CustomError{
		code: code,
		err: New(
			customErrorCodePrefix+code,
			message,
		),
	}
}

func (e *CustomError) Code() string {
	return e.code
}

func (e *CustomError) AppError() *AppError {
	return e.err
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error code=%s: {%s}", e.code, e.err.Error())
}
