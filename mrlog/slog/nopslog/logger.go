package nopslog

import "github.com/mondegor/go-sysmess/mrlog/slog"

// New - создаёт объект slog.LoggerAdapter.
func New() *slog.LoggerAdapter {
	return slog.New(handler{})
}
