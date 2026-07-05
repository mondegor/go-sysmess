package casttype

import (
	"strconv"
)

// UintToPhone - преобразует числовой телефонный номер в строку с префиксом "+".
// Возвращает пустую строку, если value равен 0.
func UintToPhone(value uint64) string {
	if value > 0 {
		return "+" + strconv.FormatUint(value, 10)
	}

	return ""
}
