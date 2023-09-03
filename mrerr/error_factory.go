package mrerr

import (
    "errors"
    "fmt"
    "runtime"

    "github.com/google/uuid"
    "github.com/mondegor/go-sysmess/mrmsg"
)

const (
    traceIdPattern = "ER-%s"
)

type (
    AppErrorFactory struct {
        id string
        kind ErrorKind
        message string
        argsNames []string
        callerSkip int
    }
)

func NewFactory(id string, kind ErrorKind, message string) *AppErrorFactory {
    return &AppErrorFactory{
        id: id,
        kind: kind,
        message: message,
        argsNames: mrmsg.ParseArgsNames(message),
    }
}

func (e *AppErrorFactory) Caller(skip int) *AppErrorFactory {
    return &AppErrorFactory{
        id: e.id,
        kind: e.kind,
        message: e.message,
        argsNames: e.argsNames,
        callerSkip: e.callerSkip + skip,
    }
}

func (e *AppErrorFactory) New(args ...any) *AppError {
    return e.new(nil, args)
}

func (e *AppErrorFactory) Wrap(err error, args ...any) *AppError {
    if err == nil {
        err = fmt.Errorf("specified error is nil, wrapping is not necessary")
    }

    return e.new(err, args)
}

// Is - see: AppError::Is
func (e *AppErrorFactory) Is(err error) bool {
    return errors.Is(err, &AppError{id: e.id})
}

func (e *AppErrorFactory) new(err error, args []any) *AppError {
    newErr := &AppError{
        id: e.id,
        kind: e.kind,
        message: e.message,
        argsNames: e.argsNames,
        args: args,
        err: err,
    }

    e.init(newErr)

    return newErr
}

func (e *AppErrorFactory) init(newErr *AppError) {
    newErr.setErrorIfArgsNotEqual(3)

    if newErr.err != nil {
        appErr, ok := newErr.err.(*AppError)

        // raising to the top
        if ok && appErr.traceId != nil {
            newErr.traceId = appErr.traceId
            appErr.traceId = nil
            return
        }
    }

    if e.kind != ErrorKindInternal && e.kind != ErrorKindSystem {
        return
    }

    if newErr.traceId == nil {
        newErr.traceId = new(string)
        *newErr.traceId = fmt.Sprintf(traceIdPattern, uuid.New().String())
    }

    _, file, line, ok := runtime.Caller(e.callerSkip + 3)

    if ok {
        if file == "" {
            file = "???"
        }

        newErr.file = new(string)
        *newErr.file = file
        newErr.line = line
    }
}
