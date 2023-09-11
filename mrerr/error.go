package mrerr

import (
    "fmt"

    "github.com/mondegor/go-sysmess/mrmsg"
)

const (
    ErrorIdInternal = "errInternal"

    ErrorKindInternal ErrorKind = iota
    ErrorKindInternalNotice
    ErrorKindSystem
    ErrorKindUser
)

type (
    ErrorKind int

    AppError struct {
        id string
        kind ErrorKind
        traceId *string
        message string
        argsNames []string
        args []any
        err error
        file *string
        line int
    }
)

func New(id string, message string, args ...any) *AppError {
    newErr := &AppError{
        id: id,
        kind: ErrorKindUser,
        message: message,
        argsNames: mrmsg.ParseArgsNames(message),
        args: args,
    }

    newErr.setErrorIfArgsNotEqual(1)

    return newErr
}

func (e *AppError) setErrorIfArgsNotEqual(callerSkip int) {
    if len(e.argsNames) == len(e.args) {
        return
    }

    var errMessage string

    if len(e.argsNames) > len(e.args) {
        errMessage = "not enough arguments in message '%s'"
    } else {
        errMessage = "too many arguments in message '%s'"
    }

    argsErrorFactory := AppErrorFactory{
        id: ErrorIdInternal,
        kind: ErrorKindInternal,
        message: fmt.Sprintf(errMessage, e.message),
        callerSkip: callerSkip,
    }

    e.err = argsErrorFactory.new(e.err, nil)
}

func (e *AppError) Id() string {
    return e.id
}

func (e *AppError) Kind() ErrorKind {
    return e.kind
}

func (e *AppError) TraceId() string {
    if e.traceId == nil {
        return ""
    }

    return *e.traceId
}

func (e *AppError) Error() string {
    var buf []byte

    if e.traceId != nil {
        buf = append(buf, '[')
        buf = append(buf, *e.traceId...)
        buf = append(buf, ']', ' ')
    }

    buf = append(buf, e.renderMessage()...)

    if e.file != nil {
        buf = append(buf, fmt.Sprintf(" in %s:%d", *e.file, e.line)...)
    }

    if e.err != nil {
        buf = append(buf, ';', ' ')
        buf = append(buf, e.err.Error()...)
    }

    return string(buf)
}

func (e *AppError) Is(err error) bool {
    if v, ok := err.(*AppError); ok && e.id == v.id {
        return true
    }

    return false
}

func (e *AppError) Unwrap() error {
    return e.err
}
