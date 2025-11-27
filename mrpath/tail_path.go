package mrpath

import (
	"strings"
)

type (
	tailBuilder struct {
		basePath string
	}
)

// NewTail - создаёт объект Builder.
// sample: /base/dir/ -> /base/dir/path
func NewTail(basePath string) Builder {
	return &tailBuilder{
		basePath: strings.TrimRight(basePath, "/") + "/",
	}
}

// BuildPath - возвращает полный путь добавляя в хвост базового указанный путь.
func (p *tailBuilder) BuildPath(path string) string {
	if path == "" {
		return ""
	}

	return p.basePath + strings.TrimLeft(path, "/")
}
