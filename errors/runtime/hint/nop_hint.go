package hint

import (
	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	nopHint struct{}
)

// Nop - создаёт объект Hint, который ничего не делает.
func Nop() Hint {
	return nopHint{}
}

// ErrorKind - возвращает неопределённый тип ошибки.
func (h nopHint) ErrorKind() kind.Enum {
	return 0
}

// ErrorID - возвращает пустой ID ошибки.
func (h nopHint) ErrorID() string {
	return ""
}

// StackTraceIterator - возвращает итератор на пустой стек вызовов.
func (h nopHint) StackTraceIterator() func() (index int, name, file string, line int) {
	return func() (index int, name, file string, line int) {
		return -1, "", "", 0
	}
}
