package errors

import (
	"github.com/mondegor/go-sysmess/errors/wrap"
)

type (
	// Wrapper - помощник для оборачивания ошибок.
	Wrapper = wrap.ErrorWrapper
)

// NewInfraStorageWrapper - используется в инфраструктурном слое
// для оборачивания нераспознанных ошибок в ErrInternalStorageQueryFailed.
func NewInfraStorageWrapper() Wrapper {
	return wrap.NewShellErrorWrapper(
		func(err error, _ []any) (wrappedErr error, ok bool) {
			return err, Is(err, ErrEventStorageNoRecordFound)
		},
		wrap.NewKindlessErrorWrapper(
			ErrInternalStorageQueryFailed,
		),
	)
}

// NewServiceOperationFailedWrapper - используется в сервисном слое
// для оборачивания нераспознанных ошибок в ErrInternalServiceOperationFailed.
func NewServiceOperationFailedWrapper() Wrapper {
	return wrap.NewKindlessErrorWrapper(
		ErrInternalServiceOperationFailed,
	)
}

// NewServiceRecordNotFoundWrapper - используется в сервисном слое
// для оборачивания ошибок ErrEventStorageNoRecordFound в ErrRecordNotFound.
// Все остальные нераспознанные ошибки оборачиваются в ErrInternalServiceOperationFailed.
func NewServiceRecordNotFoundWrapper() Wrapper {
	return wrap.NewShellErrorWrapper(
		func(err error, _ []any) (wrappedErr error, ok bool) {
			if Is(err, ErrEventStorageNoRecordFound) {
				return ErrRecordNotFound, true
			}

			return err, Is(err, ErrRecordNotFound)
		},
		wrap.NewKindlessErrorWrapper(
			ErrInternalServiceOperationFailed,
		),
	)
}

// NewServiceRecordVersionConflictWrapper - используется в сервисном слое
// для оборачивания ошибок ErrEventStorageNoRecordFound в ErrRecordVersionConflict.
// Все остальные нераспознанные ошибки оборачиваются в ErrInternalServiceOperationFailed.
func NewServiceRecordVersionConflictWrapper() Wrapper {
	return wrap.NewShellErrorWrapper(
		func(err error, _ []any) (wrappedErr error, ok bool) {
			if Is(err, ErrEventStorageNoRecordFound) {
				return ErrRecordVersionConflict, true
			}

			return err, false
		},
		wrap.NewKindlessErrorWrapper(
			ErrInternalServiceOperationFailed,
		),
	)
}
