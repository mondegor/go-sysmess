package mrtype

import (
	"github.com/mondegor/go-sysmess/mrtype/enums"
)

type (
	// PageParams - параметры для выборки части списка элементов.
	PageParams struct {
		Index uint64 // pageIndex
		Size  uint64 // pageSize
	}

	// PageCursor - параметры для выборки части списка элементов.
	PageCursor struct {
		LastID uint64 // lastItemID
		Size   uint64 // pageSize
	}

	// SortParams - параметры для сортировки списка элементов по указанному полю.
	SortParams struct {
		FieldName string              // sortField
		Direction enums.SortDirection // sortDirection
	}

	// ListSorter - интерфейс для проверки полей, которые могут участвовать в сортировке.
	ListSorter interface {
		HasField(name string) bool
		DefaultSort() SortParams
	}
)
