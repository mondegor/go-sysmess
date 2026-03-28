package wrap

import (
	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// kindlessErrorWrapper - помощник для оборачивания ошибок,
	// которые не реализуют interface{ Kind() kind.Enum }.
	kindlessErrorWrapper struct {
		defaultWrapper ErrorWrapper
	}
)

// NewKindlessErrorWrapper - создаёт объект ErrorWrapper.
func NewKindlessErrorWrapper(
	defaultWrapper ErrorWrapper,
) ErrorWrapper {
	if defaultWrapper == nil {
		defaultWrapper = nopWrapper{}
	}

	return &kindlessErrorWrapper{
		defaultWrapper: defaultWrapper,
	}
}

// Wrap - возвращает указанную ошибку как есть, если она реализует метод Kind(),
// иначе её оборачивает в defaultWrapper и возвращает результат.
// Если ошибка типа kind.Internal и указаны атрибуты, то она
// дополнительно будет обёрнута в defaultWrapper.
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
