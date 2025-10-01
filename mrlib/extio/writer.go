package extio

import (
	"context"
	"io"
)

type (
	logger interface {
		Error(ctx context.Context, msg string, args ...any)
	}
)

// Write - адаптер для вызова io.Writer указанного объекта с логированием ошибки.
func Write(ctx context.Context, logger logger, w io.Writer, bytes []byte) {
	if len(bytes) == 0 {
		return
	}

	if _, err := w.Write(bytes); err != nil {
		logger.Error(ctx, "write failed", "error", err)
	}
}
