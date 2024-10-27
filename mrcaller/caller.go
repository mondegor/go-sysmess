package mrcaller

import (
	"runtime"
)

const (
	defaultDepth        = 1
	defaultShowFuncName = false
	callStackMaxDepth   = 32
)

type (
	// Caller - обёртка runtime.Callers для более удобного формирования стека вызовов.
	Caller struct {
		depth                uint8                            // максимальное кол-во элементов в стеке вызовов
		showFuncName         bool                             // формирование имени функций в стеке вызовов
		filterStackTraceFunc func(frames []uintptr) []uintptr // функция фильтрации стека вызовов
	}
)

// New - создаёт объект Caller.
func New(opts ...CallerOption) *Caller {
	c := &Caller{
		depth:        defaultDepth,
		showFuncName: defaultShowFuncName,
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
