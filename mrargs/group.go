package mrargs

import (
	"sort"
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/mrtype"
)

type (
	// Group - группа данных с возможностью их отображения в виде строки.
	Group map[string]any
)

// ValueString - возвращает значение в виде строки для указанного ключа.
// Если ключ не найден, то возвращается пустая строка.
func (g Group) ValueString(key string) string {
	if val, ok := g[key]; ok {
		return mrtype.ToString(val)
	}

	return ""
}

// ToStringMap - конвертирует все значения в строку и возвращает результат.
// Если map неинициализированная, то новая map вернётся тоже неинициализированной.
func (g Group) ToStringMap() map[string]string {
	if g == nil {
		return nil
	}

	data := make(map[string]string, len(g))

	for k := range g {
		data[k] = mrtype.ToString(g[k])
	}

	return data
}

// // ToPairArgs - возвращает непрерывный список аргументов ключ/значение в виде массива
// // (без гарантии сохранения первоначальной упорядоченности ключей).
// func (g Group) ToPairArgs() []any {
// 	if len(g) == 0 {
// 		return nil
// 	}
//
// 	args := make([]any, 0, len(g)*2)
//
// 	for k := range g {
// 		args = append(args, k, g[k])
// 	}
//
// 	return args
// }

// String - возвращает группу данных преобразованную в строку.
func (g Group) String() string {
	if len(g) == 0 {
		return "{}"
	}

	// предварительная сортировка ключей т.к. map не гарантирует их порядок
	keys := make([]string, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	var buf strings.Builder

	buf.WriteByte('{')

	for i, key := range keys {
		if i > 0 {
			buf.WriteString(", ")
		}

		buf.WriteString(strconv.Quote(key))
		buf.WriteString(": ")
		buf.WriteString(mrtype.ToJSONValue(g[key]))
	}

	buf.WriteByte('}')

	return buf.String()
}
