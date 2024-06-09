package factory

import (
	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/features"
)

type (
	// Options - опции для создания ProtoAppError.
	Options struct {
		Code            string
		Kind            mrerr.ErrorKind
		Message         string
		WithIDGenerator bool
		WithCaller      bool
	}
)

// NewProtoAppError - создаёт объект ProtoAppError с указанными опциями.
// При этом для генерации ID и стека используются функции из пакета features.
func NewProtoAppError(opts Options) *mrerr.ProtoAppError {
	generateIDFunc := features.GenerateInstanceID
	callerFunc := func() mrerr.StackTracer {
		return features.NewStackTrace()
	}

	if !opts.WithIDGenerator {
		generateIDFunc = nil
	}

	if !opts.WithCaller {
		callerFunc = nil
	}

	return mrerr.NewProtoWithExtra(
		opts.Code,
		opts.Kind,
		opts.Message,
		generateIDFunc,
		callerFunc,
	)
}
