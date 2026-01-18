package slog

import (
	"github.com/mondegor/go-sysmess/mrlog/level"
	"github.com/mondegor/go-sysmess/mrlog/slog"
	"github.com/mondegor/go-sysmess/mrtrace"
)

// InitTracer - создаёт и инициализирует mrtrace.Tracer на основе slog.LoggerAdapter.
func InitTracer(logger *slog.LoggerAdapter) mrtrace.Tracer {
	if logger.Enabled(level.Debug) {
		return logger
	}

	return mrtrace.NopTracer()
}
