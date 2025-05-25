package nop

import "github.com/mondegor/go-sysmess/mrlog/slog"

// NewLoggerAdapter - создаёт объект slog.LoggerAdapter.
func NewLoggerAdapter() *slog.LoggerAdapter {
	return slog.New(handler{})
}
