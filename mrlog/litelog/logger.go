package litelog

import (
	"context"

	"github.com/mondegor/go-sysmess/mrlog"
)

type (
	// Logger - логгер без использования контекста, который используется,
	// например, при инициализации приложения.
	Logger struct {
		wrapped mrlog.Logger
	}
)

// NewLogger - создаёт объект Logger.
func NewLogger(l mrlog.Logger) *Logger {
	return &Logger{
		wrapped: l,
	}
}

// Debug - логирует сообщения на уровне mrlog.LevelDebug.
func (l *Logger) Debug(msg string, args ...any) {
	l.wrapped.Debug(context.Background(), msg, args...)
}

// DebugFunc - логирует сообщения на уровне mrlog.LevelDebug с их отложенным созданием.
// Применяется для того, чтобы исключить формирование в продуктовой среде больших отладочных
// сообщений с использованием многочисленных параметров.
func (l *Logger) DebugFunc(createMsg func() string, args ...any) {
	l.wrapped.DebugFunc(context.Background(), createMsg, args...)
}

// Info - логирует сообщения на уровне mrlog.LevelInfo.
func (l *Logger) Info(msg string, args ...any) {
	l.wrapped.Info(context.Background(), msg, args...)
}

// Warn - логирует сообщения на уровне mrlog.LevelWarn.
func (l *Logger) Warn(msg string, args ...any) {
	l.wrapped.Warn(context.Background(), msg, args...)
}

// Error - логирует сообщения на уровне mrlog.LevelError.
func (l *Logger) Error(msg string, args ...any) {
	l.wrapped.Error(context.Background(), msg, args...)
}

// ContextLogger - возвращает логгер с расширенным интерфейсом с использованием контекста.
func (l *Logger) ContextLogger() mrlog.Logger {
	return l.wrapped
}
