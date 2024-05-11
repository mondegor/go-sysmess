package mrcaller

import (
	"strings"
)

type (
	// CallStack объект с уже сформированным стеком вызовов функций.
	CallStack struct {
		stack  []uintptr
		prefix string
	}
)

// Empty - проверяется, что CallStack пустой или нет.
func (c *CallStack) Empty() bool {
	return len(c.stack) == 0
}

// NewIterator - возвращает итератор для обхода сформированного CallStack.
func (c *CallStack) NewIterator() CallStackIterator {
	return CallStackIterator{
		cs: c,
	}
}

func (c *CallStack) callStackItem(pos int) (item CallStackItem, ok bool) {
	if pos >= len(c.stack) {
		return CallStackItem{}, false
	}

	frame := runtimeFrame((c.stack)[pos])
	file, line := frame.FileLine()

	if isBreak(file) {
		return CallStackItem{}, false
	}

	return CallStackItem{
		frame: frame,
		file:  c.shortenFilePath(file),
		line:  line,
	}, true
}

func (c *CallStack) shortenFilePath(file string) string {
	const (
		replacer     = "..."
		minPrefixLen = len(replacer) + 1
	)

	if len(c.prefix) >= minPrefixLen && strings.HasPrefix(file, c.prefix) {
		return replacer + file[len(c.prefix)-1:]
	}

	return file
}

func isBreak(file string) bool {
	return file == "" || strings.HasSuffix(file, callStackBreak)
}
