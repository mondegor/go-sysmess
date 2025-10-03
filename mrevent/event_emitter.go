package mrevent

import (
	"context"
	"strings"
)

const (
	// SourceEventSeparator - разделитель между источником и названием события.
	SourceEventSeparator = ":"

	// DefaultSource - название источника по умолчанию.
	DefaultSource = "general"
)

type (
	// Emitter - отправитель событий.
	Emitter interface {
		Emit(ctx context.Context, eventName string, args ...any)
	}

	// Receiver - получатель событий.
	Receiver interface {
		Receive(ctx context.Context, eventName string, args ...any)
	}

	// ReceiveFunc - получатель событий в виде функции.
	ReceiveFunc func(ctx context.Context, eventName string, args ...any)
)

// Receive - реализация интерфейса Receiver в виде функции для получения события.
func (f ReceiveFunc) Receive(ctx context.Context, eventName string, args ...any) {
	f(ctx, eventName, args...)
}

// ExtractEventName - получение из указанного значение отдельно источник и название события.
func ExtractEventName(value string) (source, eventName string) {
	if s, n, ok := strings.Cut(value, SourceEventSeparator); ok {
		return s, n
	}

	return "", value
}

type (
	emitter struct {
		receivers []Receiver
	}
)

// NewEmitter - создаёт объект emitter.
func NewEmitter(receivers ...Receiver) Emitter {
	return &emitter{
		receivers: receivers,
	}
}

// Emit - отправляет указанное событие (не гарантирует их успешное получение).
func (e *emitter) Emit(ctx context.Context, eventName string, args ...any) {
	for _, r := range e.receivers {
		go func() {
			r.Receive(ctx, eventName, args...)
		}()
	}
}

type (
	// emitterWrapper - расширяет возможности Emitter добавляя к нему источник данных.
	emitterWrapper struct {
		eventEmitter Emitter
		source       string
	}
)

// NewEmitterWrapper - создаёт объект emitterWrapper.
func NewEmitterWrapper(eventEmitter Emitter, source string) Emitter {
	if source == "" {
		source = DefaultSource
	}

	return &emitterWrapper{
		eventEmitter: eventEmitter,
		source:       source,
	}
}

// Emit - отправляет указанное событие добавляя первоисточник.
func (e *emitterWrapper) Emit(ctx context.Context, eventName string, args ...any) {
	separator := SourceEventSeparator

	// в базовом варианте eventName будет заменено на {source}:{eventName},
	// но если название события уже содержит источник,
	// то после первоисточника будет стоять слеш: {source}/{eventSource}:{eventName}
	if strings.Contains(eventName, separator) {
		separator = "/"
	}

	e.eventEmitter.Emit(ctx, e.source+separator+eventName, args...)
}
