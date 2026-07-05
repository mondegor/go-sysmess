package parse

import (
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/mrtype/errors"
)

const (
	typeRequiredBool = "RequiredBool"
	typeNullableBool = "NullableBool"
)

// RequiredBool - парсит строку в значение bool.
// Возвращает ошибку, если значение пустое или не является допустимым булевым значением.
func RequiredBool(value string) (bool, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return false, errors.NewParamEmptyError(typeRequiredBool)
	}

	item, err := strconv.ParseBool(value)
	if err != nil {
		return false, errors.NewParamIncorrectError(typeRequiredBool, err)
	}

	return item, nil
}

// NullableBool - парсит строку в указатель на bool.
// Возвращает nil, если значение пустое.
// Возвращает ошибку, если значение не является допустимым булевым значением.
func NullableBool(value string) (*bool, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, nil //nolint:nilnil
	}

	item, err := strconv.ParseBool(value)
	if err != nil {
		return nil, errors.NewParamIncorrectError(typeNullableBool, err)
	}

	return &item, nil
}
