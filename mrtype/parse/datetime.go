package parse

import (
	"strings"
	"time"
)

const (
	typeDateTime   = "DateTime"
	maxLenDateTime = 64
)

// DateTime - возвращает time.Time значение из указанной строки.
// Значение строкового параметра должно быть указано в формате RFC3339.
// Если параметр пустой, то в зависимости от required возвращается нулевое время или ошибка.
func DateTime(value string, required bool) (time.Time, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return time.Time{}, NewParamEmptyError(typeDateTime)
		}

		return time.Time{}, nil
	}

	if len(value) > maxLenDateTime {
		return time.Time{}, NewParamLenMaxError(typeDateTime, maxLenDateTime)
	}

	item, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return time.Time{}, NewParamIncorrectError(typeDateTime, err)
	}

	return item, nil
}
