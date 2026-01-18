package stacktrace

type (
	stackTrace []source

	// source - элемент объекта StackTrace с вызовом конкретной функции.
	source struct {
		function string
		file     string
		line     int
	}
)

// Iterator - возвращает итератор, для чтения элементов стека:
// имя функции указанного элемента, путь к файлу и номер строки кода.
// Если вернётся index = -1, значит стек прочитан полностью.
func (t stackTrace) Iterator() func() (index int, name, file string, line int) {
	i := -1

	return func() (index int, name, file string, line int) {
		i++

		if i >= len(t) {
			return -1, "", "", 0
		}

		return i, t[i].function, t[i].file, t[i].line
	}
}

// // Source - возвращает имя функции указанного элемента, путь к файлу и номер строки кода.
// // Если i превысит кол-во элементов в стеке вызовов, то будет вызвана panic.
// func (t stackTrace) Source(i int) (name, file string, line int) {
// 	if i < 0 || i >= len(t) {
// 		panic(fmt.Sprintf("index out of range [%d] with length %d", i, len(t)))
// 	}
//
// 	return t[i].function, t[i].file, t[i].line
// }
//
// // Count - возвращается количество элементов в стеке вызовов.
// func (t stackTrace) Count() int {
// 	return len(t)
// }
