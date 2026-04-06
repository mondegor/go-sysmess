package mrpath

import (
	"strings"
)

type (
	// tailBuilder - построитель путей путём добавления переданного значения в конец базового пути.
	tailBuilder struct {
		basePath string
	}
)

// NewTail - создаёт Builder, который добавляет переданный путь в конец basePath.
// Параметр basePath - базовый путь, к которому будет добавляться значение (например, "/base/dir/").
// Нормализует trailing slash: убирает все завершающие "/" и добавляет одну.
// Пример: basePath="/base/dir/" → "/base/dir/path".
func NewTail(basePath string) Builder {
	return &tailBuilder{
		basePath: strings.TrimRight(basePath, "/") + "/",
	}
}

// BuildPath - добавляет указанный путь в конец базового пути.
// Убирает ведущий "/" из path для корректной склейки.
// Возвращает пустую строку, если path пуст.
func (p *tailBuilder) BuildPath(path string) string {
	if path == "" {
		return ""
	}

	return p.basePath + strings.TrimLeft(path, "/")
}
