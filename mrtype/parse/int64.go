package parse

import (
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/mrtype"
)

const (
	typeInt64         = "Int64"
	typeRangeInt64Min = "RangeInt64.Min"
	typeRangeInt64Max = "RangeInt64.Max"
	maxLenInt64       = 32
	maxLenInt64List   = 256
)

// Int64 - парсит строку в значение int64.
// Если значение пустое и required=true, возвращает ошибку.
// Если значение пустое и required=false, возвращает 0.
func Int64(value string, required bool) (int64, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return 0, NewParamEmptyError(typeInt64)
		}

		return 0, nil
	}

	if len(value) > maxLenInt64 {
		return 0, NewParamLenMaxError(typeInt64, maxLenInt64)
	}

	item, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, NewParamIncorrectError(typeInt64, err)
	}

	return item, nil
}

// Int64List - парсит строку с разделителями-запятыми в список int64.
func Int64List(value string) ([]int64, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return nil, nil
	}

	if len(value) > maxLenInt64List {
		return nil, NewParamLenMaxError(typeInt64, maxLenInt64List)
	}

	itemsTmp := strings.Split(value, ",")
	items := make([]int64, 0, len(itemsTmp))

	for i := range itemsTmp {
		item, err := strconv.ParseInt(strings.TrimSpace(itemsTmp[i]), 10, 64)
		if err != nil {
			return nil, NewParamIncorrectError(typeInt64, err)
		}

		items = append(items, item)
	}

	return items, nil
}

// RangeInt64 - парсит два строковых значения в RangeInt64.
// Если maxValue меньше minValue и maxValue > 0, автоматически меняет их местами.
func RangeInt64(minValue, maxValue string) (mrtype.RangeInt64, error) {
	parsedMinValue, err := Int64(minValue, false)
	if err != nil {
		return mrtype.RangeInt64{}, NewParamIncorrectError(typeRangeInt64Min, err)
	}

	parsedMaxValue, err := Int64(maxValue, false)
	if err != nil {
		return mrtype.RangeInt64{}, NewParamIncorrectError(typeRangeInt64Max, err)
	}

	if parsedMaxValue > 0 && parsedMinValue > parsedMaxValue { // change
		return mrtype.RangeInt64{
			Min: parsedMaxValue,
			Max: parsedMinValue,
		}, nil
	}

	return mrtype.RangeInt64{
		Min: parsedMinValue,
		Max: parsedMaxValue,
	}, nil
}
