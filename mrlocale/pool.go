package mrlocale

import (
	"golang.org/x/text/language"
)

type (
	// Pool - загружает и транслирует сообщения, ошибки,
	// справочники объектов на различные языки.
	Pool struct {
		bundle           *Bundle
		localizers       map[language.Tag]*Localizer
		defaultLocalizer *Localizer
	}
)

// NewPool - создаёт объект Pool.
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

// Localizer - возвращает локализатор сообщений с использованием самого близкого языка для перевода из указанных.
// Если такой язык найти не удалось, то возвращается локализатор с языком по умолчанию.
func (p *Pool) Localizer(langs ...language.Tag) *Localizer {
	if len(langs) > 0 {
		lang, _, _ := p.bundle.languageMatcher.Match(langs...)

		if l, ok := p.localizers[lang]; ok {
			return l
		}
	}

	return p.defaultLocalizer
}
