package mrerr

import (
	"strings"
)

type (
	FieldErrorList []*FieldError
)

func (l *FieldErrorList) Add(fieldID string, err error) {
	*l = append(*l, NewFieldError(fieldID, err))
}

func (l *FieldErrorList) AddAppErr(fieldID string, err *AppError) {
	*l = append(*l, NewFieldErrorAppErr(fieldID, err))
}

func (l FieldErrorList) Error() string {
	var buf strings.Builder

	for i := 0; i < len(l); i++ {
		buf.WriteString(l[i].Error())
		buf.WriteString("\n")
	}

	return strings.TrimSpace(buf.String())
}
