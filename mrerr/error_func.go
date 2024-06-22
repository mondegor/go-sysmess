package mrerr

import (
	"strconv"

	"github.com/mondegor/go-sysmess/mrmsg"
)

// WithExtra - создаёт новую ProtoAppError на основе указанной с дополнительными параметрами.
func WithExtra(proto ProtoAppError, extra ProtoExtra) ProtoAppError {
	proto.caller = extra.Caller
	proto.onCreated = extra.OnCreated

	return proto
}

// WithoutStackTrace - возвращает ошибку без стека вызовов.
func WithoutStackTrace(err *AppError) *AppError {
	if err == nil {
		return ErrErrorIsNilPointer.New()
	}

	if err.stackTrace.val == nil {
		return err
	}

	c := *err
	c.stackTrace.val = nil

	return &c
}

// Cast - преобразует в ошибку AppError без вызова generateID и caller.
func Cast(proto *ProtoAppError) *AppError {
	if proto == nil {
		return ErrErrorIsNilPointer.New()
	}

	return &AppError{
		pureError: proto.pureError,
	}
}

func makeArgs(args []any, minLength int) []any {
	if len(args) >= minLength {
		return args
	}

	l := len(args)
	newArgs := make([]any, minLength)

	// копируются все переданные параметры в новый массив
	for i := 0; i < l; i++ {
		newArgs[i] = args[i]
	}

	// копируются недостающие параметры
	for i := l; i < minLength; i++ {
		newArgs[i] = "missed-arg" + strconv.Itoa(i+1)
	}

	return newArgs
}

func appendAttr(attrs []mrmsg.NamedArg, name string, value any) []mrmsg.NamedArg {
	if name == "" {
		name = attrNameByDefault
	}

	attrs = append(
		attrs,
		mrmsg.NamedArg{
			Name:  name,
			Value: value,
		},
	)

	return attrs
}
