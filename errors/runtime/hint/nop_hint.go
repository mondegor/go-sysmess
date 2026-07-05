package hint

import (
	"github.com/mondegor/go-core/errors/kind"
)

type (
	// nopHint - заглушка, реализующая интерфейс Hint.
	// Возвращает пустые/нулевые значения для всех методов.
	nopHint struct{}
)

// Nop - создаёт Hint-заглушку, возвращающую пустые значения.
// Используется когда у ошибки нет дополнительных данных (Hint).
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
