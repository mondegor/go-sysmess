package parse

import (
	"math"

	"github.com/mondegor/go-sysmess/mrtype"
	"github.com/mondegor/go-sysmess/mrtype/errors"
)

const (
	typePageParamsIndex = "PageParams.Index"
	typePageParamsSize  = "PageParams.Size"
)

// PageParams - парсит параметры страничной пагинации из строк.
// Параметры:
//   - indexValue - индекс страницы;
//   - sizeValue - количество элементов на странице;
//
// Возвращает ошибку, если значения превышают math.MaxInt.
func PageParams(indexValue, sizeValue string) (mrtype.PageParams, error) {
	parsedIndex, err := Uint64(indexValue, false)
	if err != nil {
		return mrtype.PageParams{}, errors.NewParamIncorrectError(typePageParamsIndex, err)
	}

	parsedSize, err := Uint64(sizeValue, false)
	if err != nil {
		return mrtype.PageParams{}, errors.NewParamIncorrectError(typePageParamsSize, err)
	}

	if parsedSize > math.MaxInt {
		return mrtype.PageParams{}, errors.NewParamMaxValueError(typePageParamsSize, math.MaxInt)
	}

	if parsedIndex > math.MaxInt {
		return mrtype.PageParams{}, errors.NewParamMaxValueError(typePageParamsIndex, math.MaxInt)
	}

	return mrtype.PageParams{
		Index: int(parsedIndex),
		Size:  int(parsedSize),
	}, nil
}
