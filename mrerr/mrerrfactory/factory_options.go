package mrerrfactory

import (
	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/features"
)

// NewProtoAppError - создаёт объект ProtoAppError с указанными опциями.
// При этом для генерации ID и стека используются функции из пакета features.
func NewProtoAppError(code string, kind mrerr.ErrorKind, message string, withIDGenerator, withCaller bool) *mrerr.ProtoAppError {
	generateIDFunc := features.GenerateInstanceID
	callerFunc := func() mrerr.StackTracer {
		return features.NewStackTrace()
	}

	if !withIDGenerator {
		generateIDFunc = nil
	}

	if !withCaller {
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
