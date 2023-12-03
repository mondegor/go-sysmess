package mrerr

import "fmt"

const (
	fieldErrorIDPrefix = "field_err"
)

type (
	FieldError struct {
		id  string
		err *AppError
	}
)

func NewFieldError(id string, err error) *FieldError {
	if err == nil {
		return NewFieldErrorMessage(id, "err is nil")
	}

	appArr, ok := err.(*AppError)

	if !ok {
		appArr = New(
			fieldErrorIDPrefix+"_"+id,
			err.Error(),
		)
	}

	return &FieldError{
		id:  id,
		err: appArr,
	}
}

func NewFieldErrorAppError(id string, err *AppError) *FieldError {
	if err == nil {
		return NewFieldErrorMessage(id, "appErr is nil")
	}

	return &FieldError{
		id:  id,
		err: err,
	}
}

func NewFieldErrorMessage(id string, message string) *FieldError {
	return &FieldError{
		id: id,
		err: New(
			fieldErrorIDPrefix+"_"+id,
			message,
		),
	}
}

func (e *FieldError) ID() string {
	return e.id
}

func (e *FieldError) AppError() *AppError {
	return e.err
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("fieldId=%s; err={%s}", e.id, e.err.Error())
}
