package mrerr

import (
    "fmt"

    "github.com/mondegor/go-sysmess/mrmsg"
)

const (
    ErrorKindInternal ErrorKind = iota
    ErrorKindInternalNotice
    ErrorKindSystem
    ErrorKindUser

    ErrorIdInternal = "errInternal"
)

type (
    ErrorId string
    ErrorKind int32

    appError struct {
        id ErrorId
        kind ErrorKind
        message string
        argsNames []string
        args []any
        err error
        file *string
        line int
        eventId *string
    }
)

func New(id ErrorId, message string, args ...any) AppError {
    newErr := &appError{
        id: id,
        kind: ErrorKindUser,
        message: message,
        argsNames: mrmsg.ParseArgsNames(message),
        args: args,
    }

    newErr.setErrorIfArgsNotEqual(1)

    return newErr
}

func (e *appError) setErrorIfArgsNotEqual(callerSkip int) {
    if len(e.argsNames) == len(e.args) {
        return
    }

    var errMessage string

    if len(e.argsNames) > len(e.args) {
        errMessage = "not enough arguments in message '%s'"
    } else {
        errMessage = "too many arguments in message '%s'"
    }

    argsErrorFactory := appErrorFactory{
        id: ErrorIdInternal,
        kind: ErrorKindInternal,
        message: fmt.Sprintf(errMessage, e.message),
        callerSkip: callerSkip,
    }

    e.err = argsErrorFactory.new(e.err, nil)
}

func (e *appError) Id() ErrorId {
    return e.id
}

func (e *appError) Kind() ErrorKind {
    return e.kind
}

func (e *appError) Is(err error) bool {
    if v, ok := err.(*appError); ok && e.id == v.id {
        return true
    }

    return false
}

func (e *appError) Error() string {
    var buf []byte

    if e.eventId != nil {
        buf = append(buf, '[')
        buf = append(buf, *e.eventId...)
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

func (e *appError) Unwrap() error {
    return e.err
}

func (e *appError) EventId() string {
    if e.eventId == nil {
        return ""
    }

    return *e.eventId
}
