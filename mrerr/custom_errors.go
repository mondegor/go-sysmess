package mrerr

import (
	"strings"
)

type (
	// CustomErrorList - список пользовательских ошибок.
	CustomErrorList []*CustomError
)

// Error - возвращает список ошибок в виде строки.
func (l CustomErrorList) Error() string {
	var buf strings.Builder

	for i := 0; i < len(l); i++ {
		if i > 0 {
			buf.WriteString("\n")
		}

		buf.WriteString(l[i].Error())
	}

	return buf.String()
}
