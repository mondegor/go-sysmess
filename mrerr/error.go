package mrerr

import (
	"fmt"
	"strings"

	"github.com/mondegor/go-sysmess/mrmsg"
)

const (
	ErrorInternalID = "errInternal"
)

type (
	AppError struct {
		id        string
		kind      ErrorKind
		traceID   string
		message   string
		argsNames []string
		args      []any
		err       error
		file      string
		line      int
	}
)

func New(id, message string, args ...any) *AppError {
	newErr := &AppError{
		id:        id,
		kind:      ErrorKindUser,
		message:   message,
		argsNames: mrmsg.ParseArgsNames(message),
		args:      args,
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
		id:         ErrorInternalID,
		kind:       ErrorKindInternal,
		message:    fmt.Sprintf(errMessage, e.message),
		callerSkip: callerSkip,
	}

	e.err = argsErrorFactory.new(e.err, nil)
}

func (e *AppError) ID() string {
	return e.id
}

func (e *AppError) Kind() ErrorKind {
	return e.kind
}

func (e *AppError) TraceID() string {
	return e.traceID
}

func (e *AppError) Error() string {
	var buf strings.Builder

	if e.traceID != "" {
		buf.WriteByte('[')
		buf.WriteString(e.traceID)
		buf.Write([]byte{']', ' '})
	}

	// buf.WriteString(e.id)
	// buf.Write([]byte{':', ' '})

	buf.Write(e.renderMessage())

	if e.file != "" {
		buf.WriteString(fmt.Sprintf(" in %s:%d", e.file, e.line))
	}

	if e.err != nil {
		buf.Write([]byte{';', ' '})
		buf.WriteString(e.err.Error())
	}

	return buf.String()
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
