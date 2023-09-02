package mrerr

import (
    "github.com/mondegor/go-sysmess/mrlang"
    "github.com/mondegor/go-sysmess/mrmsg"
)

// Translate - translate error message for user
func (e *appError) Translate(loc mrlang.Locale) mrlang.ErrorMessage {
    if e.kind != ErrorKindInternal {
        return loc.TranslateError(string(e.Id()), e.message, e.getNamedArgs()...)
    }

    return loc.TranslateError(ErrorIdInternal, ErrorIdInternal)
}

func (e *appError) renderMessage() []byte {
    if len(e.argsNames) == 0 || len(e.argsNames) != len(e.args) {
        return []byte(e.message)
    }

    message, err := mrmsg.Render(e.message, e.getNamedArgs())

    if err != nil {
        return []byte(e.message)
    }

    return []byte(message)
}

func (e *appError) getNamedArgs() []mrmsg.NamedArg {
    var namedArgs []mrmsg.NamedArg

    for i, argName := range e.argsNames {
        namedArgs = append(namedArgs, mrmsg.NewArg(argName, e.args[i]))
    }

    return namedArgs
}
