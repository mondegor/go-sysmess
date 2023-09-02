package mrerr

const (
    errorIdField = "errMessageForField"
)

type (
    FieldError struct {
        Id string
        Err AppError
    }
)

func newFieldError(id string, err error) FieldError {
    appArr, ok := err.(AppError)

    if !ok {
        appArr = New(
            errorIdField,
            err.Error(),
        )
    }

    return FieldError{
        Id: id,
        Err: appArr,
    }
}
