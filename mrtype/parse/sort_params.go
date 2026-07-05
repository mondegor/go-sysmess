package parse

import (
	"regexp"
	"strings"

	"github.com/mondegor/go-sysmess/mrtype"
	"github.com/mondegor/go-sysmess/mrtype/errors"
	"github.com/mondegor/go-sysmess/mrtype/sortdirection"
)

const (
	typeSortParamsColumn    = "SortParams.Column"
	typeSortParamsDirection = "SortParams.Direction"
	maxLenSortColumn        = 32
)

var regexpSorterColumn = regexp.MustCompile(`^[a-z]([a-zA-Z0-9]+)?[a-zA-Z0-9]$`)

// SortParams - парсит параметры сортировки из строк.
// Параметры:
//   - columnValue - имя колонки сортировки (допустимые символы: латинские буквы и цифры, начинается с буквы);
//   - directionValue - направление сортировки ("ASC" или "DESC", регистр не важен);
//
// Если columnValue пустое, возвращает пустые SortParams.
// По умолчанию направление сортировки - ASC.
func SortParams(columnValue, directionValue string) (mrtype.SortParams, error) {
	columnValue = strings.TrimSpace(columnValue)

	if columnValue == "" {
		return mrtype.SortParams{}, nil
	}

	if len(columnValue) > maxLenSortColumn {
		return mrtype.SortParams{}, errors.NewParamLenMaxError(typeSortParamsColumn, maxLenSortColumn)
	}

	if !regexpSorterColumn.MatchString(columnValue) {
		return mrtype.SortParams{}, errors.NewParamRegexpError(typeSortParamsColumn, regexpSorterColumn.String())
	}

	params := mrtype.SortParams{
		Column:    columnValue,
		Direction: sortdirection.ASC,
	}

	if directionValue != "" {
		sortDirection, err := sortdirection.Parse(strings.ToUpper(directionValue))
		if err != nil {
			return mrtype.SortParams{}, errors.NewParamIncorrectError(typeSortParamsDirection, err)
		}

		params.Direction = sortDirection
	}

	return params, nil
}
