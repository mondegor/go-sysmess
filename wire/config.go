package wire

import (
	"github.com/mondegor/go-sysmess/mrtrace/context"
	"github.com/mondegor/go-sysmess/mrtrace/process"
)

type (
	// ErrorConfig - конфигурация для инициализации runtime-ошибок.
	// Определяет, использовать ли стек вызовов и с какой глубиной.
	ErrorConfig struct {
		// HasCaller - включает сбор информации (файл/строка/функция).
		HasCaller bool

		// CallerDepth - глубина стека вызовов (количество кадров).
		CallerDepth uint8

		// CallerShowFunc - включает отображение имени функции.
		CallerShowFunc bool

		// CallerUpperBounds - список имён пакетов/функций для ограничения стека вызова.
		CallerUpperBounds []string
	}

	// LoggerConfig - конфигурация для создания slog.LoggerAdapter.
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
		Key:    context.KeyCorrelationID,
		GetID:  context.CorrelationID,
		WithID: context.WithCorrelationID,
	},
	{
		Key:    context.KeyRequestID,
		GetID:  context.RequestID,
		WithID: context.WithRequestID,
	},
	{
		Key:    context.KeyProcessID,
		GetID:  context.ProcessID,
		WithID: context.WithProcessID,
	},
	{
		Key:    context.KeyWorkerID,
		GetID:  context.WorkerID,
		WithID: context.WithWorkerID,
	},
	{
		Key:    context.KeyTaskID,
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
