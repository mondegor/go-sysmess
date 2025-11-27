package sortdirection

import (
	"fmt"
)

// Возможные направления сортировки.
const (
	ASC  Enum = iota // по возрастанию
	DESC             // по убыванию
)

const (
	enumName = "SortDirection"
)

type (
	// Enum - перечисление направлений сортировки.
	Enum uint8
)

var (
	enumKeys = map[Enum]string{ //nolint:gochecknoglobals
		ASC:  "ASC",
		DESC: "DESC",
	}

	enumValues = map[string]Enum{ //nolint:gochecknoglobals
		"ASC":  ASC,
		"DESC": DESC,
	}
)

// String - возвращает значение в виде строки.
func (e Enum) String() string {
	return enumKeys[e]
}

// Parse - парсит указанное значение и если оно валидно, то устанавливает его числовое значение.
func Parse(value string) (Enum, error) {
	if parsedValue, ok := enumValues[value]; ok {
		return parsedValue, nil
	}

	return 0, fmt.Errorf("key is not found in source (source='%s', key='%s')", enumName, value)
}
