package mrtype

import (
	"github.com/mondegor/go-sysmess/mrtype/sortdirection"
)

type (
	// CursorParams - параметры для выборки части списка элементов с помощью курсора.
	CursorParams struct {
		Value string // название последовательности и указатель на её элемент, с которого должны быть выбраны данные
		Limit int    // максимальное количество элементов в выборке
	}

	// PageParams - параметры для выборки части списка элементов с помощью смещения.
	PageParams struct {
		Index int // pageIndex, индекс страницы
		Size  int // pageSize, количество элементов на страницу
	}

	// SortParams - параметры для сортировки списка элементов по указанному полю.
	SortParams struct {
		Column    string             // sortColumn
		Direction sortdirection.Enum // sortDirection
	}

	// ListSorter - контролирует поля участвующие в сортировке.
	ListSorter interface {
		HasColumn(name string) bool
		DefaultSort() SortParams
	}
)
