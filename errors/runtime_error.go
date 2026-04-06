package errors

import (
	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/runtime"
)

type (
	// RuntimeProtoError - прототип внутренних (kind.Internal) и системных (kind.System) ошибок.
	RuntimeProtoError = runtime.ProtoError
)

// NewInternalProto - создаёт прототип ошибки типа kind.Internal.
func NewInternalProto(text string) RuntimeProtoError {
	return runtime.NewDelayed(kind.Internal, text)
}

// NewSystemProto - создаёт прототип ошибки типа kind.System.
func NewSystemProto(text string) RuntimeProtoError {
	return runtime.NewDelayed(kind.System, text)
}

// errInternal - internal error,
// обобщённая внутренняя ошибка системы которая может быть решена только силами разработки.
// Для неё всегда должен формироваться стек вызовов и ID ошибки.
var errInternal = NewInternalProto("internal error")

// NewInternalError - создаёт ошибку типа kind.Internal с указанным описанием и атрибутами.
// Параметр details - текстовое пояснение, attrs - пары ключ(string)/значение(any) для логирования.
func NewInternalError(details string, attrs ...any) error {
	return errInternal.WithDetails(details, attrs...) //nolint:wrapcheck
}

// WrapInternalError - создаёт ошибку типа kind.Internal, обёртывающую указанную ошибку.
// Параметры:
//   - err - исходная ошибка для обёртывания;
//   - details - текстовое пояснение;
//   - attrs - пары ключ/значение для логирования.
func WrapInternalError(err error, details string, attrs ...any) error {
	return errInternal.WithError(err, details, attrs...) //nolint:wrapcheck
}
