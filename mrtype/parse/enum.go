package parse

import (
	"regexp"
	"strings"
)

const (
	typeEnum       = "Enum"
	maxLenEnum     = 64
	maxLenEnumList = 256
)

var regexpEnum = regexp.MustCompile(`^[A-Z]([A-Z0-9_]+)?[A-Z0-9]$`)

// Enum - возвращает Enum значение из указанной строки.
// Если параметр пустой, то в зависимости от required возвращается пустая строка или ошибка.
func Enum(value string, required bool) (string, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return "", NewParamEmptyError(typeEnum)
		}

		return "", nil
	}

	if len(value) > maxLenEnum {
		return "", NewParamLenMaxError(typeEnum, maxLenEnum)
	}

	value = strings.ToUpper(value)

	if !regexpEnum.MatchString(value) {
		return "", NewParamRegexpError(typeEnum, regexpEnum.String())
	}

	return value, nil
}

// EnumList - возвращает массив Enum значений из указанной строки.
// Если параметр пустой, то возвращается пустой массив.
func EnumList(value string) ([]string, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, nil
	}

	if len(value) > maxLenEnumList {
		return nil, NewParamLenMaxError(typeEnum, maxLenEnumList)
	}

	items := strings.Split(strings.ToUpper(value), ",")

	for i, item := range items {
		item = strings.TrimSpace(item)

		if !regexpEnum.MatchString(item) {
			return nil, NewParamRegexpError(typeEnum, regexpEnum.String())
		}

		items[i] = item
	}

	return items, nil
}
