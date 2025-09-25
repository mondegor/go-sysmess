package slog

import (
	"log/slog"
)

// Err - возвращает указанную ошибку в виде атрибута сообщения логгера.
func Err(value error) slog.Attr {
	return slog.Any("error", value)
}
