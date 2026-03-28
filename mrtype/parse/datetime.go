package parse

import (
	"strings"
	"time"

	"github.com/mondegor/go-sysmess/mrtype"
)

const (
	typeDateTime          = "DateTime"
	typeRangeDateTimeFrom = "DateTime.From"
	typeRangeDateTimeTo   = "DateTime.To"
	maxLenDateTime        = 64
)

// DateTime - возвращает time.Time значение из указанной строки.
// Значение строкового параметра должно быть указано в формате RFC3339.
// Если параметр пустой, то в зависимости от required возвращается нулевое время или ошибка.
func DateTime(value string, required bool) (time.Time, error) {
	return dateTime(value, time.RFC3339, required)
}

// Date - возвращает time.Time значение из указанной строки.
// Значение строкового параметра должно быть указано в формате DateOnly.
// Если параметр пустой, то в зависимости от required возвращается нулевое время или ошибка.
func Date(value string, required bool) (time.Time, error) {
	return dateTime(value, time.DateOnly, required)
}

func dateTime(value, format string, required bool) (time.Time, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return time.Time{}, NewParamEmptyError(typeDateTime)
		}

		return time.Time{}, nil
	}

	if len(value) > maxLenDateTime {
		return time.Time{}, NewParamLenMaxError(typeDateTime, maxLenDateTime)
	}

	item, err := time.Parse(format, value)
	if err != nil {
		return time.Time{}, NewParamIncorrectError(typeDateTime, err)
	}

	return item, nil
}

// RangeDateTime - возвращает RangeDateTime из строковых параметров в формате RFC3339.
func RangeDateTime(fromValue, toValue string) (mrtype.RangeDateTime, error) {
	return rangeDateTime(fromValue, toValue, time.RFC3339)
}

// RangeDate - возвращает RangeDateTime из строковых параметров в формате DateOnly.
func RangeDate(fromValue, toValue string) (mrtype.RangeDateTime, error) {
	return rangeDateTime(fromValue, toValue, time.DateOnly)
}

func rangeDateTime(fromValue, toValue, format string) (mrtype.RangeDateTime, error) {
	parsedFromValue, err := dateTime(fromValue, format, false)
	if err != nil {
		return mrtype.RangeDateTime{}, NewParamIncorrectError(typeRangeDateTimeFrom, err)
	}

	parsedToValue, err := dateTime(toValue, format, false)
	if err != nil {
		return mrtype.RangeDateTime{}, NewParamIncorrectError(typeRangeDateTimeTo, err)
	}

	if !parsedToValue.IsZero() && parsedToValue.Before(parsedFromValue) { // change
		return mrtype.RangeDateTime{
			From: parsedToValue,
			To:   parsedFromValue,
		}, nil
	}

	return mrtype.RangeDateTime{
		From: parsedFromValue,
		To:   parsedToValue,
	}, nil
}
