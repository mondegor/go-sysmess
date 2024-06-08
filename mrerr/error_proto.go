package mrerr

import (
	"errors"

	"github.com/mondegor/go-sysmess/mrmsg"
)

type (
	// AppErrorProto - ошибка с поддержкой параметров, ID экземпляра ошибки и стека вызовов.
	AppErrorProto struct {
		pureError
		generateID func() string
		caller     func() StackTracer
	}

	// StackTracer - предоставляет доступ к стеку вызовов.
	StackTracer interface {
		Count() int
		FileLine(index int) (file string, line int)
	}
)

// вспомогательная ошибка, чтобы отметить, что Wrap применяется для Nil ошибки
var errSpecifiedErrorIsNil = errors.New("[WARNING!!! specified error is nil, wrapping is not necessary]")

// NewProto - создаёт объект AppErrorProto.
func NewProto(code string, kind ErrorKind, message string) *AppErrorProto {
	argsNames := mrmsg.ParseArgsNames(message)

	return &AppErrorProto{
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

func NewProtoWithExtra(code string, kind ErrorKind, message string, generateID func() string, caller func() StackTracer) *AppErrorProto {
	proto := NewProto(code, kind, message)
	proto.generateID = generateID
	proto.caller = caller

	return proto
}

// New - всегда создаёт новую копию текущего объекта,
// при этом вызываются функции generateID и stackTrace.caller
func (e *AppErrorProto) New(args ...any) *AppError {
	c := &AppError{
		pureError: e.pureError,
	}

	if len(args) > 0 {
		// если аргументов передано больше c.argsNames,
		// то при вызове Error() ошибки будет выведено предупреждение
		c.args = makeArgs(args, len(c.argsNames))
	}

	if e.generateID != nil {
		c.instanceID = e.generateID()
	}

	if e.caller != nil {
		c.stackTrace.val = e.caller()
		c.stackTrace.has = true
	}

	return c
}

func (e *AppErrorProto) Wrap(err error, args ...any) *AppError {
	if err == nil {
		err = errSpecifiedErrorIsNil
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

	// WARNING: c.err должна быть именно типа *AppErrorProto, а не вложенные в неё ошибки
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

	if e.generateID != nil && c.errInstanceID == nil {
		c.instanceID = e.generateID()
	}

	if e.caller != nil && !c.stackTrace.has {
		c.stackTrace.val = e.caller()
		c.stackTrace.has = true
	}

	return c
}

// Error - возвращает ошибку в виде строки.
func (e *AppErrorProto) Error() string {
	return mrmsg.Render(e.message, e.getNamedArgs())
}
