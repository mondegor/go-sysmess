package stacktrace

import (
	"strconv"
	"strings"
)

// FindBottomBound - возвращает функцию, которая определяет,
// является ли указанное название функции нижней границей стека вызовов.
// После этой границы дальнейший стек вызовов не представляет интереса.
// Как правило, далее идёт вызов общих пакетов и сторонних библиотечных функций.
// В массиве bounds указываются все названия пакетов, которые можно опустить,
// стек вызовов будет срезан по самому верхнему из них.
func FindBottomBound(bounds []string) func(funcName string) bool {
	boundMap := make(map[string]bool, len(bounds))
	for _, item := range bounds {
		boundMap[item] = true
	}

	return func(funcName string) bool {
		return boundMap[cutPostfix(funcName)]
	}
}

// Примеры:
// - dir1/dir2/func.Name.X -> dir1/dir2/func
// - dir1/dir2/func_Name_X -> dir1/dir2.
func cutPostfix(value string) string {
	if i := strings.LastIndexByte(value, '/'); i >= 0 {
		if j := strings.IndexByte(value[i+1:], '.'); j >= 0 {
			return value[0 : i+j+1]
		}

		return value[0:i]
	}

	return value
}

// ToStrings - преобразует итератор стека вызовов в слайс строк.
// Формат каждой строки: "[имя_функции] файл:строка" или "файл:строка" (если имя функции пусто).
func ToStrings(stackIterator func() (index int, name, file string, line int)) []string {
	if stackIterator == nil {
		return nil
	}

	var (
		buf  strings.Builder
		list []string
	)

	for {
		index, name, file, line := stackIterator()
		if index < 0 {
			break
		}

		if list == nil {
			buf.Grow((len(name) + len(file)) * stackTraceMaxDepth / 8)
			list = make([]string, 0, stackTraceMaxDepth/8)
		}

		if name != "" {
			buf.WriteByte('[')
			buf.WriteString(name)
			buf.WriteString("] ")
		}

		buf.WriteString(file)
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))

		list = append(list, buf.String())
		buf.Reset()
	}

	return list
}
