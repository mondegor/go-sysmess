package mrevent

import (
	"context"
	"strings"
)

const (
	// sourceEventSeparator - разделитель между источником и названием события.
	sourceEventSeparator = ":"

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

// Receive - реализация интерфейса Receiver в виде функции для получения событий.
func (f ReceiveFunc) Receive(ctx context.Context, eventName string, args ...any) {
	f(ctx, eventName, args...)
}

type (
	emitter struct {
		receivers []Receiver
	}
)

// NewEmitter - создаёт объект Emitter.
func NewEmitter(receivers ...Receiver) Emitter {
	return &emitter{
		receivers: receivers,
	}
}

// NewEmitterWithSource - создаёт объект Emitter для отправки событий через указанных получателей добавляя первоисточник.
func NewEmitterWithSource(source string, receivers ...Receiver) Emitter {
	if source == "" {
		source = defaultSource
	}

	return &emitter{
		receivers: []Receiver{
			ReceiveFunc(
				func(ctx context.Context, sourceEventName string, args ...any) {
					// в базовом варианте eventName будет заменено на {source}:{eventName},
					// но если название события уже содержит источник события,
					// то новый источник становится после уже имеющегося источника c разделителем "/":
					// {parentSource}/{source}:{eventName}
					parentSource, eventName := ExtractEventName(sourceEventName)
					if parentSource != source {
						if parentSource != "" {
							sourceEventName = parentSource + SourceSeparator + source + sourceEventSeparator + eventName
						} else {
							sourceEventName = source + sourceEventSeparator + eventName
						}
					}

					for _, receive := range receivers {
						receive.Receive(ctx, sourceEventName, args...)
					}
				},
			),
		},
	}
}

// Emit - отправляет указанное событие всем зарегистрированным получателям.
// Гарантия успешного получение событий возлагается на Receiver.
func (e *emitter) Emit(ctx context.Context, eventName string, args ...any) {
	for _, r := range e.receivers {
		r.Receive(ctx, eventName, args...)
	}
}

// EmitterWithSource - создаёт объект Emitter
// для отправки событий через указанный базовый Emitter добавляя первоисточник.
func EmitterWithSource(base Emitter, source string) Emitter {
	return NewEmitterWithSource(source, ReceiveFunc(base.Emit))
}

// ExtractEventName - возвращает отдельно источник и название события из указанной строки.
func ExtractEventName(value string) (source, eventName string) {
	if s, n, ok := strings.Cut(value, sourceEventSeparator); ok {
		return s, n
	}

	return "", value
}
