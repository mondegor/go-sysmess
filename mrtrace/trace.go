package mrtrace

import (
	"context"
)

//go:generate mockgen -source=trace.go -destination=./mock/trace.go

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
	// Tracer - трейсер для фиксации запросов к сервисам и их ответов.
	Tracer interface {
		Enabled() bool
		Trace(ctx context.Context, args ...any)
	}

	// ContextManager - отвечает за установку ID процессов в контекст и за доступ к ним используемых в трейсинге.
	ContextManager interface {
		ID(ctx context.Context) string
		WithID(ctx context.Context, id string) context.Context
		WithGeneratedID(ctx context.Context) context.Context

		NewContextWithIDs(originalCtx context.Context) context.Context
		ExtractCorrelationID(ctx context.Context) string
		ExtractKeysValues(ctx context.Context) []any
	}
)
