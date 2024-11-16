package mrmsg

import (
	"sort"
	"strings"
)

type (
	// Data - произвольные данные с возможностью их отображения в виде строки.
	Data map[string]any
)

// ValueString - возвращает значение в виде строки для указанного ключа.
// Если ключ не найден, то возвращается пустая строка.
func (d Data) ValueString(key string) string {
	if val, ok := d[key]; ok {
		return ToString(val)
	}

	return ""
}

// String - преобразовывает данные в строку.
func (d Data) String() string {
	// предварительная сортировка ключей т.к. map не гарантирует их порядок
	keys := make([]string, 0, len(d))
	for k := range d {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var buf strings.Builder

	buf.WriteByte('{')

	firstItem := true
	for _, key := range keys {
		if firstItem {
			firstItem = false
		} else {
			buf.WriteString(", ")
		}

		buf.WriteString(key)
		buf.WriteString(": ")
		buf.WriteString(ToString(d[key]))
	}

	buf.WriteByte('}')

	return buf.String()
}
