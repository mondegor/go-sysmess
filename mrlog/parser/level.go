package parser

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrlog"
)

// ParseLevel - парсит строку и возвращает указанный в ней
// уровень логирования или ошибку, если извлечь уровень не удалось.
func ParseLevel(value string) (mrlog.Level, error) {
	switch value {
	case "DEBUG":
		return mrlog.LevelDebug, nil
	case "INFO":
		return mrlog.LevelInfo, nil
	case "WARN":
		return mrlog.LevelWarn, nil
	case "ERROR":
		return mrlog.LevelError, nil
	}

	return mrlog.LevelError, fmt.Errorf("value '%s' is not found in parser.ParseLevel()", value)
}
