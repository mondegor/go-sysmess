package mrerrors

import (
	"context"
	"strings"
)

//go:generate mockgen -source=proto_error.go -destination=./mock/proto_error.go

const (
	missingArg = "!MISSINGARG"
)

type (
	// ProtoError - прототип ошибки с поддержкой параметров, ID экземпляра ошибки и стека вызовов.
	ProtoError struct {
		*pureError
		caller    func() StackTracer
		onCreated func(ctx context.Context, err error) (instanceID string)
	}

	// StackTracer - предоставляет доступ к стеку вызовов ошибки.
	StackTracer interface {
		Count() int
		Source(index int) (function, file string, line int)
	}
)

// NewProto - создаёт объект ProtoError для создания конкретных ошибок данного типа.
func NewProto(message string, opts ...ProtoOption) *ProtoError {
	p := &ProtoError{
		pureError: &pureError{
			kind:            ErrorKindInternal,
			message:         message,
			messageReplacer: newNoArgsRenderer(message),
		},
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// New - создаётся новая ошибка на основе текущей ProtoError ошибки.
// Подробнее см. WrapContext().
func (e *ProtoError) New(args ...any) *InstantError {
	return e.NewContext(context.Background(), args...)
}

// NewContext - создаётся новая ошибка на основе текущей ProtoError ошибки,
// при этом вызываются обработчики caller и onCreated если они были установлены.
// Контекст используется только для передачи его обработчику onCreated.
func (e *ProtoError) NewContext(ctx context.Context, args ...any) *InstantError {
	c := &InstantError{
		pureError: e.pureError,
		args:      e.makeArgsAndAttrs(args),
	}

	// важно вызвать получение StackTrace как можно раньше
	if e.caller != nil {
		c.stack = e.caller()
	}

	if e.onCreated != nil {
		if instanceID := e.onCreated(ctx, c); instanceID != "" {
			c.id = &instanceID
		}
	}

	return c
}

// Wrap - создаёт новую ошибку на основе прототипа и оборачивает в неё указанную.
// Подробнее см. WrapContext().
func (e *ProtoError) Wrap(err error, args ...any) *InstantError {
	return e.WrapContext(context.Background(), err, args...)
}

// WrapContext - создаёт новую ошибку на основе текущей ProtoError и оборачивает в неё указанную.
// Контекст используется только для передачи его обработчику onCreated.
// Если указанная ошибка типа InstantError, то проверяется, был ли у этой ошибки
// сгенерированы ID и стек, и если да, то у новой ошибки эти параметры не генерируются,
// даже если соответствующие обработчики установлены для этой ошибки.
func (e *ProtoError) WrapContext(ctx context.Context, err error, args ...any) *InstantError {
	if err == nil {
		return e.NewContext(ctx, args...)
	}

	c := &InstantError{
		pureError: e.pureError,
		args:      e.makeArgsAndAttrs(args),
		err:       err,
	}

	if wrappedErr, ok := c.err.(*InstantError); ok { //nolint:errorlint
		// id is raising to the top
		if wrappedErr.id != nil {
			c.id = wrappedErr.id
		}

		if len(wrappedErr.args) > wrappedErr.messageReplacer.CountArgs() {
			c.args = append(c.args, wrappedErr.args[wrappedErr.messageReplacer.CountArgs():]...)
		}

		// если стек был установлен во вложенном объекте,
		// то запрещается генерация стека текущему объекту
		if wrappedErr.stack != nil {
			c.stack = wrappedErr.stack
		}

		c.err = (*wrappedError)(wrappedErr)
	}

	if e.caller != nil && c.stack == nil {
		c.stack = e.caller()
	}

	if e.onCreated != nil && c.id == nil {
		if instanceID := e.onCreated(ctx, c); instanceID != "" {
			c.id = &instanceID
		}
	}

	return c
}

// Error - возвращает ошибку в виде строки.
func (e *ProtoError) Error() string {
	var buf strings.Builder

	buf.WriteString(e.message)

	buf.WriteString(" [")
	buf.WriteString(e.Kind().String())

	if e.code != "" {
		buf.WriteString(", ")
		buf.WriteString(e.code)
	}

	buf.WriteByte(']')

	return buf.String()
}

// Is - сообщает, имеет ли указанная ошибка тот же
// прототип ошибки (для возможности использования errors.Is).
func (e *ProtoError) Is(target error) bool {
	if e == target {
		return true
	}

	if t, ok := target.(*InstantError); ok {
		return e.pureError == t.pureError
	}

	return false
}

// As - сообщает, имеет ли указанная ошибка тот же
// прототип ошибки (для возможности использования errors.As).
func (e *ProtoError) As(target any) bool {
	if target == nil {
		panic("mrerr: target cannot be nil")
	}

	//nolint:dupl
	switch x := target.(type) {
	case **ProtoError:
		if x == nil {
			panic("mrerr: target must be a non-nil pointer")
		}

		*x = e

		return true
	case *any:
		if _, ok := (*x).(*ProtoError); ok {
			*x = e

			return true
		}
	case *ProtoError:
		panic("mrerr: target must be a non-nil pointer")
	}

	return false
}

func (e *ProtoError) makeArgsAndAttrs(args []any) []any {
	countKeys := e.messageReplacer.CountArgs()

	if len(args) == 0 && countKeys == 0 {
		return nil
	}

	var attrs []any

	if len(args) > countKeys {
		attrs = args[countKeys:]
		args = args[:countKeys]
	}

	newArgs := make([]any, countKeys+len(attrs))

	if e.kind == ErrorKindUser {
		for i := 0; i < len(args); i++ {
			// для пользовательских ошибок все аргументы типа InstantError
			// заменяются на тип userArgError для уменьшения выводимой информации
			if ee, ok := args[i].(*InstantError); ok {
				newArgs[i] = (*userArgError)(ee)

				continue
			}

			newArgs[i] = args[i]
		}
	} else {
		copy(newArgs, args)
	}

	for i := len(args); i < countKeys; i++ {
		newArgs[i] = missingArg
	}

	if len(attrs) > 0 {
		copy(newArgs[countKeys:], attrs)
	}

	return newArgs
}
