package parser

import (
	"fmt"
	"time"
)

// ParseDateTimeFormat - возвращает формат времени предназначенный
// для использования Time.Format при логировании.
func ParseDateTimeFormat(value string) (string, error) {
	switch value {
	case "RFC3339":
		return time.RFC3339, nil
	case "RFC3339Nano":
		return time.RFC3339Nano, nil
	case "DateTime":
		return time.DateTime, nil
	case "TimeOnly":
		return time.TimeOnly, nil
	case "Kitchen":
		return time.Kitchen, nil
	}

	return time.RFC3339, fmt.Errorf("value '%s' is not found in parser.ParseDateTimeFormat()", value)
}
