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

// Float64 - возвращает Float64 значение из указанной строки.
// Если параметр пустой, то в зависимости от required возвращается 0 или ошибка.
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

// RangeFloat64 - возвращает RangeFloat64 из строковых параметров.
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
