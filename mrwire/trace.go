package mrwire

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrlib/crypt"
	"github.com/mondegor/go-sysmess/mrlog"
	"github.com/mondegor/go-sysmess/mrtrace"
	"github.com/mondegor/go-sysmess/mrtrace/context"
	"github.com/mondegor/go-sysmess/mrtrace/logtracer"
	"github.com/mondegor/go-sysmess/mrtrace/noptracer"
	"github.com/mondegor/go-sysmess/mrtrace/process"
)

type (
	// TracerConfig - опции создаваемого трейсера.
	TracerConfig struct {
		Environment string
		Version     string
		Enabled     bool
	}
)

// InitTracer - создаёт и инициализирует mrtrace.Tracer.
func InitTracer(cfg TracerConfig, logger mrlog.Logger) mrtrace.Tracer {
	if cfg.Enabled {
		if logger != nil && logger.Enabled(mrlog.LevelDebug) {
			return logtracer.NewTracer(logger)
		}
	}

	return noptracer.NewTracer()
}

// InitTraceContextManager - создаёт и инициализирует process.ContextManager.
func InitTraceContextManager(logger mrlog.Logger) (*process.ContextManager, error) {
	cm, err := process.NewContextManager(
		traceProcessIDs(),
		[]string{
			mrtrace.KeyCorrelationID,
			mrtrace.KeyRequestID,
			mrtrace.KeyTaskID,
			mrtrace.KeyWorkerID,
			mrtrace.KeyProcessID,
		},
		crypt.IDGeneratorFunc(process.GenerateID),
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("InitTraceContextManager: %w", err)
	}

	return cm, nil
}

func traceProcessIDs() []process.KeyGetIDWithID {
	return []process.KeyGetIDWithID{
		{
			mrtrace.KeyCorrelationID,
			context.CorrelationID,
			context.WithCorrelationID,
		},
		{
			mrtrace.KeyRequestID,
			context.RequestID,
			context.WithRequestID,
		},
		{
			mrtrace.KeyProcessID,
			context.ProcessID,
			context.WithProcessID,
		},
		{
			mrtrace.KeyWorkerID,
			context.WorkerID,
			context.WithWorkerID,
		},
		{
			mrtrace.KeyTaskID,
			context.TaskID,
			context.WithTaskID,
		},
	}
}
