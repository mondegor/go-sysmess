package parse

import (
	"fmt"
)

type (
	// ParamIncorrectError - ошибка некорректного значения параметра.
	// Содержит исходную ошибку парсинга или валидации.
	ParamIncorrectError struct {
		Type  string
		Cause error
	}
)

// NewParamIncorrectError - создаёт ошибку ParamIncorrectError для указанного типа параметра.
// cause - исходная ошибка, приведшая к ошибке парсинга.
func NewParamIncorrectError(paramType string, cause error) *ParamIncorrectError {
	return &ParamIncorrectError{
		Type:  paramType,
		Cause: cause,
	}
}

func (e *ParamIncorrectError) Error() string {
	return fmt.Sprintf("param contains incorrect value '%s' (type='%s')", e.Cause.Error(), e.Type)
}
