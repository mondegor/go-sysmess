package parse

import (
	"fmt"
)

type (
	// ParamIncorrectError - param contains incorrect value.
	ParamIncorrectError struct {
		Type  string
		Cause error
	}
)

// NewParamIncorrectError - создаёт объект ParamIncorrectError.
func NewParamIncorrectError(paramType string, cause error) *ParamIncorrectError {
	return &ParamIncorrectError{
		Type:  paramType,
		Cause: cause,
	}
}

func (e *ParamIncorrectError) Error() string {
	return fmt.Sprintf("param contains incorrect value: %s (type='%s')", e.Cause.Error(), e.Type)
}
