package parse

import (
	"regexp"
	"strings"

	"github.com/mondegor/go-sysmess/mrtype"
	"github.com/mondegor/go-sysmess/mrtype/sortdirection"
)

const (
	typeSortParamsField     = "SortParams.Field"
	typeSortParamsDirection = "SortParams.Direction"
	maxLenSortField         = 32
)

var regexpSorterField = regexp.MustCompile(`^[a-z]([a-zA-Z0-9]+)?[a-zA-Z0-9]$`)

// SortParams - возвращает SortParams из строковых параметров по указанным ключам.
func SortParams(fieldValue, directionValue string) (mrtype.SortParams, error) {
	fieldValue = strings.TrimSpace(fieldValue)

	if fieldValue == "" {
		return mrtype.SortParams{}, nil
	}

	if len(fieldValue) > maxLenSortField {
		return mrtype.SortParams{}, NewParamLenMaxError(typeSortParamsField, maxLenSortField)
	}

	if !regexpSorterField.MatchString(fieldValue) {
		return mrtype.SortParams{}, NewParamRegexpError(typeSortParamsField, regexpSorterField.String())
	}

	params := mrtype.SortParams{
		Column:    fieldValue,
		Direction: sortdirection.ASC,
	}

	if directionValue != "" {
		sortDirection, err := sortdirection.Parse(strings.ToUpper(directionValue))
		if err != nil {
			return mrtype.SortParams{}, NewParamIncorrectError(typeSortParamsDirection, err)
		}

		params.Direction = sortDirection
	}

	return params, nil
}
