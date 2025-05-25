package mrerrors

import (
	"context"
)

type (
	// ProtoOption - настройка объекта ProtoError.
	ProtoOption func(e *ProtoError)
)

// WithProtoKind - устанавливает тип ошибки.
func WithProtoKind(value ErrorKind) ProtoOption {
	return func(e *ProtoError) {
		if e.kind >= ErrorKindInternal && e.kind <= ErrorKindUser {
			e.kind = value
		}
	}
}

// WithProtoCode - устанавливает код ошибки.
func WithProtoCode(value string) ProtoOption {
	return func(e *ProtoError) {
		e.code = value
	}
}

// WithProtoArgsReplacer - устанавливает функцию возвращающую объект для замены аргументов, указанных в сообщении на их конкретные значения.
func WithProtoArgsReplacer(value func(message string) MessageReplacer) ProtoOption {
	return func(e *ProtoError) {
		if value != nil {
			e.messageReplacer = value(e.message)
		} else {
			e.messageReplacer = newNoArgsRenderer(e.message)
		}
	}
}

// WithProtoCaller - устанавливает функцию, которая создаёт стек вызовов при создании экземпляра ошибки.
func WithProtoCaller(value func() StackTracer) ProtoOption {
	return func(e *ProtoError) {
		e.caller = value
	}
}

// WithProtoOnCreated - устанавливает обработчик события создания экземпляра ошибки.
func WithProtoOnCreated(value func(ctx context.Context, err error) (instanceID string)) ProtoOption {
	return func(e *ProtoError) {
		e.onCreated = value
	}
}
