package mrtrace

import (
	"context"
)

type (
	// Tracer - интерфейс для трассировки запросов между сервисами.
	// Фиксирует входящие/исходящие запросы и их параметры для аудита и отладки.
	Tracer interface {
		Trace(ctx context.Context, args ...any)
	}
)
