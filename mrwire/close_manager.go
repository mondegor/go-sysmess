package mrwire

import (
	"fmt"
	"io"

	"github.com/mondegor/go-sysmess/mrerr/mr"
	"github.com/mondegor/go-sysmess/mrlog"
)

type (
	// CloseManager - менеджер для одновременного
	// закрытия всех зарегистрированных ресурсов.
	CloseManager struct {
		logger       mrlog.Logger
		resourceList []io.Closer
	}
)

// NewCloseManager - создаёт объект CloseManager.
func NewCloseManager(logger mrlog.Logger) *CloseManager {
	return &CloseManager{
		logger: logger,
	}
}

// Register - регистрирует указанный ресурс в своём списке.
func (rl *CloseManager) Register(resource io.Closer) {
	rl.resourceList = append(rl.resourceList, resource)

	mrlog.Debug(rl.logger, fmt.Sprintf("Resource %T was registered", resource))
}

// Close - последовательно вызывается метод Close() для всех зарегистрированных ресурсов.
func (rl *CloseManager) Close() {
	for _, resource := range rl.resourceList {
		rl.close(resource)
	}
}

func (rl *CloseManager) close(resource io.Closer) {
	if err := resource.Close(); err != nil {
		mrlog.Error(
			rl.logger,
			"CloseManager.Close()",
			"error", mr.ErrInternalFailedToClose.Wrap(err),
			"resource", fmt.Sprintf("%#v", resource),
		)

		return
	}

	mrlog.Debug(rl.logger, fmt.Sprintf("Resource %T was closed", resource))
}
