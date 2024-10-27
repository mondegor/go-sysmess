package mrerr

import (
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/mrmsg"
)

const (
	attrNameByDefault       = "unnamed"
	messageTooManyArguments = "[WARNING!!! too many arguments in error message] "
)

type (
	// AppError - ошибка с поддержкой параметров, ID экземпляра ошибки и стека вызовов.
	AppError struct {
		pureError
		instanceID    string // собственный ID (устанавливается только если не установлено во вложенной ошибке)
		attrs         []mrmsg.NamedArg
		err           error
		errInstanceID *string // ID вложенной ошибки
		stackTrace    stackTrace
	}

	stackTrace struct {
		val StackTracer
		has bool // признак, что стек есть у текущего объекта или у одного из вложенных
	}
)

// WithAttr - возвращает новую ошибку с прикреплённым к нему именованным атрибутом.
func (e *AppError) WithAttr(name string, value any) *AppError {
	c := *e
	c.attrs = appendAttr(c.attrs, name, value)

	return &c
}

// InstanceID - возвращает уникальный идентификатор случившейся ошибки.
// Но только если в фабрике, породившей эту ошибку, был установлен генератор ID ошибок,
// В противном случае вернётся пустая строка.
func (e *AppError) InstanceID() string {
	if e.errInstanceID != nil {
		return *e.errInstanceID
	}

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

	if len(e.argsNames) == 0 {
		buf.WriteString(e.message)
	} else {
		if len(e.args) > len(e.argsNames) {
			buf.WriteString(messageTooManyArguments)
		}

		buf.WriteString(mrmsg.MustRenderWithNamedArgs(e.message, e.getNamedArgs()))
	}

	if len(e.attrs) > 0 {
		buf.Write([]byte{' ', '('})

		for i, attr := range e.attrs {
			if i > 0 {
				buf.Write([]byte{',', ' '})
			}

			buf.WriteString(attr.Name)
			buf.Write([]byte{'='})
			buf.WriteString(attr.ValueString())
		}

		buf.WriteByte(')')
	}

	if e.stackTrace.val != nil {
		cnt := e.stackTrace.val.Count()

		for i := 0; i < cnt; i++ {
			name, file, line := e.stackTrace.val.Item(i)

			if i == 0 {
				buf.WriteString(" in ")
			} else {
				buf.Write([]byte{' ', ',', ' '})
			}

			if name != "" {
				buf.WriteByte('[')
				buf.WriteString(name)
				buf.Write([]byte{']', ' '})
			}

			buf.WriteString(file)
			buf.WriteByte(':')
			buf.WriteString(strconv.Itoa(line))
		}
	}

	if e.err != nil {
		buf.Write([]byte{':', ' '})
		buf.WriteString(e.err.Error())
	}

	return buf.String()
}

// Unwrap - возвращает вложенную ошибку (errors.Is использует этот интерфейс).
func (e *AppError) Unwrap() error {
	return e.err
}
