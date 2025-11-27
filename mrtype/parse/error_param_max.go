package parse

import (
	"fmt"
)

type (
	// ParamMaxValueError - param contains value greater then max.
	ParamMaxValueError struct {
		Type     string
		MaxValue any
	}
)

// NewParamMaxValueError - создаёт объект ParamMaxValueError.
func NewParamMaxValueError(paramType string, maxValue any) *ParamMaxValueError {
	return &ParamMaxValueError{
		Type:     paramType,
		MaxValue: maxValue,
	}
}

func (e *ParamMaxValueError) Error() string {
	return fmt.Sprintf("param contains value greater then max (type='%s', maxValue='%v')", e.Type, e.MaxValue)
}
