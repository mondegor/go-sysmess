package mrtrace

import (
	"context"
)

const (
	// KeyCorrelationID - название ключа ID корреляции.
	KeyCorrelationID = "correlation_id"

	// KeyRequestID - название ключа ID запроса.
	KeyRequestID = "request_id"

	// KeyProcessID - название ключа ID процесса.
	KeyProcessID = "process_id"

	// KeyWorkerID - название ключа ID воркера.
	KeyWorkerID = "worker_id"

	// KeyTaskID - название ключа ID задачи.
	KeyTaskID = "task_id"
)

type (
	// Tracer - интерфейс для трассировки запросов между сервисами.
	// Фиксирует входящие/исходящие запросы и их параметры для аудита и отладки.
	Tracer interface {
		Trace(ctx context.Context, args ...any)
	}

	// ContextManager - управляет идентификаторами процессов в контексте для трассировки.
	// Позволяет получать, устанавливать и генерировать ID процессов (request_id, process_id и др.).
	ContextManager interface {
		ProcessID(ctx context.Context, key string) string
		WithProcessID(ctx context.Context, key, value string) context.Context
		WithGeneratedProcessID(ctx context.Context, key string) context.Context
		ExtractCorrelationID(ctx context.Context) string
		ExtractKeysValues(ctx context.Context) []any
	}
)
