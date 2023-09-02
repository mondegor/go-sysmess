package mrerr

import "fmt"

type (
    FieldErrorList []FieldError
)

func NewList(items ...FieldError) FieldErrorList {
    if len(items) > 0 {
        return append(FieldErrorList{}, items...)
    }

    return FieldErrorList{}
}

func NewListWith(fieldId string, err error) FieldErrorList {
    return FieldErrorList{newFieldError(fieldId, err)}
}

func (e *FieldErrorList) Add(fieldId string, err error) {
    *e = append(*e, newFieldError(fieldId, err))
}

func (e *FieldErrorList) Error() string {
    return fmt.Sprintf("%+v", *e)
}
