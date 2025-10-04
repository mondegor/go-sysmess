package errorwrapper

import (
	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/mr"
)

type (
	// InfraStorage - помощник для оборачивания перехваченных ошибок
	// в часто используемые ошибки инфраструктурного слоя приложения при работе с БД.
	InfraStorage struct {
		attrs []any // атрибуты должны быть указаны попарно: название/значение
	}
)

// NewInfraStorage - создаёт объект InfraStorage.
func NewInfraStorage(source string) *InfraStorage {
	return &InfraStorage{
		attrs: []any{mrerr.ErrorSourceKey, source},
	}
}

// WithAttrs - возвращает новый InfraStorage с прикреплёнными атрибутами.
func (w *InfraStorage) WithAttrs(attrs ...any) *InfraStorage {
	c := *w
	c.attrs = append(c.attrs, attrs...)

	return &c
}

// WrapError - возвращает ошибку с указанием источника.
// Если ошибка не mrerrors.InstantError или не mrerrors.ProtoError, то она оборачивается в ErrStorageQueryFailed.
// Ошибки ErrStorageNoRowFound, ErrStorageRowsNotAffected, ErrStorageQueryFailed и пользовательские ошибки не оборачиваются!
func (w *InfraStorage) WrapError(err error, attrs ...any) error {
	if mr.ErrStorageNoRowFound.Is(err) || mr.ErrStorageRowsNotAffected.Is(err) || mr.ErrStorageQueryFailed.Is(err) {
		return err
	}

	if e, ok := err.(interface{ Kind() mrerr.ErrorKind }); ok && e.Kind() == mrerr.ErrorKindUser {
		return err
	}

	return mr.ErrStorageQueryFailed.Wrap(err, w.attrs...).WithAttrs(attrs...)
}
