package mrcaller

type (
	// CallStackItem - элемент сформированного CallStack.
	CallStackItem struct {
		frame runtimeFrame
		file  string
		line  int
	}
)

// Name - возвращает имя функции этого элемента, если оно известно.
func (it *CallStackItem) Name() string {
	return it.frame.Name()
}

// File - возвращает путь к файлу, где расположена вызванная функция этого элемента.
func (it *CallStackItem) File() string {
	return it.file
}

// Line - возвращает номер строки источника кода, где расположена вызванная функция этого элемента.
func (it *CallStackItem) Line() int {
	return it.line
}
