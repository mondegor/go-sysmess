package mrlog

import (
	"github.com/mondegor/go-sysmess/mrlog/level"
	"github.com/mondegor/go-sysmess/mrlog/logger"
)

type (
	// Logger - интерфейс логирования сообщений и ошибок с использованием контекста.
	Logger = logger.Logger
)

// WithAttrs - возвращает логгер с прикреплёнными к нему атрибутами.
func WithAttrs(l logger.Logger, attrs ...any) logger.Logger {
	return logger.WithAttrs(l, attrs...)
}

// DebugEnabled - сообщает логирует ли указанный логгер сообщения уровня level.Debug.
// Если logger = nil, то будет возвращено false.
func DebugEnabled(l logger.Logger) bool {
	return logger.Enabled(l, level.Debug)
}

// InfoEnabled - сообщает логирует ли указанный логгер сообщения уровня level.Info.
// Если logger = nil, то будет возвращено false.
func InfoEnabled(l logger.Logger) bool {
	return logger.Enabled(l, level.Info)
}

// Debug - логирует сообщения на уровне level.Debug без использования контекста.
func Debug(l logger.Logger, msg string, args ...any) {
	logger.Debug(l, msg, args...)
}

// DebugFunc - логирует сообщения на уровне level.Debug без использования контекста с их отложенным созданием сообщения.
func DebugFunc(l logger.Logger, createMsg func() string, args ...any) {
	logger.DebugFunc(l, createMsg, args...)
}

// Info - логирует сообщения на уровне level.Info без использования контекста.
func Info(l logger.Logger, msg string, args ...any) {
	logger.Info(l, msg, args...)
}

// Warn - логирует сообщения на уровне level.Warn без использования контекста.
func Warn(l logger.Logger, msg string, args ...any) {
	logger.Warn(l, msg, args...)
}

// Error - логирует сообщения на уровне level.Error без использования контекста.
func Error(l logger.Logger, msg string, args ...any) {
	logger.Error(l, msg, args...)
}

// FatalError - логирует сообщения на уровне level.Fatal без использования контекста и останавливает приложение.
func FatalError(l logger.Logger, msg string, args ...any) {
	logger.Fatal(l, msg, args...)
}

// NopLogger - создаёт объект Logger, который ничего не делает.
func NopLogger() Logger {
	return logger.Nop()
}
