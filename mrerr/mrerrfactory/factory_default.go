package mrerrfactory

import (
	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/features"
)

// NewProtoAppErrorByDefault - создаёт стандартный прототип ошибки ProtoAppError,
// а если указан один из видов ErrorKindInternal или ErrorKindSystem,
// то к ним добавляется генерация ID и стека вызовов из пакета features.
func NewProtoAppErrorByDefault(code string, kind mrerr.ErrorKind, message string) *mrerr.ProtoAppError {
	generateIDFunc := features.GenerateInstanceID
	callerFunc := func() mrerr.StackTracer {
		return features.NewStackTrace()
	}

	if kind == mrerr.ErrorKindUser {
		generateIDFunc = nil
		callerFunc = nil
	}

	return mrerr.NewProtoWithExtra(
		code,
		kind,
		message,
		generateIDFunc,
		callerFunc,
	)
}
