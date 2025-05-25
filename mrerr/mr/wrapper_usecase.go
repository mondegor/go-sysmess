package mr

import (
	"errors"

	"github.com/mondegor/go-sysmess/mrerr"
)

type (
	// UseCaseErrorWrapper - помощник оборачивания перехваченных ошибок
	// в часто используемые ошибки бизнес-логики приложения.
	UseCaseErrorWrapper struct {
		attrs []any // атрибуты должны быть указаны попарно: название/значение
	}
)

// NewUseCaseErrorWrapper - создаёт объект UseCaseErrorWrapper.
func NewUseCaseErrorWrapper(source string) *UseCaseErrorWrapper {
	return &UseCaseErrorWrapper{
		attrs: []any{"usecase-source", source},
	}
}

// WithAttrs - возвращает новый UseCaseErrorWrapper с прикреплёнными атрибутами.
func (w *UseCaseErrorWrapper) WithAttrs(attrs ...any) *UseCaseErrorWrapper {
	c := *w
	c.attrs = append(c.attrs, attrs...)

	return &c
}

// IsNotFoundOrNotAffectedError - сообщает, связанна ли ошибка с отсутствием запрошенной записи,
// или её изменение не потребовалось.
func (w *UseCaseErrorWrapper) IsNotFoundOrNotAffectedError(err error) bool {
	return errors.Is(err, ErrStorageNoRowFound) ||
		errors.Is(err, ErrStorageRowsNotAffected)
}

// WrapErrorFailed - возвращает ошибку с указанием источника, обёрнутую в
// ErrUseCaseTemporarilyUnavailable или ErrUseCaseOperationFailed.
// Ошибки ErrUseCaseOperationFailed, ErrUseCaseTemporarilyUnavailable и пользовательские ошибки не оборачиваются!
func (w *UseCaseErrorWrapper) WrapErrorFailed(err error, attrs ...any) error {
	return w.wrapErrorFailed(err, attrs)
}

// WrapErrorNotFoundOrFailed - возвращает ошибку с указанием источника, обёрнутую в
// ErrUseCaseEntityNotFound, ErrUseCaseTemporarilyUnavailable или ErrUseCaseOperationFailed.
// Ошибки ErrUseCaseOperationFailed, ErrUseCaseTemporarilyUnavailable и пользовательские ошибки не оборачиваются!
func (w *UseCaseErrorWrapper) WrapErrorNotFoundOrFailed(err error, attrs ...any) error {
	if w.IsNotFoundOrNotAffectedError(err) {
		return ErrUseCaseEntityNotFound.New()
	}

	return w.wrapErrorFailed(err, attrs)
}

func (w *UseCaseErrorWrapper) wrapErrorFailed(err error, attrs []any) error {
	if ErrUseCaseOperationFailed.Is(err) {
		return err
	}

	if e, ok := err.(interface{ Kind() mrerr.ErrorKind }); ok {
		if e.Kind() == mrerr.ErrorKindSystem {
			if ErrUseCaseTemporarilyUnavailable.Is(err) {
				return err
			}

			return ErrUseCaseTemporarilyUnavailable.Wrap(err, w.attrs...).WithAttrs(attrs...)
		}

		if e.Kind() == mrerr.ErrorKindUser {
			return err
		}
	}

	return ErrUseCaseOperationFailed.Wrap(err, w.attrs...).WithAttrs(attrs...)
}
