package mrerr

import (
	"github.com/mondegor/go-sysmess/mrmsg"
)

//go:generate mockgen -source=error_proto.go -destination=./mock/error_proto.go

type (
	// ProtoAppError - прототип ошибки с поддержкой параметров, ID экземпляра ошибки и стека вызовов.
	ProtoAppError struct {
		pureError
		caller    func() StackTracer
		onCreated func(err *AppError) (instanceID string)
	}

	// ProtoExtra - дополнительные опции для создания ProtoAppError.
	ProtoExtra struct {
		Caller    func() StackTracer
		OnCreated func(err *AppError) (instanceID string)
	}

	// StackTracer - предоставляет доступ к стеку вызовов.
	StackTracer interface {
		Count() int
		Item(index int) (name, file string, line int)
	}
)

// ErrErrorIsNilPointer - указанная ошибка - nil pointer.
var ErrErrorIsNilPointer = NewProto(
	"errErrorIsNilPointer", ErrorKindInternal, "specified error is nil")

// NewProto - создаёт объект ProtoAppError.
func NewProto(code string, kind ErrorKind, message string) *ProtoAppError {
	argsNames := mrmsg.ParseArgsNames(message)

	return &ProtoAppError{
		pureError: pureError{
			code:      code,
			kind:      kind,
			message:   message,
			argsNames: argsNames,

			// параметризованные сообщения по умолчанию
			// используют фиктивные значения параметров
			args: makeArgs(nil, len(argsNames)),
		},
	}
}

// NewProtoWithExtra - создаёт объект ProtoAppError с дополнительными параметрами.
func NewProtoWithExtra(code string, kind ErrorKind, message string, extra ProtoExtra) *ProtoAppError {
	proto := NewProto(code, kind, message)
	proto.caller = extra.Caller
	proto.onCreated = extra.OnCreated

	return proto
}

// New - всегда создаёт новую копию текущего объекта,
// при этом вызываются функции caller и onCreated если они были установлены.
func (e *ProtoAppError) New(args ...any) *AppError {
	c := &AppError{
		pureError: e.pureError,
	}

	if len(args) > 0 {
		// если аргументов передано больше c.argsNames,
		// то при вызове Error() ошибки будет выведено предупреждение
		c.args = makeArgs(args, len(c.argsNames))
	}

	if e.caller != nil {
		c.stackTrace.val = e.caller()
		c.stackTrace.has = true
	}

	if e.onCreated != nil {
		c.instanceID = e.onCreated(c)
	}

	return c
}

// Wrap - создаёт новую ошибку на основе прототипа и оборачивает
// в неё указанную. Если указанная ошибка типа AppError, то проверяется
// был ли у этой ошибки сгенерированы ID и стек, и если да, то
// у новой ошибки эти параметры не генерятся, даже если соответствующие
// генераторы установлены для этой ошибки.
func (e *ProtoAppError) Wrap(err error, args ...any) *AppError {
	if err == nil {
		err = ErrErrorIsNilPointer.New()
	}

	c := &AppError{
		pureError: e.pureError,
		err:       err,
	}

	if len(args) > 0 {
		// если аргументов передано больше c.argsNames,
		// то при вызове Error() ошибки будет выведено предупреждение
		c.args = makeArgs(args, len(c.argsNames))
	}

	if wrappedErr, ok := c.err.(*AppError); ok { //nolint:errorlint
		// instanceID is raising to the top
		if wrappedErr.errInstanceID != nil {
			c.errInstanceID = wrappedErr.errInstanceID
		} else if wrappedErr.instanceID != "" {
			c.errInstanceID = &wrappedErr.instanceID
		}

		// если стек был установлен во вложенном объекте,
		// то запрещается генерация стека текущему объекту
		if wrappedErr.stackTrace.has {
			c.stackTrace.has = true
		}
	}

	if e.caller != nil && !c.stackTrace.has {
		c.stackTrace.val = e.caller()
		c.stackTrace.has = true
	}

	if e.onCreated != nil && c.errInstanceID == nil {
		c.instanceID = e.onCreated(c)
	}

	return c
}

// Error - возвращает ошибку в виде строки.
func (e *ProtoAppError) Error() string {
	return mrmsg.MustRenderWithNamedArgs(e.message, e.getNamedArgs())
}
