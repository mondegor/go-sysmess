package mrerr

import (
	"strings"
)

type (
	// CustomErrors - список пользовательских ошибок.
	CustomErrors []*CustomError
)

// Error - возвращает список ошибок в виде строки.
func (l CustomErrors) Error() string {
	var buf strings.Builder

	for i := 0; i < len(l); i++ {
		if i > 0 {
			buf.WriteString("\n")
		}

		buf.WriteString(l[i].Error())
	}

	return buf.String()
}
