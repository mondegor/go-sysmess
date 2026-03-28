package wrap

type (
	// ErrorWrapper - помощник для оборачивания ошибок.
	ErrorWrapper interface {
		Wrap(err error, attrs ...any) error
	}

	shellErrorWrapper struct {
		wrapFunc       func(err error, attrs []any) (wrappedErr error, ok bool)
		defaultWrapper ErrorWrapper
	}
)

// NewShellErrorWrapper - создаёт объект ErrorWrapper.
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

// Wrap - анализирует указанную ошибку, при необходимости её преобразует/оборачивает и возвращает результат.
// Сначала преобразование происходит с помощью wrapFunc, если попытка не удалась,
// то ошибка оборачивается в defaultWrapper.
func (w *shellErrorWrapper) Wrap(err error, attrs ...any) error {
	if wrappedErr, ok := w.wrapFunc(err, attrs); ok {
		return wrappedErr
	}

	return w.defaultWrapper.Wrap(err, attrs...)
}

type (
	nopErrorWrapper struct{}
)

// NopErrorWrapper - создаёт объект ErrorWrapper, который возвращает переданную ему ошибку как есть.
func NopErrorWrapper() ErrorWrapper {
	return nopErrorWrapper{}
}

// Wrap - возвращает указанную ошибку, реализуя ErrorWrapper интерфейс.
func (t nopErrorWrapper) Wrap(err error, _ ...any) error { return err }
