package wire

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrlog"
	"github.com/mondegor/go-sysmess/mrtrace"
	"github.com/mondegor/go-sysmess/mrtrace/process"
	"github.com/mondegor/go-sysmess/util/crypt"
)

// InitTraceContextManager - создаёт и инициализирует process.ContextManager.
func InitTraceContextManager(
	processIDs []process.KeyGetIDWithID,
	logger mrlog.Logger,
) (*process.ContextManager, error) {
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
