package mrerr

import (
	"github.com/mondegor/go-sysmess/mrmsg"
)

// Translate - translate error message for user
func (e *AppError) Translate(t mrmsg.ErrorTranslator) mrmsg.ErrorMessage {
	if e.kind == ErrorKindInternal || e.kind == ErrorKindInternalNotice {
		if !t.CheckError(e.code) {
			return t.TranslateError(ErrorCodeInternal, ErrorCodeInternal)
		}
	} else if e.kind == ErrorKindSystem {
		if !t.CheckError(e.code) {
			return t.TranslateError(ErrorCodeSystem, ErrorCodeSystem)
		}
	}

	return t.TranslateError(e.code, e.message, e.getNamedArgs()...)
}

func (e *AppError) renderMessage() []byte {
	if len(e.argsNames) == 0 || len(e.argsNames) != len(e.args) {
		return []byte(e.message)
	}

	return []byte(mrmsg.Render(e.message, e.getNamedArgs()))
}

func (e *AppError) getNamedArgs() []mrmsg.NamedArg {
	namedArgs := make([]mrmsg.NamedArg, len(e.argsNames))

	for i, argName := range e.argsNames {
		namedArgs[i] = mrmsg.NamedArg{
			Name:  argName,
			Value: e.args[i],
		}
	}

	return namedArgs
}
