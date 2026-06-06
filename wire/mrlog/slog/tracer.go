package slog

import (
	"github.com/mondegor/go-sysmess/mrlog/level"
	"github.com/mondegor/go-sysmess/mrlog/slog"
	"github.com/mondegor/go-sysmess/mrtrace"
)

// InitTracer - создаёт mrtrace.Tracer на основе slog.LoggerAdapter.
// Если уровень логирования включает Debug, возвращает сам логгер как Tracer.
// В противном случае возвращает заглушку (NopTracer), чтобы избежать накладных расходов.
func InitTracer(logger *slog.LoggerAdapter) mrtrace.Tracer {
	if logger.Enabled(level.Debug) {
		return logger
	}

	return mrtrace.NopTracer()
}
