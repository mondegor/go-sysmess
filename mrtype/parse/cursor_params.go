package parse

import (
	"math"

	"github.com/mondegor/go-sysmess/mrtype"
)

const (
	typeCursorParamsValue = "CursorParams.Value"
	typeCursorParamsLimit = "CursorParams.Limit"
)

// CursorParams - возвращает CursorParams из строковых параметров по указанным ключам.
func CursorParams(cursorValue, limitValue string) (mrtype.CursorParams, error) {
	parsedValue, err := String(cursorValue, false)
	if err != nil {
		return mrtype.CursorParams{}, NewParamIncorrectError(typeCursorParamsValue, err)
	}

	parsedLimit, err := Uint64(limitValue, false)
	if err != nil {
		return mrtype.CursorParams{}, NewParamIncorrectError(typeCursorParamsLimit, err)
	}

	if parsedLimit > math.MaxInt {
		return mrtype.CursorParams{}, NewParamMaxValueError(typeCursorParamsLimit, math.MaxInt)
	}

	return mrtype.CursorParams{
		Value: parsedValue,
		Limit: int(parsedLimit),
	}, nil
}
