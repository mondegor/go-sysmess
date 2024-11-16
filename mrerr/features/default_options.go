package features

import (
	"github.com/mondegor/go-sysmess/mrerr"
)

// WithProtoCaller - возвращается опция формирования стека вызовов.
func WithProtoCaller() mrerr.ProtoOption {
	return mrerr.WithProtoCaller(
		func() mrerr.StackTracer {
			return NewStackTrace()
		},
	)
}

// WithProtoOnCreated - возвращается опция с обработчиком создания экземпляра ошибки.
func WithProtoOnCreated() mrerr.ProtoOption {
	return mrerr.WithProtoOnCreated(
		func(_ *mrerr.AppError) (instanceID string) {
			return GenerateInstanceID()
		},
	)
}

// DefaultOptionsHandler - возвращается обработчик формирования опций по умолчанию,
// используемый при создании Proto ошибок.
func DefaultOptionsHandler() mrerr.ProtoOptionsHandlerFunc {
	options := []mrerr.ProtoOption{
		WithProtoCaller(),
		WithProtoOnCreated(),
	}

	return func(_ string, kind mrerr.ErrorKind) []mrerr.ProtoOption {
		if kind == mrerr.ErrorKindUser {
			return nil
		}

		return options
	}
}
