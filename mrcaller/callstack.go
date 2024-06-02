package mrcaller

import (
	"fmt"
)

type (
	// StackTrace - объект с уже сформированным стеком вызовов функций.
	StackTrace struct {
		items []StackItem
	}

	// StackItem - объект с уже сформированным стеком вызовов функций.
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

// FileLine - возвращает путь к файлу и номер строки кода,
// где расположена вызванная функция указанного элемента.
// Если i превысит кол-во элементов в стеке вызовов, то будет вызвана panic.
func (c *StackTrace) FileLine(i int) (file string, line int) {
	c.check(i)
	return c.items[i].File, c.items[i].Line
}

// Item - возвращает имя функции указанного элемента, включая данные из FileLine.
// Если i превысит кол-во элементов в стеке вызовов, то будет вызвана panic.
func (c *StackTrace) Item(i int) (name, file string, line int) {
	c.check(i)
	return c.items[i].Name, c.items[i].File, c.items[i].Line
}

func (c *StackTrace) check(i int) {
	if i < 0 || i >= len(c.items) {
		panic(fmt.Sprintf("index out of range [%d] with length %d", i, len(c.items)))
	}
}
