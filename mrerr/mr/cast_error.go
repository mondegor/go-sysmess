package mr

import (
	"github.com/mondegor/go-sysmess/mrerrors"
)

// CastOrWrapUnexpectedInternal - приводит указанную ошибку к InstantError или если
// это невозможно, то оборачивает ошибку в ErrUnexpectedInternal, затем возвращает результат.
func CastOrWrapUnexpectedInternal(err error) *mrerrors.InstantError {
	return mrerrors.CastOrWrap(
		err,
		func(err error) *mrerrors.InstantError {
			return ErrUnexpectedInternal.Wrap(err)
		},
	)
}
