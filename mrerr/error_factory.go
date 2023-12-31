package mrerr

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/mondegor/go-sysmess/mrmsg"
)

type (
	AppErrorFactory struct {
		id         string
		kind       ErrorKind
		message    string
		argsNames  []string
		callerSkip int
	}
)

func NewFactory(id string, kind ErrorKind, message string) *AppErrorFactory {
	return &AppErrorFactory{
		id:         id,
		kind:       kind,
		message:    message,
		argsNames:  mrmsg.ParseArgsNames(message),
		callerSkip: 3, // skip: New, new, init
	}
}

func (e *AppErrorFactory) Caller(skip int) *AppErrorFactory {
	if skip == 0 {
		return e
	}

	c := *e
	c.callerSkip += skip

	return &c
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

func (e *AppErrorFactory) ErrorID() string {
	return e.id
}

// Is - see: AppError::Is
func (e *AppErrorFactory) Is(err error) bool {
	return errors.Is(err, &AppError{id: e.id})
}

func (e *AppErrorFactory) new(err error, args []any) *AppError {
	newErr := &AppError{
		id:        e.id,
		kind:      e.kind,
		message:   e.message,
		argsNames: e.argsNames,
		args:      args,
		err:       err,
	}

	e.init(newErr)

	return newErr
}

func (e *AppErrorFactory) init(newErr *AppError) {
	newErr.setErrorIfArgsNotEqual(3)

	if newErr.err != nil {
		appErr, ok := newErr.err.(*AppError)

		// raising to the top
		if ok && appErr.traceID != "" {
			newErr.traceID = appErr.traceID
			appErr.traceID = ""
			return
		}
	}

	if e.kind != ErrorKindInternal && e.kind != ErrorKindSystem {
		return
	}

	if newErr.traceID == "" {
		newErr.traceID = e.generateTraceID()
	}

	newErr.callStack = caller.CallStack(e.callerSkip)
}

// 'hex(unix time)' - 'hex(4 rand bytes)' -> 64e9c0f1-1e97228f
func (e *AppErrorFactory) generateTraceID() string {
	value := make([]byte, 4)
	_, err := rand.Read(value)

	if err != nil {
		value = []byte{0x0, 0xee, 0xee, 0x0}
	}

	return strconv.FormatInt(time.Now().Unix(), 16) + "-" + hex.EncodeToString(value)
}
