package stacktrace

type (
	// Option - функция-опция для настройки Caller.
	Option func(o *caller)
)

// WithDepth - устанавливает максимальную глубину (количество кадров) стека вызовов.
// Значение должно быть в диапазоне [1, 32]. Вне диапазона используется defaultDepth = 1.
func WithDepth(value int) Option {
	return func(o *caller) {
		o.depth = value
	}
}

// WithShowFunc - включает/отключает отображение имён функций в стеке вызовов.
func WithShowFunc(value bool) Option {
	return func(o *caller) {
		o.showFuncName = value
	}
}

// WithFindBottomBoundFunc - устанавливает функцию определения нижней границы стека вызовов.
// Функция принимает имя функции и возвращает true, если дальше этого кадра
// идут только вызовы общих пакетов и библиотек, не представляющие интереса.
func WithFindBottomBoundFunc(fn func(funcName string) bool) Option {
	return func(o *caller) {
		o.findBottomBoundFunc = fn
	}
}
