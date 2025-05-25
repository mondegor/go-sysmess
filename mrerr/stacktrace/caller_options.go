package stacktrace

type (
	// Option - настройка объекта Caller.
	Option func(c *Caller)
)

// WithDepth - устанавливает глубину стека для Caller.
func WithDepth(value uint8) Option {
	return func(c *Caller) {
		if value > 0 && value <= callStackMaxDepth {
			c.depth = value
		}
	}
}

// WithStackTraceFilter - устанавливает функцию фильтрации стека вызовов.
func WithStackTraceFilter(fn func(frames []uintptr) []uintptr) Option {
	return func(c *Caller) {
		if fn != nil {
			c.filterStackTraceFunc = fn
		}
	}
}
