package mrerr

import (
	"errors"
	"fmt"

	"github.com/mondegor/go-sysmess/mrcaller"
	"github.com/mondegor/go-sysmess/mrmsg"
)

const (
	attrNameByDefault = "unnamed"
)

type (
	// AppErrorFactory - фабрика ошибок, с поддержкой типов, параметров,
	// формирования CallStack, с возможностью wrap ошибок.
	AppErrorFactory struct {
		code            string
		kind            ErrorKind
		message         string
		argsNames       []string
		attrs           []mrmsg.NamedArg
		generateID      func() string
		caller          func(skip int) mrcaller.CallStack
		callerSkipFrame int
	}
)

// NewFactory - создаётся объект AppErrorFactory.
func NewFactory(code string, etype ErrorType, message string) *AppErrorFactory {
	return &AppErrorFactory{
		code:       code,
		kind:       etype.Kind,
		message:    message,
		argsNames:  mrmsg.ParseArgsNames(message),
		generateID: etype.GenerateIDFunc,
		caller:     etype.CallerFunc,
	}
}

// WithCallerSkipFrame - возвращает AppErrorFactory с установленным
// пропуском функций, которые не должны попасть в CallStack.
func (e *AppErrorFactory) WithCallerSkipFrame(skip int) *AppErrorFactory {
	c := *e
	c.callerSkipFrame += skip

	return &c
}

// WithAttr - возвращает AppErrorFactory с прикреплённым к нему именованным параметром.
func (e *AppErrorFactory) WithAttr(name string, value any) *AppErrorFactory {
	if name == "" {
		name = attrNameByDefault
	}

	c := *e
	c.attrs = append(
		c.attrs,
		mrmsg.NamedArg{
			Name:  name,
			Value: value,
		},
	)

	return &c
}

// New - создаётся объект AppError с использованием параметров фабрики.
func (e *AppErrorFactory) New(args ...any) *AppError {
	return e.new(nil, args)
}

// Wrap - возвращает ошибку с вложенной в неё указанной ошибки.
func (e *AppErrorFactory) Wrap(err error, args ...any) *AppError {
	if err == nil {
		err = fmt.Errorf("specified error is nil, wrapping is not necessary")
	}

	return e.new(err, args)
}

// Code - возвращает код ошибки.
func (e *AppErrorFactory) Code() string {
	return e.code
}

// Is - see: AppError::Is
func (e *AppErrorFactory) Is(err error) bool {
	return errors.Is(err, &AppError{code: e.code})
}

func (e *AppErrorFactory) new(err error, args []any) *AppError {
	newErr := &AppError{
		code:      e.code,
		kind:      e.kind,
		message:   e.message,
		argsNames: e.argsNames,
		args:      args,
		attrs:     e.attrs,
		err:       err,
	}

	e.init(newErr)

	return newErr
}

func (e *AppErrorFactory) init(newErr *AppError) {
	const skipFrame = 3
	newErr.setErrorIfArgsNotEqual(skipFrame)

	hasInstanceID := false
	hasCallStack := false

	if newErr.err != nil {
		if wrappedErr, ok := newErr.err.(*AppError); ok {
			// instanceID is raising to the top
			if wrappedErr.instanceID != "" {
				newErr.instanceID = wrappedErr.instanceID
				wrappedErr.instanceID = ""
				hasInstanceID = true
			}

			if !wrappedErr.callStack.Empty() {
				hasCallStack = true
			}
		}
	}

	if e.generateID != nil && !hasInstanceID {
		newErr.instanceID = e.generateID()
	}

	if e.caller != nil && !hasCallStack {
		newErr.callStack = e.caller(e.callerSkipFrame)
	}
}
