package parse

import (
	"math"

	"github.com/mondegor/go-core/mrtype"
	"github.com/mondegor/go-core/mrtype/errors"
)

const (
	typeCursorParamsValue = "CursorParams.Value"
	typeCursorParamsLimit = "CursorParams.Limit"
)

// CursorParams - парсит параметры курсорной пагинации из строк.
// Параметры:
//   - cursorValue - значение курсора (может быть пустым для первой страницы);
//   - limitValue - максимальное количество элементов.
//
// Возвращает ошибку, если limit превышает math.MaxInt.
func CursorParams(cursorValue, limitValue string) (mrtype.CursorParams, error) {
	parsedValue, err := String(cursorValue, false)
	if err != nil {
		return mrtype.CursorParams{}, errors.NewParamIncorrectError(typeCursorParamsValue, err)
	}

	parsedLimit, err := Uint64(limitValue, false)
	if err != nil {
		return mrtype.CursorParams{}, errors.NewParamIncorrectError(typeCursorParamsLimit, err)
	}

	if parsedLimit > math.MaxInt {
		return mrtype.CursorParams{}, errors.NewParamMaxValueError(typeCursorParamsLimit, math.MaxInt)
	}

	return mrtype.CursorParams{
		Value: parsedValue,
		Limit: int(parsedLimit),
	}, nil
}
