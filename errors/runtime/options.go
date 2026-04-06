package runtime

import (
	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// Option - функция-опция для настройки прототипа runtime-ошибки (protoError).
	Option func(o *protoError)
)

// WithOnCreate - устанавливает обработчик, вызываемый при создании каждого экземпляра ошибки.
// Обработчик получает тип ошибки и обёрнутую ошибку, возвращает дополнительные данные
// (обычно hint.Hint с ID ошибки и стеком вызовов), которая ассоциируется с экземпляром ошибки.
func WithOnCreate(value func(kindErr kind.Enum, wrappedErr error) (hint any)) Option {
	return func(o *protoError) {
		o.onCreate = value
	}
}
