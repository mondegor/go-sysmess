package mrcaller

import (
	"path/filepath"
	"runtime"
)

const (
	callStackMaxDepth = 32
	callStackBreak    = "src/runtime/proc.go"
)

type (
	// Caller - обёртка runtime.Callers для более удобного формирования CallStack.
	Caller struct {
		depth        int  // глубина отображения CallStack
		useShortPath bool // при true обрезает начало путей файлов

		// Префикс пути, который будет обрезан у имён файлов при выводе CallStack.
		// Если значение пустое, то используется родительский путь от первого файла в CallStack.
		// (используется только при useShortPath = true)
		rootPath string
	}
)

// New - создаётся объект Caller.
func New(opts ...CallerOption) *Caller {
	c := &Caller{}
	c.applyOptions(opts)

	return c
}

// CallStack - формирует CallStack для текущего вызова и возвращает его.
func (c *Caller) CallStack(skip int) CallStack {
	var pcs [callStackMaxDepth]uintptr
	n := runtime.Callers(skip+1, pcs[:])

	if n > c.depth+1 {
		n = c.depth + 1
	}

	return CallStack{
		stack:  pcs[0:n],
		prefix: c.filesPrefix(runtimeFrame(pcs[0])),
	}
}

func (c *Caller) applyOptions(opts []CallerOption) {
	for _, f := range opts {
		f(c)
	}
}

func (c *Caller) filesPrefix(fr runtimeFrame) string {
	if !c.useShortPath {
		return ""
	}

	file, _ := fr.FileLine()

	if isBreak(file) {
		return ""
	}

	if c.rootPath == "" {
		return filepath.ToSlash(filepath.Dir(filepath.Dir(file))) + "/"
	}

	return c.rootPath
}
