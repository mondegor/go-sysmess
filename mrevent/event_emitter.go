package mrevent

import (
	"context"
	"strings"
)

const (
	// SourceSeparator - разделитель между источниками событий при вложенности.
	SourceSeparator = "/"

	// sourceEventSeparator - разделитель между источником и названием события.
	sourceEventSeparator = ":"

	// defaultSource - источник событий по умолчанию, если не указан явно.
	defaultSource = "general.events"
)

type (
	// Emitter - интерфейс отправителя событий.
	// Позволяет отправлять события всем зарегистрированным получателям.
	Emitter interface {
		Emit(ctx context.Context, eventName string, args ...any)
	}

	// Receiver - интерфейс получателя событий.
	// Получает события от Emitter и обрабатывает их.
	Receiver interface {
		Receive(ctx context.Context, eventName string, args ...any)
	}

	// ReceiveFunc - тип функции для реализации интерфейса Receiver.
	// Позволяет использовать обычную функцию как получатель событий.
	ReceiveFunc func(ctx context.Context, eventName string, args ...any)
)

// Receive - реализует интерфейс Receiver для типа ReceiveFunc.
// Позволяет использовать обычную функцию как получатель событий.
func (f ReceiveFunc) Receive(ctx context.Context, eventName string, args ...any) {
	f(ctx, eventName, args...)
}

type (
	emitter struct {
		receivers []Receiver
	}
)

// NewEmitter - создаёт новый Emitter, который отправляет события
// всем указанным получателям (Receiver) в порядке их передачи.
func NewEmitter(receivers ...Receiver) Emitter {
	return &emitter{
		receivers: receivers,
	}
}

// NewEmitterWithSource - создаёт Emitter, который добавляет источник к имени события
// перед отправкой указанным получателям.
// Если событие уже содержит источник, новый источник добавляется после существующего
// через SourceSeparator (например: "parent/source:event").
// Если параметр source пустой, используется defaultSource.
// Это позволяет отслеживать иерархию источников событий.
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
// Метод не гарантирует успешную доставку - ответственность за обработку ошибок
// лежит на каждом отдельном Receiver.
func (e *emitter) Emit(ctx context.Context, eventName string, args ...any) {
	for _, r := range e.receivers {
		r.Receive(ctx, eventName, args...)
	}
}

// EmitterWithSource - создаёт обёртку над базовым Emitter, которая добавляет
// источник к имени события перед отправкой.
// Аналогичен NewEmitterWithSource, но принимает существующий Emitter вместо Receiver.
func EmitterWithSource(base Emitter, source string) Emitter {
	return NewEmitterWithSource(source, ReceiveFunc(base.Emit))
}

// ExtractEventName - разделяет строку события на источник и имя события.
// Ожидается формат "source:eventName". Если разделитель отсутствует,
// возвращается пустой источник и полное значение как eventName.
func ExtractEventName(value string) (source, eventName string) {
	if s, n, ok := strings.Cut(value, sourceEventSeparator); ok {
		return s, n
	}

	return "", value
}
