package casttype

import (
	"time"

	"golang.org/x/exp/constraints"
)

// BoolToNumber - преобразует bool в числовой тип: true → 1, false → 0.
func BoolToNumber[Number constraints.Integer | constraints.Float](value bool) Number {
	if value {
		return 1
	}

	return 0
}

// BoolToPointer - возвращает указатель на bool.
// Если required не указан или false, для значения false возвращается nil.
func BoolToPointer(value bool, required ...bool) *bool {
	if isNullable(required) && !value {
		return nil
	}

	return &value
}

// NumberToPointer - возвращает указатель на числовое значение.
// Если required не указан или false, для нулевого значения возвращается nil.
func NumberToPointer[Number constraints.Integer | constraints.Float](value Number, required ...bool) *Number {
	if isNullable(required) && value == 0 {
		return nil
	}

	return &value
}

// StringToPointer - возвращает указатель на строку.
// Если required не указан или false, для пустой строки возвращается nil.
func StringToPointer(value string, required ...bool) *string {
	if isNullable(required) && value == "" {
		return nil
	}

	return &value
}

// TimeToPointer - возвращает указатель на значение времени.
// Если required не указан или false, для нулевого времени возвращается nil.
func TimeToPointer(value time.Time, required ...bool) *time.Time {
	if isNullable(required) && value.IsZero() {
		return nil
	}

	return &value
}

func isNullable(required []bool) bool {
	return len(required) == 0 || !required[0]
}
