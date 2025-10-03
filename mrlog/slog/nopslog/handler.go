package nopslog

import (
	"context"
	"log/slog"
)

type handler struct{}

// Enabled - реализация метода для интерфейса slog.Handler.
func (handler) Enabled(context.Context, slog.Level) bool {
	return false
}

// Handle - реализация метода для интерфейса slog.Handler.
func (handler) Handle(context.Context, slog.Record) error {
	return nil
}

// WithAttrs - реализация метода для интерфейса slog.Handler.
func (handler) WithAttrs([]slog.Attr) slog.Handler {
	return handler{}
}

// WithGroup - реализация метода для интерфейса slog.Handler.
func (handler) WithGroup(string) slog.Handler {
	return handler{}
}
