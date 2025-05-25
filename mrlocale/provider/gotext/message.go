package gotext

import (
	"strconv"

	"github.com/mondegor/go-sysmess/mrmsg"
)

// NewMessageFormatter - создаёт объект mrmsg.MessageFormatter.
func NewMessageFormatter(leftDelim, rightDelim string) *mrmsg.MessageFormatter {
	return mrmsg.NewMessageFormatter(
		leftDelim,
		rightDelim,
		func(_ string, i int) (newPlaceholder string) {
			return "%[" + strconv.Itoa(i+1) + "]s"
		},
	)
}

// MessageConverter - возвращает функцию для замены указанных аргументов в указанное сообщение.
func MessageConverter(leftDelim, rightDelim string) func(msg string, args []any) (newMsg string, newArgs []any) {
	formatter := NewMessageFormatter(leftDelim, rightDelim)

	return func(msg string, args []any) (newMsg string, newArgs []any) {
		newMsg, _ = formatter.Format(msg)

		return newMsg, args
	}
}
