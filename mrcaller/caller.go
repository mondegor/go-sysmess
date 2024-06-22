package mrcaller

import (
	"runtime"
)

const (
	callStackMaxDepth = 32
)

type (
	// Caller - обёртка runtime.Callers для более удобного формирования стека вызовов.
	Caller struct {
		depth                int                              // максимальное кол-во элементов в стеке вызовов
		showFuncName         bool                             // формирование имени функций в стеке вызовов
		filterStackTraceFunc func(frames []uintptr) []uintptr // функция фильтрации стека вызовов
	}
)

// New - создаёт объект Caller.
func New(opts ...CallerOption) *Caller {
	c := &Caller{}
	c.applyOptions(opts)

	const minDepth = 1
	if c.depth < minDepth {
		c.depth = minDepth
	} else if c.depth > callStackMaxDepth {
		c.depth = callStackMaxDepth
	}

	if c.filterStackTraceFunc == nil {
		c.filterStackTraceFunc = func(frames []uintptr) []uintptr {
			return frames
		}
	}

	return c
}

// StackTrace - формирует стек вызовов для текущего вызова функции и возвращает его.
func (c *Caller) StackTrace() *StackTrace {
	var pcs [callStackMaxDepth]uintptr
	n := runtime.Callers(2, pcs[:])
	fpcs := c.filterStackTraceFunc(pcs[0:n])

	if len(fpcs) > c.depth {
		fpcs = fpcs[0:c.depth]
	}

	items := make([]StackItem, len(fpcs))

	for i := 0; i < len(fpcs); i++ {
		frame := runtimeFrame(fpcs[i])
		file, line := frame.FileLine()

		items[i] = StackItem{
			File: file,
			Line: line,
		}

		if c.showFuncName {
			items[i].Name = frame.Name()
		}
	}

	return &StackTrace{
		items: items,
	}
}

func (c *Caller) applyOptions(opts []CallerOption) {
	for _, f := range opts {
		f(c)
	}
}
