package parse

import (
	"strconv"
	"strings"
)

const (
	typeRequiredBool = "RequiredBool"
	typeNullableBool = "NullableBool"
)

// RequiredBool - возвращает Bool значение из указанной строки.
// Если параметр пустой, то возвращается ошибка.
func RequiredBool(value string) (bool, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return false, NewParamEmptyError(typeRequiredBool)
	}

	item, err := strconv.ParseBool(value)
	if err != nil {
		return false, NewParamIncorrectError(typeRequiredBool, err)
	}

	return item, nil
}

// NullableBool - возвращает Bool значение из указанной строки.
// Если параметр пустой, то возвращается nil.
func NullableBool(value string) (*bool, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, nil //nolint:nilnil
	}

	item, err := strconv.ParseBool(value)
	if err != nil {
		return nil, NewParamIncorrectError(typeNullableBool, err)
	}

	return &item, nil
}
