package mrevent

import (
	"context"
)

type (
	// Emitter - заглушка реализующая интерфейс отправителя событий.
	nopEmitter struct{}
)

// NopEmitter - создаёт объект Emitter, который ничего не делает.
func NopEmitter() Emitter {
	return nopEmitter{}
}

// Emit - имитирует отправку указанного события.
func (e nopEmitter) Emit(_ context.Context, _ string, _ ...any) {}
