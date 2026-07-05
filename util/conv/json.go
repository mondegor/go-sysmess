package conv

import (
	"strconv"
)

const (
	nilValue = "<NIL>"
)

// JSONValue - преобразует значение в строковое представление для JSON.
// Строковые значения оборачиваются в двойные кавычки через strconv.Quote.
// nil-значения преобразуются в "null".
func JSONValue(value any) string {
	str, needQuote := toString(value)

	if needQuote {
		return strconv.Quote(str)
	}

	if str == nilValue {
		return "null"
	}

	return str
}
