package parse

import (
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/mrtype"
)

const (
	typeFloat64         = "Float64"
	typeRangeFloat64Min = "RangeFloat64.Min"
	typeRangeFloat64Max = "RangeFloat64.Max"
	maxLenFloat64       = 64
)

// Float64 - парсит строку в значение float64.
// Если значение пустое и required=true, возвращает ошибку.
// Если значение пустое и required=false, возвращает 0.
func Float64(value string, required bool) (float64, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return 0, NewParamEmptyError(typeFloat64)
		}

		return 0, nil
	}

	if len(value) > maxLenFloat64 {
		return 0, NewParamLenMaxError(typeFloat64, maxLenFloat64)
	}

	item, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, NewParamIncorrectError(typeFloat64, err)
	}

	return item, nil
}

// RangeFloat64 - парсит два строковых значения в RangeFloat64.
// Если maxValue меньше minValue и maxValue > 0, автоматически меняет их местами.
func RangeFloat64(minValue, maxValue string) (mrtype.RangeFloat64, error) {
	parsedMinValue, err := Float64(minValue, false)
	if err != nil {
		return mrtype.RangeFloat64{}, NewParamIncorrectError(typeRangeFloat64Min, err)
	}

	parsedMaxValue, err := Float64(maxValue, false)
	if err != nil {
		return mrtype.RangeFloat64{}, NewParamIncorrectError(typeRangeFloat64Max, err)
	}

	if parsedMaxValue > 0 && parsedMinValue > parsedMaxValue { // change
		return mrtype.RangeFloat64{
			Min: parsedMaxValue,
			Max: parsedMinValue,
		}, nil
	}

	return mrtype.RangeFloat64{
		Min: parsedMinValue,
		Max: parsedMaxValue,
	}, nil
}
