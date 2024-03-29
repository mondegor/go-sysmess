package mrerr

import (
	"strings"
)

type (
	CustomErrorList []*CustomError
)

func (l CustomErrorList) Error() string {
	var buf strings.Builder

	for i := 0; i < len(l); i++ {
		buf.WriteString(l[i].Error())
		buf.WriteString("\n")
	}

	return strings.TrimSpace(buf.String())
}
