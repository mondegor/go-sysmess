package errorwrapper

import (
	"github.com/mondegor/go-sysmess/mrerr/mr"
	"github.com/mondegor/go-sysmess/mrerrors"
)

type (
	// InfraStorage - помощник для оборачивания перехваченных ошибок
	// в ошибки инфраструктурного слоя приложения при работе с хранилищем данных (БД).
	InfraStorage struct {
		wrapper *mrerrors.Wrapper
	}
)

// NewInfraStorage - создаёт объект InfraStorage.
func NewInfraStorage() *InfraStorage {
	w, err := mrerrors.NewWrapper(
		mr.ErrStorageQueryFailed,
		nil,
		mr.ErrStorageNoRowFound,
	)
	if err != nil {
		panic(err)
	}

	return &InfraStorage{
		wrapper: w,
	}
}

// WrapError - возвращает ошибку, обёрнутую в mr.ErrStorageQueryFailed.
// Пользовательские ошибки и ошибка mr.ErrStorageNoRowFound не оборачиваются.
// Если указанная ошибка совпадает с mr.ErrStorageQueryFailed, то она только дополняется атрибутами.
func (w *InfraStorage) WrapError(err error, attrs ...any) error {
	return w.wrapper.WrapError(err, attrs...) //nolint:wrapcheck
}
