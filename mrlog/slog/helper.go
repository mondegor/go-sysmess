package slog

import (
	"log/slog"

	"github.com/mondegor/go-sysmess/mrlog"
)

// ReplaceArg - функция для ранней замены переданного в логгер аргумента.
func ReplaceArg(value any) (newValue any) {
	if attr, ok := value.(mrlog.Attr); ok {
		return slog.Any(attr.Key, attr.Value)
	}

	return value
}
