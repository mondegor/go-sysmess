package mrtrace

import (
	"context"

	tracectx "github.com/mondegor/go-sysmess/mrtrace/context"
)

type (
	// ContextManager - отвечает за установку ID процессов в контекст и за доступ к ним используемых в трейсинге.
	ContextManager struct {
		idGenerator idGenerator
	}

	idGenerator interface {
		GenerateID() string
	}
)

// NewContextManager - создаёт объект ContextManager.
func NewContextManager(idGenerator idGenerator) *ContextManager {
	return &ContextManager{
		idGenerator: idGenerator,
	}
}

// WithCorrelationID - возвращает результат вызова функции tracectx.WithCorrelationID.
// Метод WithGeneratedRequestID не предусмотрен, т.к. Correlation предполагается получать из вне, а не генерировать.
func (e *ContextManager) WithCorrelationID(ctx context.Context, id string) context.Context {
	return tracectx.WithCorrelationID(ctx, id)
}

// CorrelationID - возвращает результат вызова функции tracectx.CorrelationID.
func (e *ContextManager) CorrelationID(ctx context.Context) string {
	return tracectx.CorrelationID(ctx)
}

// WithGeneratedRequestID - генерирует ID запроса и возвращает результат вызова функции tracectx.WithRequestID.
func (e *ContextManager) WithGeneratedRequestID(ctx context.Context) context.Context {
	return tracectx.WithRequestID(ctx, e.idGenerator.GenerateID())
}

// RequestID - возвращает результат вызова функции tracectx.RequestID.
func (e *ContextManager) RequestID(ctx context.Context) string {
	return tracectx.RequestID(ctx)
}

// WithGeneratedProcessID - генерирует ID процесса и возвращает результат вызова функции tracectx.WithRequestID.
func (e *ContextManager) WithGeneratedProcessID(ctx context.Context) context.Context {
	return tracectx.WithProcessID(ctx, e.idGenerator.GenerateID())
}

// ProcessID - возвращает результат вызова функции tracectx.ProcessID.
func (e *ContextManager) ProcessID(ctx context.Context) string {
	return tracectx.ProcessID(ctx)
}

// WithGeneratedWorkerID - генерирует ID воркера и возвращает результат вызова функции tracectx.WithRequestID.
func (e *ContextManager) WithGeneratedWorkerID(ctx context.Context) context.Context {
	return tracectx.WithWorkerID(ctx, e.idGenerator.GenerateID())
}

// WorkerID - возвращает результат вызова функции tracectx.WorkerID.
func (e *ContextManager) WorkerID(ctx context.Context) string {
	return tracectx.WorkerID(ctx)
}

// WithGeneratedTaskID - генерирует ID корреляции запроса и возвращает результат вызова функции tracectx.WithCorrelationID.
func (e *ContextManager) WithGeneratedTaskID(ctx context.Context) context.Context {
	return tracectx.WithTaskID(ctx, e.idGenerator.GenerateID())
}

// TaskID - возвращает результат вызова функции tracectx.TaskID.
func (e *ContextManager) TaskID(ctx context.Context) string {
	return tracectx.TaskID(ctx)
}

// NewContextWithIDs - возвращает результат вызова функции NewContextWithIDs.
func (e *ContextManager) NewContextWithIDs(originalCtx context.Context) context.Context {
	return NewContextWithIDs(originalCtx)
}

// ExtractCorrelationID - возвращает результат вызова функции ExtractCorrelationID.
func (e *ContextManager) ExtractCorrelationID(ctx context.Context) string {
	return ExtractCorrelationID(ctx)
}

// ExtractKeysValues - возвращает результат вызова функции ExtractKeysValues.
func (e *ContextManager) ExtractKeysValues(ctx context.Context) []any {
	return ExtractKeysValues(ctx)
}
