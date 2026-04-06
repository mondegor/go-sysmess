package mrlog

import (
	stdlog "log"

	"github.com/mondegor/go-sysmess/mrlog/logger"
)

// Fatal is equivalent to [Print] followed by a call to [os.Exit](1).
func Fatal(v ...any) {
	if l, msg, ok := extractLogger(v); ok {
		FatalError(l, msg, v[2:]...) // stop if success
	}

	stdlog.Fatal(v...)
}

// Fatalf is equivalent to [Printf] followed by a call to [os.Exit](1).
func Fatalf(format string, v ...any) {
	stdlog.Fatalf(format, v...)
}

// Panic is equivalent to [Print] followed by a call to panic().
func Panic(v ...any) {
	stdlog.Panic(v...)
}

// Panicf is equivalent to [Printf] followed by a call to panic().
func Panicf(format string, v ...any) {
	stdlog.Panicf(format, v...)
}

// extractLogger - извлекает logger.Logger и сообщение из слайса аргументов.
// Ожидает, что первый аргумент - logger.Logger, а второй - строка.
func extractLogger(v []any) (l logger.Logger, msg string, ok bool) {
	if len(v) < 2 {
		return nil, "", false
	}

	if l, ok = v[0].(logger.Logger); !ok {
		return nil, "", false
	}

	if msg, ok = v[1].(string); !ok {
		return nil, "", false
	}

	return l, msg, true
}
