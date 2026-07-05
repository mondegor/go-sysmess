package mrtype

import (
	"github.com/mondegor/go-core/mrtype/sortdirection"
)

type (
	// CursorParams - параметры курсорной пагинации.
	// Используется для выборки части списка элементов с помощью курсора.
	CursorParams struct {
		// Value - значение курсора (последовательность и указатель на элемент начала выборки).
		Value string

		// Limit - максимальное количество элементов в выборке.
		Limit int
	}

	// PageParams - параметры страничной пагинации.
	// Используется для выборки части списка элементов с помощью смещения.
	PageParams struct {
		// Index - индекс страницы (нумерация с 0 или 1 в зависимости от контекста).
		Index int

		// Size - количество элементов на странице.
		Size int
	}

	// SortParams - параметры сортировки списка элементов.
	SortParams struct {
		// Column - имя колонки для сортировки.
		Column string

		// Direction - направление сортировки (ASC/DESC).
		Direction sortdirection.Enum
	}

	// ListSorter - интерфейс проверки допустимых полей сортировки.
	// Позволяет валидировать имя поля сортировки и получить параметры сортировки по умолчанию.
	ListSorter interface {
		HasColumn(name string) bool
		DefaultSort() SortParams
	}
)
