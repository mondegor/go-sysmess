package parse

import (
	"strings"
)

const (
	typeString        = "String"
	maxLenString      = 256
	maxLenStringsList = 2048
)

// String - возвращает строковое значение после обрезки пробелов.
// Если значение пустое и required=true, возвращает ошибку.
// Если значение пустое и required=false, возвращает пустую строку.
func String(value string, required bool) (string, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return "", NewParamEmptyError(typeString)
		}

		return "", nil
	}

	if len(value) > maxLenString {
		return "", NewParamLenMaxError(typeString, maxLenString)
	}

	return value, nil
}

// StringList - парсит строку с разделителями-запятыми в список строк.
// Обрезает пробелы вокруг каждого элемента.
func StringList(value string) ([]string, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, nil
	}

	if len(value) > maxLenStringsList {
		return nil, NewParamLenMaxError(typeString, maxLenStringsList)
	}

	items := strings.Split(value, ",")

	for i, item := range items {
		items[i] = strings.TrimSpace(item)
	}

	return items, nil
}
