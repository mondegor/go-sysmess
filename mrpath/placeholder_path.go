package mrpath

import (
	"fmt"
	"strings"
)

const (
	// Placeholder - плейсхолдер в пути по умолчанию.
	Placeholder = "{{path}}"
)

type (
	placeholderBuilder struct {
		basePath string
		postfix  string
	}
)

// NewPlaceholder - создаёт объект Builder.
// sample: /dir/{{path}}/postfix -> /dir/real-value/postfix
func NewPlaceholder(basePath, placeholder string) (Builder, error) {
	if placeholder == "" {
		return nil, fmt.Errorf("placeholder is empty for path '%s'", basePath)
	}

	if i := strings.Index(basePath, placeholder); i > 0 {
		return &placeholderBuilder{
			basePath: basePath[0:i],
			postfix:  basePath[i+len(placeholder):],
		}, nil
	}

	return nil, fmt.Errorf("placeholder is not found in path (placeholder='%s', path='%s')", placeholder, basePath)
}

// BuildPath - возвращает полный путь вставляя в базовый указанный путь.
func (p *placeholderBuilder) BuildPath(path string) string {
	if path == "" {
		return ""
	}

	return p.basePath + strings.Trim(path, "/") + p.postfix
}
