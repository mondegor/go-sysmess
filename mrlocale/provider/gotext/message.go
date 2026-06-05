package gotext

import (
	"strconv"

	"github.com/mondegor/go-sysmess/internal/msg"
)

// NewMessageFormatter - создаёт MessageFormatter для преобразования
// плейсхолдеров в формат, совместимый с fmt.Sprintf.
// Параметры leftDelim и rightDelim - ограничители плейсхолдеров (например: "{{" и "}}").
// Плейсхолдеры преобразуются в формат "%[N]s", где N - порядковый номер аргумента.
func NewMessageFormatter(leftDelim, rightDelim string) *msg.MessageFormatter {
	return msg.NewMessageFormatter(
		leftDelim,
		rightDelim,
		func(_ string, i int) (newPlaceholder string) {
			return "%[" + strconv.Itoa(i+1) + "]s"
		},
	)
}

// MessageConverter - возвращает функцию для предварительной обработки сообщения.
// Преобразует плейсхолдеры с указанными ограничителями в формат fmt.Sprintf ("%[N]s").
func MessageConverter(leftDelim, rightDelim string) func(message string, args []any) (newMsg string, newArgs []any) {
	formatter := NewMessageFormatter(leftDelim, rightDelim)

	return func(message string, args []any) (newMsg string, newArgs []any) {
		newMsg, _ = formatter.Format(message)

		return newMsg, args
	}
}
