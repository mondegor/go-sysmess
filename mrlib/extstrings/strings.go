package extstrings

import (
	"strings"
)

// TrimBeforeSep - возвращает строку без её начальной части продолжающейся до сепаратора включая этот сепаратор.
// Если сепаратор не был найден, то возвращается исходная строка.
func TrimBeforeSep(s string, sep byte) string {
	if i := strings.IndexByte(s, sep); i >= 0 {
		return s[i+1:]
	}

	return s
}
