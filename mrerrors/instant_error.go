package mrerrors

import (
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/mrargs"
)

type (
	// InstantError - ошибка с поддержкой параметров, ID экземпляра ошибки и стека вызовов.
	InstantError struct {
		*pureError
		id    *string // уникальный ID конкретной ошибки (генерируется только если его уже нет во вложенной ошибке)
		args  []any
		stack StackTracer
		err   error
	}
)

// WithAttr - возвращает новую ошибку с прикреплённым к нему именованным атрибутом.
func (e *InstantError) WithAttr(name string, value any) *InstantError {
	c := *e
	c.args = append(c.args, name, value)

	return &c
}

// WithAttrs - возвращает новую ошибку с прикреплёнными к нему именованными атрибутами,
// которые должны передаваться попарно: название/значение.
func (e *InstantError) WithAttrs(attrs ...any) *InstantError {
	if len(attrs) == 0 {
		return e
	}

	c := *e
	c.args = append(c.args, attrs...)

	return &c
}

// ID - возвращает уникальный идентификатор случившейся ошибки,
// если в этой ошибке в опциях был установлен генератор ID ошибок.
// В противном случае вернётся пустая строка.
func (e *InstantError) ID() string {
	if e.id != nil {
		return *e.id
	}

	return ""
}

// Args - возвращает попарно аргументы используемые в сообщении ошибки: название/значение.
//
// Args возвращает оригинал слайса, не копию!
func (e *InstantError) Args() []any {
	return slices.Clip(e.args[:e.messageReplacer.CountArgs()])
}

// Attrs - возвращает попарно атрибуты прикреплённые к ошибке: ключ/значение.
//
// Attrs возвращает оригинал слайса, не копию!
func (e *InstantError) Attrs() []any {
	return slices.Clip(e.args[e.messageReplacer.CountArgs():])
}

// StackTrace - возвращает стек вызовов ошибки.
func (e *InstantError) StackTrace() StackTracer {
	return e.stack
}

// Error - возвращает ошибку в виде строки.
func (e *InstantError) Error() string {
	var buf strings.Builder

	countKeys := e.messageReplacer.CountArgs()

	if err := e.messageReplacer.ReplaceTo(&buf, e.args[:countKeys]); err != nil {
		buf.WriteString(e.message)
	}

	buf.WriteString(" [")
	buf.WriteString(e.Kind().String())

	if e.code != "" {
		buf.WriteString(", ")
		buf.WriteString(e.code)
	}

	if e.id != nil {
		buf.WriteString(", id=")
		buf.WriteString(*e.id)
	}

	if len(e.args) > countKeys {
		var (
			key   string
			value any
		)

		attrs := e.args[countKeys:]

		for len(attrs) > 0 {
			key, value, attrs = mrargs.PopKeyValue(attrs)

			buf.WriteString(", ")
			buf.WriteString(key)
			buf.WriteByte('=')
			buf.WriteString(mrargs.ToJSONValue(value))
		}
	}

	buf.WriteByte(']')

	if e.err != nil {
		buf.WriteString(": ")
		buf.WriteString(e.err.Error())
	}

	if e.stack != nil {
		stackCnt := e.stack.Count()

		for i := 0; i < stackCnt; i++ {
			function, file, line := e.stack.Source(i)

			if i == 0 {
				buf.WriteString(" in ")
			} else {
				buf.WriteString(", ")
			}

			if function != "" {
				buf.WriteByte('[')
				buf.WriteString(function)
				buf.WriteString("] ")
			}

			buf.WriteString(file)
			buf.WriteByte(':')
			buf.WriteString(strconv.Itoa(line))
		}
	}

	return buf.String()
}

// Is - сообщает, имеет ли указанная ошибка тот же
// прототип ошибки (для возможности использования errors.Is).
func (e *InstantError) Is(target error) bool {
	switch t := target.(type) {
	case *ProtoError:
		if e.pureError == t.pureError {
			return true
		}
	case *InstantError:
		if e.pureError == t.pureError {
			return true
		}
	}

	return errors.Is(e.Unwrap(), target)
}

// As - сообщает, имеет ли указанная ошибка тот же
// прототип ошибки (для возможности использования errors.As).
func (e *InstantError) As(target any) bool {
	if target == nil {
		panic("mrerr: target cannot be nil")
	}

	//nolint:dupl
	switch x := target.(type) {
	case **InstantError:
		if x == nil {
			panic("mrerr: target must be a non-nil pointer")
		}

		*x = e

		return true
	case *any:
		if _, ok := (*x).(*InstantError); ok {
			*x = e

			return true
		}
	case *InstantError:
		panic("mrerr: target must be a non-nil pointer")
	}

	return errors.As(e.Unwrap(), target)
}

// Unwrap - возвращает вложенную ошибку (errors.Is использует этот интерфейс).
func (e *InstantError) Unwrap() error {
	if we, ok := e.err.(*wrappedError); ok { //nolint:errorlint
		return (*InstantError)(we)
	}

	return e.err
}
