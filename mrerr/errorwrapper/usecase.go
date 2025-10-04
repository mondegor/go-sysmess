package errorwrapper

import (
	"errors"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/mr"
)

type (
	// UseCase - помощник оборачивания перехваченных ошибок
	// в часто используемые ошибки бизнес-логики приложения.
	UseCase struct{}
)

// NewUseCase - создаёт объект UseCase.
func NewUseCase() *UseCase {
	return &UseCase{}
}

// IsNotFoundOrNotAffectedError - сообщает, связанна ли ошибка с отсутствием запрошенной записи,
// или она была найдена, но её изменение не потребовалось.
func (w *UseCase) IsNotFoundOrNotAffectedError(err error) bool {
	return errors.Is(err, mr.ErrStorageNoRowFound) ||
		errors.Is(err, mr.ErrStorageRowsNotAffected)
}

// WrapErrorFailed - возвращает ошибку с указанием источника, обёрнутую в
// ErrUseCaseTemporarilyUnavailable или ErrUseCaseOperationFailed.
// Ошибки ErrUseCaseOperationFailed, ErrUseCaseTemporarilyUnavailable и пользовательские ошибки не оборачиваются!
func (w *UseCase) WrapErrorFailed(err error, attrs ...any) error {
	return w.wrapErrorFailed(err, attrs)
}

// WrapErrorNotFoundOrFailed - возвращает ошибку с указанием источника, обёрнутую в
// ErrUseCaseEntityNotFound, ErrUseCaseTemporarilyUnavailable или ErrUseCaseOperationFailed.
// Ошибки ErrUseCaseOperationFailed, ErrUseCaseTemporarilyUnavailable и пользовательские ошибки не оборачиваются!
func (w *UseCase) WrapErrorNotFoundOrFailed(err error, attrs ...any) error {
	if w.IsNotFoundOrNotAffectedError(err) {
		return mr.ErrUseCaseEntityNotFound.New()
	}

	return w.wrapErrorFailed(err, attrs)
}

func (w *UseCase) wrapErrorFailed(err error, attrs []any) error {
	if mr.ErrUseCaseOperationFailed.Is(err) {
		return err
	}

	if e, ok := err.(interface{ Kind() mrerr.ErrorKind }); ok {
		if e.Kind() == mrerr.ErrorKindSystem {
			if mr.ErrUseCaseTemporarilyUnavailable.Is(err) {
				return err
			}

			return mr.ErrUseCaseTemporarilyUnavailable.Wrap(err, attrs...)
		}

		if e.Kind() == mrerr.ErrorKindUser {
			return err
		}
	}

	return mr.ErrUseCaseOperationFailed.Wrap(err, attrs...)
}
