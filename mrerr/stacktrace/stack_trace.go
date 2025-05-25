package stacktrace

import (
	"fmt"
)

type (
	// StackTrace - объект с уже сформированным стеком вызовов функций.
	StackTrace struct {
		sources []source
	}

	// source - элемент объекта StackTrace с вызовом конкретной функции.
	source struct {
		Function string
		File     string
		Line     int
	}
)

// Count - возвращается количество элементов в стеке вызовов.
func (t *StackTrace) Count() int {
	return len(t.sources)
}

// Source - возвращает имя функции указанного элемента, путь к файлу и номер строки кода.
// Если i превысит кол-во элементов в стеке вызовов, то будет вызвана panic.
func (t *StackTrace) Source(i int) (name, file string, line int) {
	if i < 0 || i >= len(t.sources) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, len(t.sources)))
	}

	return t.sources[i].Function, t.sources[i].File, t.sources[i].Line
}
