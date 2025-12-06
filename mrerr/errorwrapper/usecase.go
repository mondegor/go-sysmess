package errorwrapper

import (
	"errors"

	"github.com/mondegor/go-sysmess/mrerr/mr"
	"github.com/mondegor/go-sysmess/mrerrors"
)

type (
	// UseCase - помощник оборачивания перехваченных ошибок
	// в часто используемые ошибки бизнес-логики приложения.
	UseCase struct {
		wrapper *mrerrors.Wrapper
	}
)

// NewUseCase - создаёт объект UseCase.
func NewUseCase() *UseCase {
	w, err := mrerrors.NewWrapper(
		mr.ErrUseCaseOperationFailed,
		mr.ErrUseCaseTemporarilyUnavailable,
	)
	if err != nil {
		panic(err)
	}

	return &UseCase{
		wrapper: w,
	}
}

// IsNotFoundError - сообщает, связанна ли ошибка с отсутствием запрошенной записи.
func (w *UseCase) IsNotFoundError(err error) bool {
	return errors.Is(err, mr.ErrStorageNoRowFound) || errors.Is(err, mr.ErrUseCaseEntityNotFound)
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
	if errors.Is(err, mr.ErrUseCaseEntityNotFound) {
		return err
	}

	if errors.Is(err, mr.ErrStorageNoRowFound) {
		return mr.ErrUseCaseEntityNotFound.Wrap(err)
	}

	return w.wrapErrorFailed(err, attrs)
}

func (w *UseCase) wrapErrorFailed(err error, attrs []any) error {
	return w.wrapper.WrapError(err, attrs) //nolint:wrapcheck
}
