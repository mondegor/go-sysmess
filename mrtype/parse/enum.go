package parse

import (
	"regexp"
	"strings"

	"github.com/mondegor/go-core/mrtype/errors"
)

const (
	typeEnum       = "Enum"
	maxLenEnum     = 64
	maxLenEnumList = 256
)

var regexpEnum = regexp.MustCompile(`^[A-Z]([A-Z0-9_]+)?[A-Z0-9]$`)

// Enum - парсит строку в значение перечисления.
// Допустимый формат: заглавные латинские буквы, цифры и подчёркивания (например: "MY_VALUE").
// Если значение пустое и required=true, возвращает ошибку.
// Если значение пустое и required=false, возвращает пустую строку.
func Enum(value string, required bool) (string, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return "", errors.NewParamEmptyError(typeEnum)
		}

		return "", nil
	}

	if len(value) > maxLenEnum {
		return "", errors.NewParamLenMaxError(typeEnum, maxLenEnum)
	}

	value = strings.ToUpper(value)

	if !regexpEnum.MatchString(value) {
		return "", errors.NewParamRegexpError(typeEnum, regexpEnum.String())
	}

	return value, nil
}

// EnumList - парсит строку с разделителями-запятыми в список значений перечисления.
// Каждый элемент приводится к верхнему регистру и проверяется по тому же формату, что и Enum.
func EnumList(value string) ([]string, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, nil
	}

	if len(value) > maxLenEnumList {
		return nil, errors.NewParamLenMaxError(typeEnum, maxLenEnumList)
	}

	items := strings.Split(strings.ToUpper(value), ",")

	for i, item := range items {
		item = strings.TrimSpace(item)

		if !regexpEnum.MatchString(item) {
			return nil, errors.NewParamRegexpError(typeEnum, regexpEnum.String())
		}

		items[i] = item
	}

	return items, nil
}
