package mrcaller

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

// FileLine - возвращает путь к файлу и номер строки кода,
// где расположена вызванная функция этого элемента.
func (f runtimeFrame) FileLine() (file string, line int) {
	fn := runtime.FuncForPC(f.pc())
	if fn == nil {
		return "???", 0
	}

	return fn.FileLine(f.pc())
}

func (f runtimeFrame) pc() uintptr {
	return uintptr(f) - 1
}
