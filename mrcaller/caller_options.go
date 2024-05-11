package mrcaller

import (
	"path/filepath"
	"strings"
)

type (
	// CallerOption - настройка объекта Caller.
	CallerOption func(c *Caller)
)

// CallerDepth - устанавливает опцию depth для Caller.
// Если value меньше 1 или больше callStackMaxDepth, то устанавливается 1.
func CallerDepth(value int) CallerOption {
	return func(c *Caller) {
		const minDepth = 1
		if value < minDepth || value > callStackMaxDepth {
			value = minDepth
		}

		c.depth = value
	}
}

// CallerUseShortPath - устанавливает опцию useShortPath для Caller.
func CallerUseShortPath(value bool) CallerOption {
	return func(c *Caller) { c.useShortPath = value }
}

// CallerRootPath - устанавливает опцию rootPath для Caller.
// Если value не пустое, то его конец подставляется "/".
func CallerRootPath(value string) CallerOption {
	return func(c *Caller) {
		if value != "" {
			value = strings.TrimRight(filepath.ToSlash(value), "/") + "/"
		}

		c.rootPath = value
	}
}
