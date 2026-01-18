package casttype

import (
	"strconv"
)

// UintToPhone - возвращает число в виде строки с "+" перед ним.
func UintToPhone(value uint64) string {
	if value > 0 {
		return "+" + strconv.FormatUint(value, 10)
	}

	return ""
}
