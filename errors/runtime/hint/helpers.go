package hint

import (
	"strconv"
	"strings"
)

// Extract - возвращает дополнительные данные, которые ассоциированы с ошибкой.
func Extract(err error) Hint {
	if e, ok := err.(interface{ Hint() any }); ok {
		if ht, ok := e.Hint().(Hint); ok {
			return ht
		}
	}

	return nopHint{}
}

// DetailedError - детальная информация об ошибке.
func DetailedError(err error) string {
	var buf strings.Builder

	buf.WriteString(err.Error())

	iterator := Extract(err).StackTraceIterator()

	for {
		index, name, file, line := iterator()
		if index < 0 {
			break
		}

		if index == 0 {
			buf.WriteString(": ")
		} else {
			buf.WriteString(" | ")
		}

		if name != "" {
			buf.WriteByte('[')
			buf.WriteString(name)
			buf.WriteString("] ")
		}

		buf.WriteString(file)
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))
	}

	return buf.String()
}
