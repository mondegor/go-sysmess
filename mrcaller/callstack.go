package mrcaller

import (
	"fmt"
)

type (
	// StackTrace - объект с уже сформированным стеком вызовов функций.
	StackTrace struct {
		items []StackItem
	}

	// StackItem - элемент объекта StackTrace с вызовом конкретной функции.
	StackItem struct {
		Name string
		File string
		Line int
	}
)

// Count - возвращается количество элементов в стеке вызовов.
func (c *StackTrace) Count() int {
	return len(c.items)
}

// Item - возвращает имя функции указанного элемента, путь к файлу и номер строки кода.
// Если i превысит кол-во элементов в стеке вызовов, то будет вызвана panic.
func (c *StackTrace) Item(i int) (name, file string, line int) {
	if i < 0 || i >= len(c.items) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, len(c.items)))
	}

	return c.items[i].Name, c.items[i].File, c.items[i].Line
}
