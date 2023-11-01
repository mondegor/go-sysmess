package mrerr

import (
    "github.com/mondegor/go-sysmess/mrlang"
    "github.com/mondegor/go-sysmess/mrmsg"
)

// Translate - translate error message for user
func (e *AppError) Translate(locale *mrlang.Locale) mrlang.ErrorMessage {
    if e.kind == ErrorKindInternal || e.kind == ErrorKindInternalNotice {
        return locale.TranslateError(ErrorIdInternal, ErrorIdInternal)
    }

    return locale.TranslateError(e.id, e.message, e.getNamedArgs()...)
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
        namedArgs[i] = mrmsg.NewArg(argName, e.args[i])
    }

    return namedArgs
}
