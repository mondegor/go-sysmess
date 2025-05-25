package mrlog

import (
	"log/slog"
)

type (
	// Level - уровень логирования.
	Level int8
)

const (
	LevelDebug = Level(slog.LevelDebug)     // LevelDebug - LevelWarn + LevelWarn + LevelInfo - отладочные сообщения.
	LevelInfo  = Level(slog.LevelInfo)      // LevelInfo - LevelError + LevelWarn + информационные сообщения.
	LevelWarn  = Level(slog.LevelWarn)      // LevelWarn - LevelError + предупреждения.
	LevelError = Level(slog.LevelError)     // LevelError - ошибки (максимальный уровень, который можно назначить в конфиге).
	LevelFatal = Level(slog.LevelError + 4) // LevelFatal - дополнительный уровень для фатальных ошибок (служебный).
	LevelTrace = Level(slog.LevelError + 8) // LevelTrace - режим трассировки (служебный, управляется отдельно).
)

// String - возвращает значение уровня в виде строки.
func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	case LevelTrace:
		return "TRACE"
	}

	return "UNKNOWN"
}
