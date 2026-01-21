package xio

import (
	"context"
	"fmt"
	"io"
)

type (
	// CloseManager - менеджер для одновременного
	// закрытия всех зарегистрированных ресурсов.
	CloseManager struct {
		logger    logger
		resources []io.Closer
	}
)

// NewCloseManager - создаёт объект CloseManager.
func NewCloseManager(logger logger) *CloseManager {
	return &CloseManager{
		logger: logger,
	}
}

// Register - регистрирует указанный ресурс в своём списке.
func (m *CloseManager) Register(resource io.Closer) {
	m.resources = append(m.resources, resource)

	m.logger.Debug(context.Background(), fmt.Sprintf("Resource %T was registered", resource))
}

// Close - последовательно вызывается метод Close() для всех зарегистрированных ресурсов.
func (m *CloseManager) Close() {
	for _, resource := range m.resources {
		m.close(resource)
	}
}

func (m *CloseManager) close(resource io.Closer) {
	if err := resource.Close(); err != nil {
		m.logger.Error(
			context.Background(),
			"CloseManager.Close(): failed to close object",
			"error", err,
			"resource", fmt.Sprintf("%#v", resource),
		)

		return
	}

	m.logger.Debug(context.Background(), fmt.Sprintf("Resource %T was closed", resource))
}
