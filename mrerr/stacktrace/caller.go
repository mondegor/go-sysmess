package stacktrace

import (
	"runtime"
)

const (
	defaultDepth      = 1
	callStackMaxDepth = 32
)

type (
	// Caller - обёртка runtime.Callers для более удобного формирования стека вызовов.
	Caller struct {
		depth                uint8                            // максимальное кол-во элементов в стеке вызовов
		filterStackTraceFunc func(frames []uintptr) []uintptr // функция фильтрации стека вызовов
	}
)

// New - создаёт объект Caller.
func New(opts ...Option) *Caller {
	c := &Caller{
		depth: defaultDepth,
		filterStackTraceFunc: func(frames []uintptr) []uintptr {
			return frames
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// StackTrace - формирует стек вызовов для текущего вызова функции и возвращает его.
func (c *Caller) StackTrace() *StackTrace {
	var pcs [callStackMaxDepth]uintptr
	n := runtime.Callers(2, pcs[:])
	fpcs := c.filterStackTraceFunc(pcs[0:n])

	if len(fpcs) > int(c.depth) {
		fpcs = fpcs[0:c.depth]
	}

	items := make([]source, len(fpcs))

	for i := 0; i < len(fpcs); i++ {
		function, file, line := runtimeFrame(fpcs[i]).Source()

		items[i] = source{
			Function: function,
			File:     file,
			Line:     line,
		}
	}

	return &StackTrace{
		sources: items,
	}
}
