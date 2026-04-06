package parse

import (
	"strconv"
	"strings"
)

const (
	typeUint64       = "Uint64"
	maxLenUint64     = 32
	maxLenUint64List = 256
)

// Uint64 - парсит строку в значение uint64.
// Если значение пустое и required=true, возвращает ошибку.
// Если значение пустое и required=false, возвращает 0.
func Uint64(value string, required bool) (uint64, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return 0, NewParamEmptyError(typeUint64)
		}

		return 0, nil
	}

	if len(value) > maxLenUint64 {
		return 0, NewParamLenMaxError(typeUint64, maxLenUint64)
	}

	item, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, NewParamIncorrectError(typeUint64, err)
	}

	return item, nil
}

// Uint64List - парсит строку с разделителями-запятыми в список uint64.
func Uint64List(value string) ([]uint64, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, nil
	}

	if len(value) > maxLenUint64List {
		return nil, NewParamLenMaxError(typeUint64, maxLenUint64List)
	}

	itemsTmp := strings.Split(value, ",")
	items := make([]uint64, 0, len(itemsTmp))

	for i := range itemsTmp {
		item, err := strconv.ParseUint(strings.TrimSpace(itemsTmp[i]), 10, 64)
		if err != nil {
			return nil, NewParamIncorrectError(typeUint64, err)
		}

		items = append(items, item)
	}

	return items, nil
}
