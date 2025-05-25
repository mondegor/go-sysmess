package stacktrace

import (
	"runtime"
)

type runtimeFrame uintptr

// Name - возвращает имя функции этого элемента, если оно известно.
func (f runtimeFrame) Name() string {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown"
	}

	return fn.Name()
}

// Source - возвращает имя функции, путь к файлу и номер строки кода,
// где расположена вызванная функция.
func (f runtimeFrame) Source() (function, file string, line int) {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "unknown", "???", 0
	}

	file, line = fn.FileLine(f.pc())

	return fn.Name(), file, line
}

func (f runtimeFrame) pc() uintptr {
	return uintptr(f) - 1
}
