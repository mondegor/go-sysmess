package mrtrace

import (
	"context"
)

const (
	// KeyTraceID - название ключа ID трейса.
	KeyTraceID = "trace_id"

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
	// Tracer - трейсер для фиксации запросов к сервисам и их ответов.
	Tracer interface {
		Enabled() bool
		Trace(ctx context.Context, args ...any)
	}

	// IdentifierGenerator - генератор уникальных идентификаторов
	// процессов используемых при трейсинге.
	IdentifierGenerator interface {
		GenerateID() string
	}

	// IdentifierGeneratorFunc - реализация интерфейса IdentifierGenerator.
	IdentifierGeneratorFunc func() string
)

// GenerateID - реализация интерфейса IdentifierGenerator.
func (f IdentifierGeneratorFunc) GenerateID() string {
	return f()
}
