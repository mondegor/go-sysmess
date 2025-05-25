package mrerrors

import (
	"strings"
)

type (
	userArgError InstantError
)

// Error - возвращает ошибку в виде строки.
func (e *userArgError) Error() string {
	var buf strings.Builder

	if err := e.messageReplacer.ReplaceTo(&buf, e.args[:e.messageReplacer.CountArgs()]); err != nil {
		buf.WriteString(e.message)
	}

	return buf.String()
}
