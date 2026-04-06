package wrap

type (
	// ErrorWrapper - помощник для обёртывания ошибок.
	// Позволяет настраивать стратегию обёртывания через композицию функций.
	ErrorWrapper interface {
		Wrap(err error, attrs ...any) error
	}

	shellErrorWrapper struct {
		wrapFunc       func(err error, attrs []any) (wrappedErr error, ok bool)
		defaultWrapper ErrorWrapper
	}
)

// NewShellErrorWrapper - создаёт обёртку с двумя уровнями обработки.
// Параметр wrapFunc - основная функция обёртывания; если она возвращает ok=true,
// результат сразу возвращается. Если ok=false, используется defaultWrapper.
func NewShellErrorWrapper(
	wrapFunc func(err error, attrs []any) (wrappedErr error, ok bool),
	defaultWrapper ErrorWrapper,
) ErrorWrapper {
	if wrapFunc == nil {
		wrapFunc = func(err error, _ []any) (wrappedErr error, ok bool) {
			return err, false
		}
	}

	if defaultWrapper == nil {
		defaultWrapper = nopErrorWrapper{}
	}

	return &shellErrorWrapper{
		wrapFunc:       wrapFunc,
		defaultWrapper: defaultWrapper,
	}
}

// Wrap - анализирует ошибку и пытается обернуть её через wrapFunc.
// Если wrapFunc вернула ok=false, ошибка оборачивается через defaultWrapper.
func (w *shellErrorWrapper) Wrap(err error, attrs ...any) error {
	if wrappedErr, ok := w.wrapFunc(err, attrs); ok {
		return wrappedErr
	}

	return w.defaultWrapper.Wrap(err, attrs...)
}

type (
	// nopErrorWrapper - заглушка, реализующая интерфейс ErrorWrapper.
	// Возвращает переданную ошибку без изменений.
	nopErrorWrapper struct{}
)

// NopErrorWrapper - создаёт ErrorWrapper-заглушку.
func NopErrorWrapper() ErrorWrapper {
	return nopErrorWrapper{}
}

// Wrap - возвращает указанную ошибку, реализуя ErrorWrapper интерфейс.
func (t nopErrorWrapper) Wrap(err error, _ ...any) error { return err }
