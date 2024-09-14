package mrcaller

type (
	// CallerOption - настройка объекта Caller.
	CallerOption func(c *Caller)
)

// WithDepth - устанавливает опцию depth для Caller.
// Если value меньше 1 или больше callStackMaxDepth, то устанавливается 1.
func WithDepth(value int) CallerOption {
	return func(c *Caller) {
		c.depth = value
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
		c.filterStackTraceFunc = fn
	}
}
