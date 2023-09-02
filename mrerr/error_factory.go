package mrerr

import (
    "errors"
    "fmt"
    "runtime"

    "github.com/google/uuid"
    "github.com/mondegor/go-sysmess/mrmsg"
)

const (
    eventIdPattern = "ER-%s"
)

type (
    AppErrorFactory interface {
        New(args ...any) AppError
        Wrap(err error, args ...any) AppError
        Caller(skip int) AppErrorFactory
        Is(err error) bool
    }

    AppError interface {
        Id() ErrorId
        Kind() ErrorKind
        Is(err error) bool
        Error() string
        Unwrap() error
        EventId() string
    }

    appErrorFactory struct {
        id ErrorId
        kind ErrorKind
        message string
        argsNames []string
        callerSkip int
    }
)

func NewFactory(id ErrorId, kind ErrorKind, message string) AppErrorFactory {
    return &appErrorFactory{
        id: id,
        kind: kind,
        message: message,
        argsNames: mrmsg.ParseArgsNames(message),
    }
}

func (e *appErrorFactory) New(args ...any) AppError {
    return e.new(nil, args)
}

func (e *appErrorFactory) Wrap(err error, args ...any) AppError {
    if err == nil {
        err = fmt.Errorf("error is nil, wrapping is not necessary")
    }

    return e.new(err, args)
}

func (e *appErrorFactory) Caller(skip int) AppErrorFactory {
    return &appErrorFactory{
        id: e.id,
        kind: e.kind,
        message: e.message,
        argsNames: e.argsNames,
        callerSkip: e.callerSkip + skip,
    }
}

// Is - see: appError::Is
func (e *appErrorFactory) Is(err error) bool {
    return errors.Is(err, &appError{id: e.id})
}

func (e *appErrorFactory) new(err error, args []any) *appError {
    newErr := &appError{
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

func (e *appErrorFactory) init(newErr *appError) {
    newErr.setErrorIfArgsNotEqual(3)

    if newErr.err != nil {
        appErr, ok := newErr.err.(*appError)

        // raising to the top
        if ok && appErr.eventId != nil {
            newErr.eventId = appErr.eventId
            appErr.eventId = nil
            return
        }
    }

    if e.kind != ErrorKindInternal && e.kind != ErrorKindSystem {
        return
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

    //if e.eventId == nil {
    //    e.eventId = (*string)(sentry.CaptureException(e))
    //}

    if newErr.eventId == nil {
        newErr.eventId = new(string)
        *newErr.eventId = fmt.Sprintf(eventIdPattern, uuid.New().String())
    }
}
