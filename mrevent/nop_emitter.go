package mrevent

import (
	"context"
)

type (
	// nopEmitter - заглушка, реализующая интерфейс Emitter.
	// Игнорирует все отправляемые события.
	nopEmitter struct{}
)

// NopEmitter - создаёт Emitter, который игнорирует все события.
// Полезно для тестов или когда отправка событий не требуется.
func NopEmitter() Emitter {
	return nopEmitter{}
}

// Emit - игнорирует событие (заглушка для интерфейса Emitter).
func (e nopEmitter) Emit(_ context.Context, _ string, _ ...any) {}
