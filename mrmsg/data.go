package mrmsg

import (
	"fmt"
)

type (
	Data map[string]any
)

func (d Data) String() string {
	var buf []byte
	firstItem := true

	buf = append(buf, '{')

	for key, value := range d {
		if firstItem {
			firstItem = false
		} else {
			buf = append(buf, ',', ' ')
		}

		buf = append(buf, key...)
		buf = append(buf, ':', ' ')
		buf = append(buf, fmt.Sprintf("%v", value)...)
	}

	buf = append(buf, '}')

	return string(buf)
}
