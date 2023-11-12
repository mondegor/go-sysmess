package mrerr

import "fmt"

type (
    FieldErrorList []FieldError
)

func NewList(items ...FieldError) *FieldErrorList {
    list := &FieldErrorList{}

    if len(items) > 0 {
        *list = append(*list, items...)
    }

    return list
}

func NewListWith(fieldID string, err error) *FieldErrorList {
    return &FieldErrorList{newFieldError(fieldID, err)}
}

func (e *FieldErrorList) Add(fieldID string, err error) {
    *e = append(*e, newFieldError(fieldID, err))
}

func (e *FieldErrorList) Error() string {
    return fmt.Sprintf("%v", *e)
}
