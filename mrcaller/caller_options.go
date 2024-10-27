package mrcaller

type (
	// CallerOption - настройка объекта Caller.
	CallerOption func(c *Caller)
)

// WithDepth - устанавливает опцию depth для Caller.
func WithDepth(value uint8) CallerOption {
	return func(c *Caller) {
		if value > 0 && value <= callStackMaxDepth {
			c.depth = value
		}
	}
}

// WithShowFuncName - устанавливает опцию noname для Caller.
func WithShowFuncName(value bool) CallerOption {
	return func(c *Caller) {
		c.showFuncName = value
	}
}

// WithFilterStackTrace - функцию фильтрации стека вызовов.
func WithFilterStackTrace(fn func(frames []uintptr) []uintptr) CallerOption {
	return func(c *Caller) {
		if fn != nil {
			c.filterStackTraceFunc = fn
		}
	}
}
