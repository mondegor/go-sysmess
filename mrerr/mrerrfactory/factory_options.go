package mrerrfactory

import (
	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/features"
)

// NewProtoAppError - создаёт объект ProtoAppError с указанными опциями.
// При этом генерации стека вызовов и генерации ID используются функции из пакета features.
func NewProtoAppError(code string, kind mrerr.ErrorKind, message string, withCaller, withIDGenerator bool) *mrerr.ProtoAppError {
	callerFunc := func() mrerr.StackTracer {
		return features.NewStackTrace()
	}

	onCreateFunc := func(_ *mrerr.AppError) (instanceID string) {
		return features.GenerateInstanceID()
	}

	if !withCaller {
		callerFunc = nil
	}

	if !withIDGenerator {
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
