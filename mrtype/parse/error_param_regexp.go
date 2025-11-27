package parse

import (
	"fmt"
)

type (
	// ParamRegexpError - param value doesn't match regexp.
	ParamRegexpError struct {
		Type   string
		Regexp string
	}
)

// NewParamRegexpError - создаёт объект ParamRegexpError.
func NewParamRegexpError(paramType, regexp string) *ParamRegexpError {
	return &ParamRegexpError{
		Type:   paramType,
		Regexp: regexp,
	}
}

func (e *ParamRegexpError) Error() string {
	return fmt.Sprintf("param value doesn't match regexp (type='%s', regexp='%s')", e.Type, e.Regexp)
}
