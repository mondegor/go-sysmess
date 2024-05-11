package mrerr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/mrcaller"
	"github.com/mondegor/go-sysmess/mrmsg"
)

const (
	// ErrorCodeInternal - обобщённый код ошибки: внутренняя ошибка приложения.
	ErrorCodeInternal = "errInternal"

	// ErrorCodeSystem - обобщённый код ошибки: системная ошибка приложения.
	ErrorCodeSystem = "errSystem"
)

type (
	// AppError - ошибка с поддержкой параметров и CallStack.
	AppError struct {
		code       string
		kind       ErrorKind
		instanceID string
		message    string
		argsNames  []string
		args       []any
		attrs      []mrmsg.NamedArg
		err        error
		callStack  mrcaller.CallStack
	}
)

// New - создаётся объект AppError.
func New(code, message string, args ...any) *AppError {
	newErr := &AppError{
		code:      code,
		kind:      ErrorKindUser,
		message:   message,
		argsNames: mrmsg.ParseArgsNames(message),
		args:      args,
	}

	const skipFrame = 1
	newErr.setErrorIfArgsNotEqual(skipFrame)

	return newErr
}

func (e *AppError) setErrorIfArgsNotEqual(callerSkipFrame int) {
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
		code:            ErrorCodeInternal,
		kind:            ErrorKindInternal,
		message:         fmt.Sprintf(errMessage, e.message),
		callerSkipFrame: callerSkipFrame,
	}

	e.err = argsErrorFactory.new(e.err, nil)
}

// Code - возвращает код ошибки.
func (e *AppError) Code() string {
	return e.code
}

// Kind - возвращает тип ошибки.
func (e *AppError) Kind() ErrorKind {
	return e.kind
}

// InstanceID - возвращает уникальный идентификатор случившейся ошибки.
// Но только если в фабрике, породившей эту ошибку, был установлен генератор ID ошибок.
func (e *AppError) InstanceID() string {
	return e.instanceID
}

// HasCallStack - возвращается, что в ошибке содержится CallStack (включая вложенные ошибки).
// Но только если в фабрике, породившей эту ошибку, было установлено формирование ID ошибок.
func (e *AppError) HasCallStack() bool {
	if !e.callStack.Empty() {
		return true
	}

	for err := e.err; err != nil; {
		if v, ok := err.(*AppError); ok {
			if !v.callStack.Empty() {
				return true
			}

			err = v.err
		}
	}

	return false
}

// Error - возвращает ошибку в виде строки.
func (e *AppError) Error() string {
	var buf strings.Builder

	if e.instanceID != "" {
		buf.WriteByte('[')
		buf.WriteString(e.instanceID)
		buf.Write([]byte{']', ' '})
	}

	// buf.WriteString(e.code)
	// buf.Write([]byte{':', ' '})

	buf.Write(e.renderMessage())

	if len(e.attrs) > 0 {
		buf.Write([]byte{' ', '('})

		for i, attr := range e.attrs {
			if i > 0 {
				buf.Write([]byte{',', ' '})
			}

			buf.WriteString(attr.Name)
			buf.Write([]byte{':', ' '})
			buf.WriteString(attr.ValueString())
		}

		buf.WriteByte(')')
	}

	for iter := e.callStack.NewIterator(); ; {
		if n, item := iter.Next(); n > 0 {
			if n == 1 {
				buf.WriteString(" in ")
			} else {
				buf.Write([]byte{' ', ',', ' '})
			}

			buf.WriteString(item.File())
			buf.WriteByte(':')
			buf.WriteString(strconv.Itoa(item.Line()))
		} else {
			break
		}
	}

	if e.err != nil {
		buf.Write([]byte{';', ' '})
		buf.WriteString(e.err.Error())
	}

	return buf.String()
}

// Is - проверяется что ошибка с указанным кодом (для возможности использования errors.Is).
func (e *AppError) Is(err error) bool {
	if v, ok := err.(*AppError); ok && e.code == v.code {
		return true
	}

	return false
}

// Unwrap - возвращает вложенную ошибку (для возможности использования errors.Is).
func (e *AppError) Unwrap() error {
	return e.err
}
