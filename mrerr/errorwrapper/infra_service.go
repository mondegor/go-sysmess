package errorwrapper

import (
	"github.com/mondegor/go-sysmess/mrerr/mr"
	"github.com/mondegor/go-sysmess/mrerrors"
)

type (
	// Service - помощник для оборачивания перехваченных ошибок в ошибки сервисного слоя приложения.
	Service struct {
		wrapper *mrerrors.Wrapper
	}
)

// NewService - создаёт объект Service.
func NewService() *Service {
	w, err := mrerrors.NewWrapper(
		mr.ErrServiceOperationFailed,
		nil,
		mr.ErrStorageNoRowFound,
		mr.ErrStorageQueryFailed,
	)
	if err != nil {
		panic(err)
	}

	return &Service{
		wrapper: w,
	}
}

// WrapError - возвращает ошибку, обёрнутую в mr.ErrServiceOperationFailed.
// Пользовательские ошибки и ошибки mr.ErrStorageNoRowFound, mr.ErrStorageQueryFailed не оборачиваются.
// Если указанная ошибка совпадает с mr.ErrServiceOperationFailed, то она только дополняется атрибутами.
func (w *Service) WrapError(err error, attrs ...any) error {
	return w.wrapper.WrapError(err, attrs...) //nolint:wrapcheck
}
