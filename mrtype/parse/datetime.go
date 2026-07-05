package parse

import (
	"strings"
	"time"

	"github.com/mondegor/go-core/mrtype"
	"github.com/mondegor/go-core/mrtype/errors"
)

const (
	typeDateTime          = "DateTime"
	typeRangeDateTimeFrom = "DateTime.From"
	typeRangeDateTimeTo   = "DateTime.To"
	maxLenDateTime        = 64
)

// DateTime - парсит строку в time.Time в формате RFC3339.
// Если значение пустое и required=true, возвращает ошибку.
// Если значение пустое и required=false, возвращает нулевое время.
func DateTime(value string, required bool) (time.Time, error) {
	return dateTime(value, time.RFC3339, required)
}

// Date - парсит строку в time.Time в формате DateOnly ("2006-01-02").
// Если значение пустое и required=true, возвращает ошибку.
// Если значение пустое и required=false, возвращает нулевое время.
func Date(value string, required bool) (time.Time, error) {
	return dateTime(value, time.DateOnly, required)
}

func dateTime(value, format string, required bool) (time.Time, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		if required {
			return time.Time{}, errors.NewParamEmptyError(typeDateTime)
		}

		return time.Time{}, nil
	}

	if len(value) > maxLenDateTime {
		return time.Time{}, errors.NewParamLenMaxError(typeDateTime, maxLenDateTime)
	}

	item, err := time.Parse(format, value)
	if err != nil {
		return time.Time{}, errors.NewParamIncorrectError(typeDateTime, err)
	}

	return item, nil
}

// RangeDateTime - парсит два строковых значения в RangeDateTime в формате RFC3339.
// Если toValue меньше fromValue, автоматически меняет их местами.
func RangeDateTime(fromValue, toValue string) (mrtype.RangeDateTime, error) {
	return rangeDateTime(fromValue, toValue, time.RFC3339)
}

// RangeDate - парсит два строковых значения в RangeDateTime в формате DateOnly.
// Если toValue меньше fromValue, автоматически меняет их местами.
func RangeDate(fromValue, toValue string) (mrtype.RangeDateTime, error) {
	return rangeDateTime(fromValue, toValue, time.DateOnly)
}

func rangeDateTime(fromValue, toValue, format string) (mrtype.RangeDateTime, error) {
	parsedFromValue, err := dateTime(fromValue, format, false)
	if err != nil {
		return mrtype.RangeDateTime{}, errors.NewParamIncorrectError(typeRangeDateTimeFrom, err)
	}

	parsedToValue, err := dateTime(toValue, format, false)
	if err != nil {
		return mrtype.RangeDateTime{}, errors.NewParamIncorrectError(typeRangeDateTimeTo, err)
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
