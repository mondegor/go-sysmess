package custom

import (
	"strings"
)

type (
	// ListError - слайс пользовательских ошибок с уточнённым кодом.
	// Реализует интерфейс error, объединяя все ошибки через перевод строки.
	// Используется для агрегации нескольких ошибок валидации.
	ListError []Error
)

// Error - возвращает строковое представление списка ошибок.
func (e ListError) Error() string {
	var buf strings.Builder

	for i := 0; i < len(e); i++ {
		if i > 0 {
			buf.WriteByte('\n')
		}

		buf.WriteString((e)[i].Error())
	}

	return buf.String()
}
