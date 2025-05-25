package mrerrors

import (
	"strings"
)

type (
	wrappedError InstantError
)

// Error - возвращает вложенную ошибку в виде строки, здесь выводится минимальная информация,
// чтобы исключить дублирование при выводе информации в основной ошибке.
func (e *wrappedError) Error() string {
	var buf strings.Builder

	if err := e.messageReplacer.ReplaceTo(&buf, e.args[:e.messageReplacer.CountArgs()]); err != nil {
		buf.WriteString(e.message)
	}

	buf.WriteString(" [")
	buf.WriteString(e.Kind().String())

	if e.code != "" {
		buf.WriteString(", ")
		buf.WriteString(e.code)
	}

	buf.WriteByte(']')

	if e.err != nil {
		buf.WriteString(": ")
		buf.WriteString(e.err.Error())
	}

	return buf.String()
}
