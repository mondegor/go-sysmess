package enums

import (
	"fmt"
)

// Направление сортировки.
const (
	SortDirectionASC  SortDirection = iota // по возрастанию
	SortDirectionDESC                      // по убыванию
)

const (
	enumNameSortDirection = "SortDirection"
)

type (
	// SortDirection - направление сортировки.
	SortDirection uint8
)

var (
	sortDirectionName = map[SortDirection]string{ //nolint:gochecknoglobals
		SortDirectionASC:  "ASC",
		SortDirectionDESC: "DESC",
	}

	sortDirectionValue = map[string]SortDirection{ //nolint:gochecknoglobals
		"ASC":  SortDirectionASC,
		"DESC": SortDirectionDESC,
	}
)

// String - возвращает значение в виде строки.
func (e SortDirection) String() string {
	return sortDirectionName[e]
}

// ParseSortDirection - парсит указанное значение и если оно валидно, то устанавливает его числовое значение.
func ParseSortDirection(value string) (SortDirection, error) {
	if parsedValue, ok := sortDirectionValue[value]; ok {
		return parsedValue, nil
	}

	return 0, fmt.Errorf("key '%s' is not found in source '%s'", value, enumNameSortDirection)
}
