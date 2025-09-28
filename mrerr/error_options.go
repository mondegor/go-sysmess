package mrerr

import (
	"context"

	"github.com/mondegor/go-sysmess/mrerr/stacktrace"
	"github.com/mondegor/go-sysmess/mrerrors"
	"github.com/mondegor/go-sysmess/mrlib/crypt"
	"github.com/mondegor/go-sysmess/mrmsg"
)

type (
	// Option - настройка объекта mrerrors.ProtoError.
	Option func(o *options)

	// OptionsHandler - обработчик формирования списка опций для указанных типов ошибок.
	OptionsHandler interface {
		Options(kind ErrorKind, code, message string) []Option
	}

	// OptionsHandlerFunc - обработчик формирования списка опций в виде функции.
	OptionsHandlerFunc func(kind ErrorKind, code, message string) []Option
)

// Options - реализация интерфейса OptionsHandler в виде функции для формирования списка опций.
func (f OptionsHandlerFunc) Options(kind ErrorKind, code, message string) []Option {
	return f(kind, code, message)
}

// WithCode - устанавливает код ошибки.
func WithCode(value string) Option {
	return func(o *options) {
		mrerrors.WithProtoCode(value)(o.proto)
	}
}

// WithArgsReplacer - устанавливает функцию возвращающую объект для замены аргументов,
// указанных в сообщении на их конкретные значения.
func WithArgsReplacer(value func(message string) mrerrors.MessageReplacer) Option {
	return func(o *options) {
		if o.changedArgsReplacer {
			return
		}

		o.changedArgsReplacer = true
		mrerrors.WithProtoArgsReplacer(value)(o.proto)
	}
}

// WithCaller - устанавливает функцию, которая создаёт стек вызовов при создании экземпляра ошибки.
func WithCaller(value func() mrerrors.StackTracer) Option {
	return func(o *options) {
		if o.changedCaller {
			return
		}

		o.changedCaller = true
		mrerrors.WithProtoCaller(value)(o.proto)
	}
}

// WithDisabledCaller - отключает формирование стека вызовов при создании экземпляра ошибки.
func WithDisabledCaller() Option {
	return WithCaller(nil)
}

// WithOnCreated - устанавливает обработчик события создания экземпляра ошибки.
func WithOnCreated(value func(ctx context.Context, err error) (instanceID string)) Option {
	return func(o *options) {
		if o.changedOnCreated {
			return
		}

		o.changedOnCreated = true
		mrerrors.WithProtoOnCreated(value)(o.proto)
	}
}

// WithDisabledOnCreated - отключает обработчик события создания экземпляра ошибки.
func WithDisabledOnCreated() Option {
	return WithOnCreated(nil)
}

// WithDefaultArgsReplacer - устанавливает функцию возвращающую объект по умолчанию для замены аргументов,
// указанных в сообщении на их конкретные значения.
func WithDefaultArgsReplacer() Option {
	return WithArgsReplacer(
		func(message string) mrerrors.MessageReplacer {
			return mrmsg.NewMessageReplacer("{", "}", message)
		},
	)
}

// WithDefaultCaller - устанавливает формирование стека вызовов с опциями по умолчанию при создании экземпляра ошибки.
func WithDefaultCaller() Option {
	const (
		stackDepth = 16
		funcPrefix = "github.com/mondegor/go-sysmess/mrerr.(*ProtoError)."
	)

	caller := stacktrace.New(
		stacktrace.WithDepth(stackDepth),
		stacktrace.WithStackTraceFilter(
			stacktrace.TrimUpperFilter(
				[]string{
					funcPrefix + "New",
					funcPrefix + "NewContext",
					funcPrefix + "Wrap",
					funcPrefix + "WrapContext",
				},
			),
		),
	)

	return WithCaller(
		func() mrerrors.StackTracer {
			return caller.StackTrace()
		},
	)
}

// WithDefaultOnCreated - устанавливает обработчик по умолчанию события создания экземпляра ошибки,
// который генерирует уникальный ID ошибки и возвращает его.
func WithDefaultOnCreated() Option {
	return WithOnCreated(
		func(_ context.Context, _ error) (instanceID string) {
			return crypt.GenerateInstanceID()
		},
	)
}

// DefaultOptionsHandler - возвращает обработчик по умолчанию для формирования опций по умолчанию,
// необходимый для вызова InitDefaultOptions().
func DefaultOptionsHandler() OptionsHandlerFunc {
	internalOpts := []Option{
		WithDefaultArgsReplacer(),
		WithDefaultCaller(),
		WithDefaultOnCreated(),
	}

	systemOpts := []Option{
		WithDefaultArgsReplacer(),
		WithDefaultOnCreated(),
	}

	userOpts := []Option{
		WithDefaultArgsReplacer(),
	}

	return func(kind ErrorKind, _, _ string) []Option {
		if kind == ErrorKindUser {
			return userOpts
		}

		if kind == ErrorKindSystem {
			return systemOpts
		}

		return internalOpts
	}
}
