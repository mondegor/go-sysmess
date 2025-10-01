package mrlog

import (
	"context"
	"os"
)

type (
	// Logger - интерфейс логирования сообщений и ошибок с использованием контекста.
	Logger interface {
		WithAttrs(args ...any) Logger
		Enabled(level Level) bool

		Debug(ctx context.Context, msg string, args ...any)
		DebugFunc(ctx context.Context, createMsg func() string, args ...any)
		Info(ctx context.Context, msg string, args ...any)
		Warn(ctx context.Context, msg string, args ...any)
		Error(ctx context.Context, msg string, args ...any)

		Log(ctx context.Context, level Level, message string, args ...any)
	}
)

// Debug - логирует сообщения на уровне LevelDebug.
func Debug(logger Logger, msg string, args ...any) {
	logger.Debug(context.Background(), msg, args...)
}

// DebugFunc - логирует сообщения на уровне LevelDebug с их отложенным созданием.
func DebugFunc(logger Logger, createMsg func() string, args ...any) {
	if !logger.Enabled(LevelDebug) {
		return
	}

	logger.Debug(context.Background(), createMsg(), args...)
}

// Info - логирует сообщения на уровне LevelInfo.
func Info(logger Logger, msg string, args ...any) {
	logger.Info(context.Background(), msg, args...)
}

// Warn - логирует сообщения на уровне LevelWarn.
func Warn(logger Logger, msg string, args ...any) {
	logger.Warn(context.Background(), msg, args...)
}

// Error - логирует сообщения на уровне LevelError.
func Error(logger Logger, msg string, args ...any) {
	logger.Error(context.Background(), msg, args...)
}

// Fatal - логирует ошибку и прекращает выполнение программы.
func Fatal(logger Logger, msg string, args ...any) {
	logger.Error(context.Background(), msg, args...)
	os.Exit(1) // //nolint:revive
}
