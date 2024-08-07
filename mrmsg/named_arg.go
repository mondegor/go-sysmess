package mrmsg

import (
	"fmt"
	"strconv"
)

type (
	// NamedArg - именованный аргумент используемый в сообщениях.
	NamedArg struct {
		Name  string
		Value any
	}
)

// ValueString - преобразовывает значение аргумента в строку.
func (a *NamedArg) ValueString() string {
	switch val := a.Value.(type) {
	case string:
		return val

	case int:
		return strconv.FormatInt(int64(val), 10)

	case int32:
		return strconv.FormatInt(int64(val), 10)

	case int64:
		return strconv.FormatInt(val, 10)

	default:
		return fmt.Sprintf("%+v", val)
	}
}
