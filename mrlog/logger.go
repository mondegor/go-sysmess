package mrlog

import (
	"context"
	"os"

	"github.com/mondegor/go-sysmess/mrlog/level"
)

type (
	// Logger - интерфейс логирования сообщений и ошибок с использованием контекста.
	Logger interface {
		Debug(ctx context.Context, msg string, args ...any)
		DebugFunc(ctx context.Context, createMsg func() string, args ...any)
		Info(ctx context.Context, msg string, args ...any)
		Warn(ctx context.Context, msg string, args ...any)
		Error(ctx context.Context, msg string, args ...any)
	}
)

// WithAttrs - возвращает логгер с прикреплёнными к нему атрибутами.
func WithAttrs(l Logger, attrs ...any) Logger {
	if l, ok := l.(interface{ WithAttrs(attrs ...any) Logger }); ok {
		return l.WithAttrs(attrs...)
	}

	Warn(l, "Logger not supported method 'WithAttrs'")

	return l
}

// Enabled - информирует включён ли указанный уровень логирования.
func Enabled(l Logger, lvl level.Enum) bool {
	if ll, ok := l.(interface{ Enabled(lvl level.Enum) bool }); ok {
		return ll.Enabled(lvl)
	}

	Warn(l, "Logger not supported method 'Enabled'")

	return false
}

// Debug - логирует сообщения на уровне level.Debug без использования контекста.
func Debug(l Logger, msg string, args ...any) {
	l.Debug(context.Background(), msg, args...)
}

// DebugFunc - логирует сообщения на уровне level.Debug без использования контекста с их отложенным созданием сообщения.
func DebugFunc(l Logger, createMsg func() string, args ...any) {
	l.DebugFunc(context.Background(), createMsg, args...)
}

// Info - логирует сообщения на уровне level.Info без использования контекста.
func Info(l Logger, msg string, args ...any) {
	l.Info(context.Background(), msg, args...)
}

// Warn - логирует сообщения на уровне level.Warn без использования контекста.
func Warn(l Logger, msg string, args ...any) {
	l.Warn(context.Background(), msg, args...)
}

// Error - логирует сообщения на уровне level.Error без использования контекста.
func Error(l Logger, msg string, args ...any) {
	l.Error(context.Background(), msg, args...)
}

// FatalError - логирует сообщения на уровне level.Fatal без использования контекста и останавливает приложение.
func FatalError(l Logger, msg string, args ...any) {
	if ll, ok := l.(interface {
		Log(ctx context.Context, lvl level.Enum, message string, args ...any)
	}); ok {
		ll.Log(context.Background(), level.Fatal, msg, args...)
	} else {
		l.Error(context.Background(), "Fatal error: "+msg, args...)
	}

	os.Exit(1) // //nolint:revive
}
