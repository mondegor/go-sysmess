package stacktrace

import (
	"runtime"
)

const (
	defaultDepth       = 1
	stackTraceMaxDepth = 32
)

type (
	// Caller - обёртка runtime.Callers для более удобного формирования стека вызовов.
	Caller interface {
		Call() StackTrace
	}

	// StackTrace - объект с уже сформированным стеком вызовов функций.
	StackTrace interface {
		Iterator() func() (index int, name, file string, line int)
	}

	caller struct {
		depth                uint8                            // максимальное кол-во элементов в стеке вызовов
		filterStackTraceFunc func(frames []uintptr) []uintptr // функция фильтрации стека вызовов
	}
)

// NewCaller - создаёт объект Caller.
func NewCaller(opts ...Option) Caller {
	c := &caller{}

	for _, opt := range opts {
		opt(c)
	}

	if c.depth < 1 || c.depth > stackTraceMaxDepth {
		c.depth = defaultDepth
	}

	if c.filterStackTraceFunc == nil {
		c.filterStackTraceFunc = func(frames []uintptr) []uintptr {
			return frames
		}
	}

	return c
}

// Call - формирует стек вызовов для текущего вызова функции и возвращает его.
func (c *caller) Call() StackTrace {
	var pcs [stackTraceMaxDepth]uintptr
	n := runtime.Callers(2, pcs[:])
	fpcs := c.filterStackTraceFunc(pcs[0:n])

	if len(fpcs) > int(c.depth) {
		fpcs = fpcs[0:c.depth]
	}

	items := make([]source, len(fpcs))

	for i := 0; i < len(fpcs); i++ {
		function, file, line := runtimeFrame(fpcs[i]).Source()

		items[i] = source{
			function: function,
			file:     file,
			line:     line,
		}
	}

	return stackTrace(items)
}
