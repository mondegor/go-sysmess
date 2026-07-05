package mrlog

import (
	"github.com/mondegor/go-core/mrtrace"
	"github.com/mondegor/go-core/mrtrace/context"
	"github.com/mondegor/go-core/mrtrace/process"
)

type (
	// LoggerConfig - конфигурация для создания mrlog.Logger.
	// Содержит параметры окружения, форматирования и контекстных идентификаторов.
	LoggerConfig struct {
		// Environment - название окружения (например: "production", "testing", "local").
		Environment string

		// Version - версия приложения.
		Version string

		// Level - минимальный уровень логирования (debug, info, warn, error).
		Level string

		// JsonFormat - включает вывод в формате JSON.
		JsonFormat bool

		// TimeFormat - формат вывода времени (RFC3339, DateTime и др.).
		TimeFormat string

		// ColorMode - включает цветной вывод в консоль (игнорируется при JsonFormat=true).
		ColorMode bool

		// ContextProcessIDs - список идентификаторов процессов для извлечения/установки в контекст.
		ContextProcessIDs []process.KeyGetIDWithID
	}
)

// defaultProcessIDs - предустановленные ключи всех поддерживаемых идентификаторов процессов.
// Включает: correlation_id, request_id, process_id, worker_id, task_id.
var defaultProcessIDs = []process.KeyGetIDWithID{ //nolint:gochecknoglobals
	{
		Key:    mrtrace.KeyCorrelationID,
		GetID:  context.CorrelationID,
		WithID: context.WithCorrelationID,
	},
	{
		Key:    mrtrace.KeyRequestID,
		GetID:  context.RequestID,
		WithID: context.WithRequestID,
	},
	{
		Key:    mrtrace.KeyProcessID,
		GetID:  context.ProcessID,
		WithID: context.WithProcessID,
	},
	{
		Key:    mrtrace.KeyWorkerID,
		GetID:  context.WorkerID,
		WithID: context.WithWorkerID,
	},
	{
		Key:    mrtrace.KeyTaskID,
		GetID:  context.TaskID,
		WithID: context.WithTaskID,
	},
}

// DefaultProcessIDs - возвращает копию списка ключей процессов по умолчанию.
// Эти ключи используются для записи и чтения идентификаторов из контекста:
// correlation_id, request_id, process_id, worker_id, task_id.
func DefaultProcessIDs() []process.KeyGetIDWithID {
	return defaultProcessIDs
}
