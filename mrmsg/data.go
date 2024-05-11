package mrmsg

import (
	"fmt"
	"strings"
)

type (
	// Data - произвольные данные с возможностью их отображения в виде строки.
	Data map[string]any
)

// String - преобразовывает данные в строку.
func (d Data) String() string {
	var buf strings.Builder
	firstItem := true

	buf.WriteByte('{')

	for key, value := range d {
		if firstItem {
			firstItem = false
		} else {
			buf.Write([]byte{',', ' '})
		}

		buf.WriteString(key)
		buf.Write([]byte{':', ' '})
		buf.WriteString(fmt.Sprintf("%v", value))
	}

	buf.WriteByte('}')

	return buf.String()
}
