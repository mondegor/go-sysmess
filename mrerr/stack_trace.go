package mrerr

import (
	"fmt"
	"runtime"
)

const (
	stackTraceMaxDepth = 4
)

type (
	stackTrace struct {
		pcs []uintptr
	}
)

func newStackTrace() *stackTrace {
	var pcs [stackTraceMaxDepth]uintptr
	n := runtime.Callers(6, pcs[:])

	if n > stackTraceMaxDepth {
		n = stackTraceMaxDepth
	}

	return &stackTrace{
		pcs: pcs[0:n],
	}
}

// Count - возвращается количество элементов в стеке вызовов.
func (s *stackTrace) Count() int {
	return len(s.pcs)
}

// FileLine - возвращает путь к файлу и номер строки кода,
// где расположена вызванная функция указанного элемента.
func (s *stackTrace) FileLine(i int) (file string, line int) {
	if i < 0 || i >= len(s.pcs) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, len(s.pcs)))
	}

	fn := runtime.FuncForPC(s.pcs[i] - 1)
	if fn == nil {
		return "", 0
	}

	return fn.FileLine(s.pcs[i] - 1)
}
