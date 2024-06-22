package mrerrfactory

import (
	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/features"
)

// NewProtoAppErrorByDefault - создаёт стандартный прототип ошибки ProtoAppError,
// а если указан один из видов ErrorKindInternal или ErrorKindSystem,
// то к ним добавляется генерация стека вызовов и генерация ID из пакета features.
func NewProtoAppErrorByDefault(code string, kind mrerr.ErrorKind, message string) *mrerr.ProtoAppError {
	callerFunc := func() mrerr.StackTracer {
		return features.NewStackTrace()
	}

	onCreateFunc := func(_ *mrerr.AppError) (instanceID string) {
		return features.GenerateInstanceID()
	}

	if kind == mrerr.ErrorKindUser {
		callerFunc = nil
		onCreateFunc = nil
	}

	return mrerr.NewProtoWithExtra(
		code,
		kind,
		message,
		mrerr.ProtoExtra{
			Caller:    callerFunc,
			OnCreated: onCreateFunc,
		},
	)
}
