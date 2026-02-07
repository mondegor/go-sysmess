package stacktrace

type (
	// Option - настройка объекта Caller.
	Option func(o *caller)
)

// WithDepth - устанавливает глубину стека для Caller.
func WithDepth(value uint8) Option {
	return func(o *caller) {
		o.depth = value
	}
}

// WithStackTraceFilter - устанавливает функцию фильтрации стека вызовов.
func WithStackTraceFilter(fn func(frames []uintptr) []uintptr) Option {
	return func(o *caller) {
		o.filterStackTraceFunc = fn
	}
}
