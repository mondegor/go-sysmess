package casttype

import (
	"time"

	"golang.org/x/exp/constraints"
)

// SliceToAnySlice - возвращает элементы преобразованные в слайс пустых интерфейсов.
func SliceToAnySlice[Type any](values []Type) []any {
	vs := make([]any, len(values))

	for i, v := range values {
		vs[i] = v
	}

	return vs
}

// BoolToNumber - возвращает преобразованный bool к Number.
func BoolToNumber[Number constraints.Integer | constraints.Float](value bool) Number {
	if value {
		return 1
	}

	return 0
}

// BoolToPointer - возвращает преобразованный bool к его указателю.
// Если свойство required не указано или равно false, то вместо значения false будет возвращено nil.
func BoolToPointer(value bool, required ...bool) *bool {
	if isNullable(required) && !value {
		return nil
	}

	return &value
}

// NumberToPointer - возвращает преобразованный number к его указателю.
// Если свойство required не указано или равно false, то вместо нулевого значения будет возвращено nil.
func NumberToPointer[Number constraints.Integer | constraints.Float](value Number, required ...bool) *Number {
	if isNullable(required) && value == 0 {
		return nil
	}

	return &value
}

// StringToPointer - возвращает преобразованную строку к его указателю.
// Если свойство required не указано или равно false, то вместо пустого значения будет возвращено nil.
func StringToPointer(value string, required ...bool) *string {
	if isNullable(required) && value == "" {
		return nil
	}

	return &value
}

// TimeToPointer - возвращает преобразованное время к его указателю.
// Если свойство required не указано или равно false, то вместо нулевого значения будет возвращено nil.
func TimeToPointer(value time.Time, required ...bool) *time.Time {
	if isNullable(required) && value.IsZero() {
		return nil
	}

	return &value
}

func isNullable(required []bool) bool {
	return len(required) == 0 || !required[0]
}
