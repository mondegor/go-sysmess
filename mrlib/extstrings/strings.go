package extstrings

import (
	"strings"
)

// InArray - сообщает есть ли указанная строка в массиве строк.
func InArray(needle string, haystack []string) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}

// TrimBeforeSep - возвращает строку без её начальной части продолжающейся до сепаратора включая этот сепаратор.
// Если сепаратор не был найден, то возвращается исходная строка.
func TrimBeforeSep(s string, sep byte) string {
	if i := strings.IndexByte(s, sep); i >= 0 {
		return s[i+1:]
	}

	return s
}
