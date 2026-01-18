package conv

import (
	"strconv"
)

const (
	nilValue = "<NIL>"
)

// JSONValue - преобразовывает значение аргумента в строку для использования
// в JSON в качестве значения, строковые выражения дополнительно помещаются в двойные кавычки.
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
