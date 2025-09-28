package mrtrace

import (
	"context"

	tracectx "github.com/mondegor/go-sysmess/mrtrace/context"
)

type (
	// ContextManager - отвечает за установку ID процессов в контекст и за доступ к ним используемых в трейсинге.
	ContextManager struct {
		idGenerator IdentifierGenerator
	}
)

// NewContextManager - создаёт объект ContextManager.
func NewContextManager(idGenerator IdentifierGenerator) *ContextManager {
	return &ContextManager{
		idGenerator: idGenerator,
	}
}

// WithTraceID - генерирует ID корреляции запроса и возвращает результат вызова функции WithTraceID.
func (e *ContextManager) WithTraceID(ctx context.Context) context.Context {
	return tracectx.WithTraceID(ctx, e.idGenerator.GenerateID())
}

// TraceID - возвращает результат вызова функции TraceID.
func (e *ContextManager) TraceID(ctx context.Context) string {
	return tracectx.TraceID(ctx)
}

// WithCorrelationID - генерирует ID корреляции запроса и возвращает результат вызова функции WithCorrelationID.
func (e *ContextManager) WithCorrelationID(ctx context.Context) context.Context {
	return tracectx.WithCorrelationID(ctx, e.idGenerator.GenerateID())
}

// CorrelationID - возвращает результат вызова функции CorrelationID.
func (e *ContextManager) CorrelationID(ctx context.Context) string {
	return tracectx.CorrelationID(ctx)
}

// WithRequestID - генерирует ID запроса и возвращает результат вызова функции WithRequestID.
func (e *ContextManager) WithRequestID(ctx context.Context) context.Context {
	return tracectx.WithRequestID(ctx, e.idGenerator.GenerateID())
}

// RequestID - возвращает результат вызова функции RequestID.
func (e *ContextManager) RequestID(ctx context.Context) string {
	return tracectx.RequestID(ctx)
}

// WithProcessID - генерирует ID процесса и возвращает результат вызова функции WithRequestID.
func (e *ContextManager) WithProcessID(ctx context.Context) context.Context {
	return tracectx.WithProcessID(ctx, e.idGenerator.GenerateID())
}

// ProcessID - возвращает результат вызова функции ProcessID.
func (e *ContextManager) ProcessID(ctx context.Context) string {
	return tracectx.ProcessID(ctx)
}

// WithWorkerID - генерирует ID воркера и возвращает результат вызова функции WithRequestID.
func (e *ContextManager) WithWorkerID(ctx context.Context) context.Context {
	return tracectx.WithWorkerID(ctx, e.idGenerator.GenerateID())
}

// WorkerID - возвращает результат вызова функции WorkerID.
func (e *ContextManager) WorkerID(ctx context.Context) string {
	return tracectx.WorkerID(ctx)
}

// WithTaskID - генерирует ID корреляции запроса и возвращает результат вызова функции WithCorrelationID.
func (e *ContextManager) WithTaskID(ctx context.Context) context.Context {
	return tracectx.WithTaskID(ctx, e.idGenerator.GenerateID())
}

// TaskID - возвращает результат вызова функции TaskID.
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
