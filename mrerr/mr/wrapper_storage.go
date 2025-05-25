package mr

import (
	"github.com/mondegor/go-sysmess/mrerr"
)

type (
	// StorageErrorWrapper - помощник для оборачивания перехваченных ошибок
	// в часто используемые ошибки инфраструктурного слоя приложения.
	StorageErrorWrapper struct {
		attrs []any // атрибуты должны быть указаны попарно: название/значение
	}
)

// NewStorageErrorWrapper - создаёт объект StorageErrorWrapper.
func NewStorageErrorWrapper(source string) *StorageErrorWrapper {
	return &StorageErrorWrapper{
		attrs: []any{"storage-source", source},
	}
}

// WithAttrs - возвращает новый StorageErrorWrapper с прикреплёнными атрибутами.
func (w *StorageErrorWrapper) WithAttrs(attrs ...any) *StorageErrorWrapper {
	c := *w
	c.attrs = append(c.attrs, attrs...)

	return &c
}

// WrapError - возвращает ошибку с указанием источника.
// Если ошибка не mrerrors.InstantError или не mrerrors.ProtoError, то она оборачивается в ErrStorageQueryFailed.
// Ошибки ErrStorageNoRowFound, ErrStorageRowsNotAffected, ErrStorageQueryFailed и пользовательские ошибки не оборачиваются!
func (w *StorageErrorWrapper) WrapError(err error, attrs ...any) error {
	if ErrStorageNoRowFound.Is(err) || ErrStorageRowsNotAffected.Is(err) || ErrStorageQueryFailed.Is(err) {
		return err
	}

	if e, ok := err.(interface{ Kind() mrerr.ErrorKind }); ok && e.Kind() == mrerr.ErrorKindUser {
		return err
	}

	return ErrStorageQueryFailed.Wrap(err, w.attrs...).WithAttrs(attrs...)
}
