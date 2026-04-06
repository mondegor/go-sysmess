package wrap

import (
	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// kindlessErrorWrapper - помощник для оборачивания ошибок, не реализующих интерфейс Kind().
	// Ошибки с известным типом (User, System, Internal) пропускаются,
	// а необработанные ошибки оборачиваются через defaultWrapper.
	kindlessErrorWrapper struct {
		defaultWrapper ErrorWrapper
	}
)

// NewKindlessErrorWrapper - создаёт объект ErrorWrapper.
func NewKindlessErrorWrapper(
	defaultWrapper ErrorWrapper,
) ErrorWrapper {
	if defaultWrapper == nil {
		defaultWrapper = nopErrorWrapper{}
	}

	return &kindlessErrorWrapper{
		defaultWrapper: defaultWrapper,
	}
}

// Wrap - возвращает ошибки с известным типом (User, System, Internal без атрибутов) как есть.
// Ошибки без Kind() и Internal с атрибутами оборачиваются через defaultWrapper.
// Алгоритм:
//   - kind.User, kind.System - возвращаются как есть.
//   - kind.Internal без атрибутов - возвращается как есть.
//   - kind.Internal с атрибутами - оборачивается в defaultWrapper.
//   - Без Kind() - оборачивается в defaultWrapper.
func (w *kindlessErrorWrapper) Wrap(err error, attrs ...any) error {
	switch kind.Extract(err) {
	case kind.User, kind.System:
		return err
	case kind.Internal:
		// если атрибутов не указано
		if len(attrs) == 0 {
			return err
		}
	default:
	}

	return w.defaultWrapper.Wrap(err, attrs...)
}
