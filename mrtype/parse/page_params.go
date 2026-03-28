package parse

import (
	"math"

	"github.com/mondegor/go-sysmess/mrtype"
)

const (
	typePageParamsIndex = "PageParams.Index"
	typePageParamsSize  = "PageParams.Size"
)

// PageParams - возвращает PageParams из строковых параметров по указанным ключам.
func PageParams(indexValue, sizeValue string) (mrtype.PageParams, error) {
	parsedIndex, err := Uint64(indexValue, false)
	if err != nil {
		return mrtype.PageParams{}, NewParamIncorrectError(typePageParamsIndex, err)
	}

	parsedSize, err := Uint64(sizeValue, false)
	if err != nil {
		return mrtype.PageParams{}, NewParamIncorrectError(typePageParamsSize, err)
	}

	if parsedSize > math.MaxInt {
		return mrtype.PageParams{}, NewParamMaxValueError(typePageParamsSize, math.MaxInt)
	}

	if parsedIndex > math.MaxInt {
		return mrtype.PageParams{}, NewParamMaxValueError(typePageParamsIndex, math.MaxInt)
	}

	return mrtype.PageParams{
		Index: int(parsedIndex),
		Size:  int(parsedSize),
	}, nil
}
