package mrerr

import "fmt"

const (
	fieldErrorID           = "errFieldMessage"
	fieldErrorUnknownError = "unknown error (err = nil)"
)

type (
	FieldError struct {
		id  string
		err *AppError
	}
)

func NewFieldError(id string, err error) *FieldError {
	if err == nil {
		return NewFieldMessage(id, fieldErrorUnknownError)
	}

	appArr, ok := err.(*AppError)

	if !ok {
		appArr = New(
			fieldErrorID,
			err.Error(),
		)
	}

	return &FieldError{
		id:  id,
		err: appArr,
	}
}

func NewFieldErrorAppErr(id string, err *AppError) *FieldError {
	if err == nil {
		return NewFieldMessage(id, fieldErrorUnknownError)
	}

	return &FieldError{
		id:  id,
		err: err,
	}
}

func NewFieldMessage(id string, message string) *FieldError {
	return &FieldError{
		id: id,
		err: New(
			fieldErrorID,
			message,
		),
	}
}

func (e *FieldError) ID() string {
	return e.id
}

func (e *FieldError) Kind() ErrorKind {
	return e.err.kind
}

func (e *FieldError) AppError() *AppError {
	return e.err
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.id, e.err.Error())
}
