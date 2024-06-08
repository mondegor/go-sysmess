package factory

import (
	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/features"
)

type (
	// Options - опции для создания AppErrorProto
	Options struct {
		Code            string
		Kind            mrerr.ErrorKind
		Message         string
		WithIDGenerator bool
		WithCaller      bool
	}
)

// NewAppErrorProto - создаётся объект AppErrorProto с указанными опциями.
func NewAppErrorProto(opts Options) *mrerr.AppErrorProto {
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
