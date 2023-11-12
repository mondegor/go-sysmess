package mrerr

const (
    fieldErrorID = "errFieldMessage"
)

type (
    FieldError struct {
        ID string
        Err *AppError
    }
)

func newFieldError(id string, err error) FieldError {
    appArr, ok := err.(*AppError)

    if !ok {
        appArr = New(
            fieldErrorID,
            err.Error(),
        )
    }

    return FieldError{
        ID: id,
        Err: appArr,
    }
}
