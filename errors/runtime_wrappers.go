package errors

import (
	"github.com/mondegor/go-sysmess/errors/helper"
)

//go:generate mockgen -source=runtime_wrappers.go -destination=./mock/runtime_wrappers.go

type (
	// Wrapper - помощник для оборачивания ошибок.
	Wrapper interface {
		Wrap(err error, attrs ...any) error
	}
)

// NewRuntimeWrapper - оборачивает ошибки в runtime ошибки.
// Ошибки совпадающие с wrapperInternal, wrapperSystem, пользовательские ошибки
// и ошибки из exceptions не оборачиваются и переданные атрибуты для них игнорируются.
func NewRuntimeWrapper(
	wrapperInternal RuntimeProtoError,
	wrapperSystem RuntimeProtoError,
	prewrapFunc func(err error, attrs []any) (wrappedErr error, ok bool),
) Wrapper {
	w, err := helper.NewRuntimeErrorWrapper(
		wrapperInternal,
		wrapperSystem,
		prewrapFunc,
	)
	if err != nil {
		panic(err)
	}

	return w
}

// NewInfraStorageWrapper - используется в инфраструктурном слое
// для оборачивания ошибок в ErrInternalStorageQueryFailed.
// Пользовательские ошибки и ErrEventStorageNoRowFound не оборачиваются.
// ErrInternalStorageQueryFailed повторно также не оборачивается.
func NewInfraStorageWrapper() Wrapper {
	return NewRuntimeWrapper(
		ErrInternalStorageQueryFailed,
		nil,
		func(err error, _ []any) (wrappedErr error, ok bool) {
			return err, Is(err, ErrEventStorageNoRowFound)
		},
	)
}

// NewServiceWrapper - используется в сервисном слое
// для оборачивания ошибок в ErrInternalServiceOperationFailed.
// Пользовательские ошибки и ErrEventStorageNoRowFound, ErrInternalStorageQueryFailed не оборачиваются.
// ErrInternalServiceOperationFailed повторно также не оборачивается.
func NewServiceWrapper() Wrapper {
	return NewRuntimeWrapper(
		ErrInternalServiceOperationFailed,
		nil,
		func(err error, _ []any) (wrappedErr error, ok bool) {
			return err, Is(err, ErrEventStorageNoRowFound) || Is(err, ErrInternalStorageQueryFailed)
		},
	)
}

// NewUseCaseWrapper - используется в слое бизнес логики
// для оборачивания ошибок в ErrInternalUseCaseOperationFailed и ErrSystemUseCaseTemporarilyUnavailable.
// Пользовательские ошибки не оборачиваются. ErrEventStorageNoRowFound оборачивается в ErrUseCaseEntityNotFound.
// ErrInternalUseCaseOperationFailed, ErrSystemUseCaseTemporarilyUnavailable,
// ErrSystemUseCaseTemporarilyUnavailable повторно также не оборачиваются.
func NewUseCaseWrapper() Wrapper {
	return NewRuntimeWrapper(
		ErrInternalUseCaseOperationFailed,
		ErrSystemUseCaseTemporarilyUnavailable,
		func(err error, _ []any) (wrappedErr error, ok bool) {
			if Is(err, ErrUseCaseEntityNotFound) {
				return err, true
			}

			if Is(err, ErrEventStorageNoRowFound) {
				return ErrUseCaseEntityNotFound.Wrap(err), true
			}

			return err, false
		},
	)
}
