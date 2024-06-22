package features

import (
	"fmt"
	"runtime"
)

const (
	stackTraceMaxDepth = 4
)

type (
	// StackTrace - стек вызов функций.
	StackTrace struct {
		pcs []uintptr
	}
)

// NewStackTrace - создаёт объект StackTrace.
func NewStackTrace() *StackTrace {
	var pcs [stackTraceMaxDepth]uintptr
	n := runtime.Callers(6, pcs[:])

	if n > stackTraceMaxDepth {
		n = stackTraceMaxDepth
	}

	return &StackTrace{
		pcs: pcs[0:n],
	}
}

// Count - возвращается количество элементов в стеке вызовов.
func (s *StackTrace) Count() int {
	return len(s.pcs)
}

// Item - возвращает имя функции указанного элемента, путь к файлу и номер строки кода.
// Если i превысит кол-во элементов в стеке вызовов, то будет вызвана panic.
func (s *StackTrace) Item(i int) (name, file string, line int) {
	if i < 0 || i >= len(s.pcs) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, len(s.pcs)))
	}

	fn := runtime.FuncForPC(s.pcs[i] - 1)
	if fn == nil {
		return "unknown", "???", 0
	}

	file, line = fn.FileLine(s.pcs[i] - 1)

	return fn.Name(), file, line
}
