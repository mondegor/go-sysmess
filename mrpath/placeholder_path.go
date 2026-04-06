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
	// placeholderBuilder - построитель путей с плейсхолдером в середине пути.
	// Разбивает базовый путь на префикс и постфикс относительно плейсхолдера,
	// затем вставляет переданное значение между ними.
	placeholderBuilder struct {
		basePath string
		postfix  string
	}
)

// NewPlaceholder - создаёт Builder, который заменяет плейсхолдер в basePath на переданный путь.
// Параметры:
//   - basePath - шаблон пути с плейсхолдером (например, "/dir/{{path}}/postfix");
//   - placeholder - строка-плейсхолдер для замены (например, "{{path}}").
//
// Возвращает ошибку, если плейсхолдер пуст или не найден в basePath.
// Пример: basePath="/dir/{{path}}/postfix", placeholder="{{path}}" → "/dir/real-value/postfix".
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

// BuildPath - вставляет указанный путь на место плейсхолдера в базовом шаблоне.
// Убирает ведущие и завершающие "/" из path для корректной склейки.
// Возвращает пустую строку, если path пуст.
func (p *placeholderBuilder) BuildPath(path string) string {
	if path == "" {
		return ""
	}

	return p.basePath + strings.Trim(path, "/") + p.postfix
}
