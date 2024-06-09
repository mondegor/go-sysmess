package mrmsg

import (
	"fmt"
	"sort"
	"strings"
)

type (
	// Data - произвольные данные с возможностью их отображения в виде строки.
	Data map[string]any
)

// String - преобразовывает данные в строку.
func (d Data) String() string {
	var buf strings.Builder

	buf.WriteByte('{')

	// предварительная сортировка ключей т.к. map не гарантирует порядок
	keys := make([]string, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	firstItem := true
	for _, key := range keys {
		if firstItem {
			firstItem = false
		} else {
			buf.Write([]byte{',', ' '})
		}

		buf.WriteString(key)
		buf.Write([]byte{':', ' '})
		buf.WriteString(fmt.Sprintf("%v", d[key]))
	}

	buf.WriteByte('}')

	return buf.String()
}
