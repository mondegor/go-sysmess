package trace

import (
	tracectx "github.com/mondegor/go-sysmess/mrtrace/context"
	"github.com/mondegor/go-sysmess/mrtrace/process"
)

const (
	// KeyTaskID - cм. tracectx.KeyTaskID.
	KeyTaskID = tracectx.KeyTaskID
)

type (
	// ContextManager - cм. process.ContextManager.
	ContextManager = process.ContextManager
)
