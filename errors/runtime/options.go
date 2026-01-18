package runtime

import (
	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// Option - настройка объекта protoError.
	Option func(o *protoError)
)

// WithOnCreate - устанавливает обработчик события создания экземпляра ошибки.
func WithOnCreate(value func(kindErr kind.Enum, wrappedErr error) (bag any)) Option {
	return func(o *protoError) {
		o.onCreate = value
	}
}
