package mrtrace

import (
	"context"
	"fmt"

	"github.com/mondegor/go-sysmess/mrtrace"
	"github.com/mondegor/go-sysmess/mrtrace/process"
	"github.com/mondegor/go-sysmess/util/crypt"
)

type (
	logger interface {
		Warn(ctx context.Context, msg string, args ...any)
	}
)

// InitTraceContextManager - создаёт process.ContextManager для управления идентификаторами
// процессов в контексте (correlation_id, request_id, task_id, worker_id, process_id).
// Параметры:
//   - processIDs - список пар (ключ, функция чтения, функция записи) для каждого идентификатора;
//   - logger - используется для предупреждений при отсутствии ключей.
func InitTraceContextManager(
	processIDs []process.KeyGetIDWithID,
	logger logger,
) (mrtrace.ContextManager, error) {
	cm, err := process.NewContextManager(
		processIDs,
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
