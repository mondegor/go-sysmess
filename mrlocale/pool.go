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
		localizers       []*Localizer
		defaultLocalizer *Localizer
	}
)

// NewPool - создаёт Pool из Bundle.
// Создаёт Localizer для каждого языка, поддерживаемого бандлом.
//
// Локализаторы хранятся срезом в том же порядке, что и Bundle.languages,
// чтобы индекс, возвращаемый language.Matcher.Match, адресовал их напрямую:
// матчер построен на том же срезе, поэтому индекс и позиция локализатора
// заведомо совпадают, и промежуточный поиск по тегу не нужен.
func NewPool(bundle *Bundle) *Pool {
	localizers := make([]*Localizer, len(bundle.languages))
	defaultLocalizer := &Localizer{
		bundle:   bundle,
		language: bundle.defaultLanguage,
	}

	for i, lang := range bundle.languages {
		if lang == bundle.defaultLanguage {
			localizers[i] = defaultLocalizer

			continue
		}

		localizers[i] = &Localizer{
			bundle:   bundle,
			language: lang,
		}
	}

	return &Pool{
		bundle:           bundle,
		localizers:       localizers,
		defaultLocalizer: defaultLocalizer,
	}
}

// Localizer - возвращает Localizer для наиболее близкого языка из списка langs.
// Использует language.Matcher для поиска лучшего совпадения.
// Если совпадение не найдено, возвращает Localizer для языка по умолчанию.
// Если langs пуст, также возвращает Localizer по умолчанию.
//
// Язык выбирается по индексу, который возвращает Match, а не по тегу: для тегов
// с регионом Match возвращает тег с расширением (например, "en-US" -> "en-u-rg-uszzzz"),
// поэтому поиск по тегу промахивался бы мимо любого языка списка.
//
// Индекс на длину локализаторов не проверяется: матчер построен на Bundle.languages,
// из которого NewPool создаёт локализаторы, поэтому индекс заведомо в границах среза.
func (p *Pool) Localizer(langs ...language.Tag) *Localizer {
	if len(langs) > 0 {
		_, index, conf := p.bundle.languageMatcher.Match(langs...)

		// при conf == language.No матчер возвращает index == 0, то есть первый язык
		// списка, а он не обязан быть языком по умолчанию, поэтому промах отсекается
		// отдельно, а не разбирается вместе с удачными подборами
		if conf != language.No {
			return p.localizers[index]
		}
	}

	return p.defaultLocalizer
}
