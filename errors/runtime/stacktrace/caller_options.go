package stacktrace

type (
	// Option - настройка объекта Caller.
	Option func(o *caller)
)

// WithDepth - устанавливает глубину стека для Caller.
func WithDepth(value int) Option {
	return func(o *caller) {
		o.depth = value
	}
}

// WithShowFunc - устанавливает вывод названий функций для Caller.
func WithShowFunc(value bool) Option {
	return func(o *caller) {
		o.showFuncName = value
	}
}

// WithFindBottomBoundFunc - устанавливает функцию распознавания границы стека вызовов.
func WithFindBottomBoundFunc(fn func(funcName string) bool) Option {
	return func(o *caller) {
		o.findBottomBoundFunc = fn
	}
}
