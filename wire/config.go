package wire

import (
	"github.com/mondegor/go-sysmess/mrtrace"
	"github.com/mondegor/go-sysmess/mrtrace/context"
	"github.com/mondegor/go-sysmess/mrtrace/process"
)

type (
	// ErrorConfig - опции инициализации ошибок.
	ErrorConfig struct {
		HasCaller         bool
		CallerDepth       uint8
		CallerUpperBounds []string
	}

	// LoggerConfig - опции создаваемого логгера.
	LoggerConfig struct {
		Environment       string
		Version           string
		Level             string
		JsonFormat        bool
		TimeFormat        string
		ColorMode         bool
		ContextProcessIDs []process.KeyGetIDWithID
	}
)

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

// DefaultProcessIDs - возвращает настроенные ключи процессов,
// которые записываются в контекст и читаются из него.
func DefaultProcessIDs() []process.KeyGetIDWithID {
	return defaultProcessIDs
}
