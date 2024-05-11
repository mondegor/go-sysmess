package mrcaller

type (
	// CallStackIterator - итератор для объекта CallStack.
	CallStackIterator struct {
		cs  *CallStack
		pos int
	}
)

// Next - возвращает номер очередного элемента и сам элемент объекта CallStack.
// Если номер элемента равен 0, то итератор завершил свою работу (при этом в item содержится пустой элемент).
func (it *CallStackIterator) Next() (number int, item CallStackItem) {
	item, ok := it.cs.callStackItem(it.pos + 1)
	if !ok {
		return 0, CallStackItem{}
	}

	it.pos++

	return it.pos, item
}
