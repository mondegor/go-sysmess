package mrerr

import (
	"github.com/mondegor/go-sysmess/mrmsg"
)

// WithoutStackTrace - возвращает ошибку без стека вызовов, если он был сформирован.
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

func makeArgs(args []any, argsNames []string) []any {
	if len(args) >= len(argsNames) {
		return args
	}

	l := len(args)
	newArgs := make([]any, len(argsNames))

	// копируются все переданные параметры в новый массив
	for i := 0; i < l; i++ {
		newArgs[i] = args[i]
	}

	// копируются недостающие параметры
	for i := l; i < len(argsNames); i++ {
		newArgs[i] = "missed-error-arg={{ ." + argsNames[i] + " }}"
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
