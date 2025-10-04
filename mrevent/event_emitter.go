package mrevent

import (
	"context"
	"strings"
)

//go:generate mockgen -source=event_emitter.go -destination=./mock/event_emitter.go

const (
	// SourceEventSeparator - разделитель между источником и названием события.
	SourceEventSeparator = ":"

	// SourceSeparator - разделитель между источниками.
	SourceSeparator = "/"

	defaultSource = "general.events"
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

// Emit - отправляет указанное событие всем зарегистрированным получателям.
// Гарантия успешного получение событий возлагается на Receiver.
func (e *emitter) Emit(ctx context.Context, eventName string, args ...any) {
	for _, r := range e.receivers {
		r.Receive(ctx, eventName, args...)
	}
}

type (
	// sourceEmitter - расширяет возможности Emitter добавляя к нему источник данных.
	sourceEmitter struct {
		base   Emitter
		source string
	}
)

// NewSourceEmitter - создаёт объект sourceEmitter.
func NewSourceEmitter(base Emitter, source string) Emitter {
	if source == "" {
		source = defaultSource
	}

	return &sourceEmitter{
		base:   base,
		source: source,
	}
}

// Emit - отправляет указанное событие добавляя первоисточник.
func (e *sourceEmitter) Emit(ctx context.Context, eventName string, args ...any) {
	separator := SourceEventSeparator

	// в базовом варианте eventName будет заменено на {source}:{eventName},
	// но если название события уже содержит источник,
	// то после первоисточника будет стоять "/": {source}/{eventSource}:{eventName}
	if strings.Contains(eventName, separator) {
		separator = SourceSeparator
	}

	e.base.Emit(ctx, e.source+separator+eventName, args...)
}
