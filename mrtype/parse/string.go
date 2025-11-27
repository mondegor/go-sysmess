package parse

import (
	"strings"
)

const (
	typeString        = "String"
	maxLenString      = 256
	maxLenStringsList = 2048
)

// String - возвращает строковое значение из указанной строки.
// Если параметр пустой, то в зависимости от required возвращается пустая строка или ошибка.
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

// StringList - возвращает массив строковых значений из указанной строки.
// Если параметр пустой, то возвращается пустой массив.
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
