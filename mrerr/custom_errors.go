package mrerr

import (
	"strings"
)

type (
	// CustomErrors - список пользовательских ошибок.
	CustomErrors []*CustomError
)

// Error - возвращает список пользовательских ошибок в виде строки.
func (es CustomErrors) Error() string {
	var buf strings.Builder

	for i := 0; i < len(es); i++ {
		if i > 0 {
			buf.WriteByte('\n')
		}

		buf.WriteString((es)[i].Error())
	}

	return buf.String()
}
