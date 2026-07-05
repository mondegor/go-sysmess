package mrstorage

import (
	"strconv"
)

// ToSQL - преобразует часть SQL-запроса в строку.
// Если part равен nil, возвращает пустую строку.
// Игнорирует аргументы, возвращает только SQL-выражение.
func ToSQL(part SQLPart) string {
	if part == nil {
		return ""
	}

	sql, _ := part.ToSQL()

	return sql
}

// NonZeroLimit - формирует SQL-конструкцию LIMIT с гарантированным минимумом в одну запись.
// Если value меньше 1, лимит принудительно устанавливается равным 1,
// что исключает формирование некорректного выражения (LIMIT 0 или отрицательный лимит).
func NonZeroLimit(value int) string {
	if value < 1 {
		return " LIMIT 1"
	}

	return " LIMIT " + strconv.Itoa(value)
}
