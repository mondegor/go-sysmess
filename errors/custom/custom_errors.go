package custom

import (
	"strings"
)

type (
	// ListError - список пользовательских ошибок с уточнённым кодом.
	ListError []Error
)

// Error - возвращает список ошибок в виде строки.
func (e ListError) Error() string {
	var buf strings.Builder

	for i := 0; i < len(e); i++ {
		if i > 0 {
			buf.WriteByte('\n')
		}

		buf.WriteString((e)[i].Error())
	}

	return buf.String()
}
