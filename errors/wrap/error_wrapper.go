package wrap

//go:generate mockgen -source=error_wrapper.go -destination=./mock/error_wrapper.go

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
		defaultWrapper = nopWrapper{}
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
	nopWrapper struct{}
)

// Wrap - возвращает указанную ошибку, реализуя wrapper интерфейс.
func (e nopWrapper) Wrap(err error, _ ...any) error {
	return err
}
