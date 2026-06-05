package wire

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrlog"
	"github.com/mondegor/go-sysmess/mrtrace/context"
	"github.com/mondegor/go-sysmess/mrtrace/process"
	"github.com/mondegor/go-sysmess/util/crypt"
)

// InitTraceContextManager - создаёт process.ContextManager для управления идентификаторами
// процессов в контексте (correlation_id, request_id, task_id, worker_id, process_id).
// Параметры:
//   - processIDs - список пар (ключ, функция чтения, функция записи) для каждого идентификатора;
//   - logger - используется для предупреждений при отсутствии ключей.
func InitTraceContextManager(
	processIDs []process.KeyGetIDWithID,
	logger mrlog.Logger,
) (process.ContextManager, error) {
	cm, err := process.NewContextManager(
		processIDs,
		[]string{
			context.KeyCorrelationID,
			context.KeyRequestID,
			context.KeyTaskID,
			context.KeyWorkerID,
			context.KeyProcessID,
		},
		crypt.IDGeneratorFunc(process.GenerateID),
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("InitTraceContextManager: %w", err)
	}

	return cm, nil
}
