package parse

import (
	"fmt"
)

type (
	// ParamLenMaxError - param has value length greater then max characters.
	ParamLenMaxError struct {
		Type   string
		MaxLen int
	}
)

// NewParamLenMaxError - создаёт объект ParamLenMaxError.
func NewParamLenMaxError(paramType string, maxValue int) *ParamLenMaxError {
	return &ParamLenMaxError{
		Type:   paramType,
		MaxLen: maxValue,
	}
}

func (e *ParamLenMaxError) Error() string {
	return fmt.Sprintf("param has value length greater then max characters (type='%s', maxLen=%d)", e.Type, e.MaxLen)
}
