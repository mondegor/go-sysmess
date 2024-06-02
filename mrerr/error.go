package mrerr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/mrmsg"
)

const (
	// ErrorCodeInternal - обобщённый код ошибки: внутренняя ошибка приложения.
	ErrorCodeInternal = "errInternal"

	// ErrorCodeSystem - обобщённый код ошибки: системная ошибка приложения.
	ErrorCodeSystem = "errSystem"
)

type (
	// AppError - ошибка с поддержкой параметров, ID экземпляра ошибки и стека вызовов.
	AppError struct {
		code       string
		kind       ErrorKind
		instanceID string
		message    string
		argsNames  []string
		args       []any
		attrs      []mrmsg.NamedArg
		err        error
		stackTrace StackTracer
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

	newErr.setErrorIfArgsNotEqual()

	return newErr
}

func (e *AppError) setErrorIfArgsNotEqual() {
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
		code:    ErrorCodeInternal,
		kind:    ErrorKindInternal,
		message: fmt.Sprintf(errMessage, e.message),
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

	if e.stackTrace != nil {
		for i := 0; i < e.stackTrace.Count(); i++ {
			file, line := e.stackTrace.FileLine(i)

			if i == 0 {
				buf.WriteString(" in ")
			} else {
				buf.Write([]byte{' ', ',', ' '})
			}

			buf.WriteString(file)
			buf.WriteByte(':')
			buf.WriteString(strconv.Itoa(line))
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
