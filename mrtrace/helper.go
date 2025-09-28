package mrtrace

import (
	"context"

	tracectx "github.com/mondegor/go-sysmess/mrtrace/context"
)

// NewContextWithIDs - возвращает новый контекст содержащий
// только все ID процессы, скопированные из указанного контекста.
func NewContextWithIDs(originalCtx context.Context) context.Context {
	ctx := context.Background()

	if originalCtx == nil || originalCtx == ctx {
		return ctx
	}

	if value := tracectx.TraceID(originalCtx); value != "" {
		ctx = tracectx.WithTraceID(ctx, value)
	}

	if value := tracectx.CorrelationID(originalCtx); value != "" {
		ctx = tracectx.WithCorrelationID(ctx, value)
	}

	if value := tracectx.ProcessID(originalCtx); value != "" {
		ctx = tracectx.WithProcessID(ctx, value)
	}

	if value := tracectx.RequestID(originalCtx); value != "" {
		ctx = tracectx.WithCorrelationID(ctx, value)
	}

	if value := tracectx.WorkerID(originalCtx); value != "" {
		ctx = tracectx.WithWorkerID(ctx, value)
	}

	if value := tracectx.TaskID(originalCtx); value != "" {
		ctx = tracectx.WithTaskID(ctx, value)
	}

	return ctx
}

// ExtractCorrelationID - возвращает первый попавшийся ID из указанного контекста,
// который можно использовать в качестве CorrelationID.
func ExtractCorrelationID(ctx context.Context) string {
	if value := tracectx.CorrelationID(ctx); value != "" {
		return value
	}

	if value := tracectx.RequestID(ctx); value != "" {
		return value
	}

	if value := tracectx.TaskID(ctx); value != "" {
		return value
	}

	if value := tracectx.WorkerID(ctx); value != "" {
		return value
	}

	if value := tracectx.ProcessID(ctx); value != "" {
		return value
	}

	return tracectx.TraceID(ctx)
}

// ExtractKeysValues - возвращает попарно (key/id-value) все имеющиеся
// ID процессов из указанного контекста.
func ExtractKeysValues(ctx context.Context) (keyValue []any) {
	if ctx == nil || ctx == context.Background() {
		return nil
	}

	keyValue = make([]any, 0, 6)

	if value := tracectx.TraceID(ctx); value != "" {
		keyValue = append(keyValue, KeyTraceID, value)
	}

	if value := tracectx.CorrelationID(ctx); value != "" {
		keyValue = append(keyValue, KeyCorrelationID, value)
	}

	if value := tracectx.ProcessID(ctx); value != "" {
		keyValue = append(keyValue, KeyProcessID, value)
	}

	if value := tracectx.RequestID(ctx); value != "" {
		keyValue = append(keyValue, KeyRequestID, value)
	}

	if value := tracectx.WorkerID(ctx); value != "" {
		keyValue = append(keyValue, KeyWorkerID, value)
	}

	if value := tracectx.TaskID(ctx); value != "" {
		keyValue = append(keyValue, KeyTaskID, value)
	}

	return keyValue[0:len(keyValue):len(keyValue)]
}
