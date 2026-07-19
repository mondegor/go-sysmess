package slog

import (
	"time"
)

// formatTime - форматирует время лога согласно настройкам логгера.
// Используется всеми обработчиками, чтобы время выводилось единообразно.
func formatTime(tm time.Time, loc *time.Location, format string) string {
	return tm.In(loc).Format(format)
}
