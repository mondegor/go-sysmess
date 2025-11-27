package parse

import (
	"fmt"
)

type (
	// ParamEmptyError - param value is empty.
	ParamEmptyError struct {
		Type string
	}
)

// NewParamEmptyError - создаёт объект ParamEmptyError.
func NewParamEmptyError(paramType string) *ParamEmptyError {
	return &ParamEmptyError{
		Type: paramType,
	}
}

func (e *ParamEmptyError) Error() string {
	return fmt.Sprintf("param value is empty (type='%s')", e.Type)
}
