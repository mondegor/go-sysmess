package mrlocale

import (
	"golang.org/x/text/language"
)

type (
	// Pool - управляет набором Localizer для всех поддерживаемых языков.
	// Создаётся из Bundle и предоставляет удобный поиск локализатора
	// по близости к указанному языку через language.Matcher.
	Pool struct {
		bundle           *Bundle
		localizers       map[language.Tag]*Localizer
		defaultLocalizer *Localizer
	}
)

// NewPool - создаёт Pool из Bundle.
// Создаёт Localizer для каждого языка, поддерживаемого провайдером.
func NewPool(bundle *Bundle) *Pool {
	languages := bundle.provider.Languages()
	localizers := make(map[language.Tag]*Localizer, len(languages))

	for _, lang := range languages {
		localizers[lang] = &Localizer{
			bundle:   bundle,
			language: lang,
		}
	}

	return &Pool{
		bundle:           bundle,
		localizers:       localizers,
		defaultLocalizer: localizers[bundle.defaultLanguage],
	}
}

// Localizer - возвращает Localizer для наиболее близкого языка из списка langs.
// Использует language.Matcher для поиска лучшего совпадения.
// Если совпадение не найдено, возвращает Localizer для языка по умолчанию.
// Если langs пуст, также возвращает Localizer по умолчанию.
func (p *Pool) Localizer(langs ...language.Tag) *Localizer {
	if len(langs) > 0 {
		lang, _, _ := p.bundle.languageMatcher.Match(langs...)

		if l, ok := p.localizers[lang]; ok {
			return l
		}
	}

	return p.defaultLocalizer
}
