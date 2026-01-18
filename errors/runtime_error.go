package errors

import (
	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/runtime"
)

type (
	// RuntimeProtoError - внутренняя/системная ошибка.
	RuntimeProtoError = runtime.ProtoError
)

// NewInternalProto - создаёт RuntimeProtoError для создания ошибок типа kind.Internal с предопределёнными опциями.
func NewInternalProto(text string) RuntimeProtoError {
	return runtime.NewDelayed(kind.Internal, text)
}

// NewSystemProto - создаёт RuntimeProtoError для создания ошибок типа kind.System с предопределёнными опциями.
func NewSystemProto(text string) RuntimeProtoError {
	return runtime.NewDelayed(kind.System, text)
}

// errInternal - internal error,
// обобщённая внутренняя ошибка системы которая может быть решена только силами разработки.
// Для неё всегда должен формироваться стек вызовов и ID ошибки.
var errInternal = NewInternalProto("internal error")

// NewInternalError - создаёт ошибку типа kind.Internal с предопределёнными опциями.
func NewInternalError(details string, attrs ...any) error {
	return errInternal.WithDetails(details, attrs...) //nolint:wrapcheck
}

// WrapInternalError - создаёт ошибку типа kind.Internal с предопределёнными опциями.
func WrapInternalError(err error, details string, attrs ...any) error {
	return errInternal.WithError(err, details, attrs...) //nolint:wrapcheck
}
