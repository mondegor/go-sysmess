package mrlang

import "fmt"

const (
	langPathPattern = "%s/%s.%s" // dir/lang.ext, ./translate/en.yaml
)

type (
	langMap map[string]*Locale

	Translator struct {
		langs         langMap
		defaultLocale *Locale
	}

	TranslatorOptions struct {
		DirPath     string
		FileType    string
		LangCodes   []string
		DefaultLang string // optional
	}
)

func NewTranslator(opt TranslatorOptions) (*Translator, error) {
	if len(opt.LangCodes) == 0 {
		return nil, fmt.Errorf("opt.LangCodes is required")
	}

	if opt.DefaultLang == "" {
		opt.DefaultLang = opt.LangCodes[0]
	} else if !defaultLangInArray(opt.DefaultLang, opt.LangCodes) {
		return nil, fmt.Errorf("opt.DefaultLang='%s' not found in opt.LangCodes", opt.DefaultLang)
	}

	tr := Translator{
		langs: make(langMap, 0),
	}

	for _, langCode := range opt.LangCodes {
		locale, err := newLocale(langCode, fmt.Sprintf(langPathPattern, opt.DirPath, langCode, opt.FileType))

		if err != nil {
			return nil, err
		}

		tr.langs[langCode] = locale

		if opt.DefaultLang == langCode {
			tr.defaultLocale = locale
		}
	}

	return &tr, nil
}

func defaultLangInArray(lang string, langs []string) bool {
	for _, value := range langs {
		if value == lang {
			return true
		}
	}

	return false
}

func (t *Translator) DefaultLocale() *Locale {
	return t.defaultLocale
}

func (t *Translator) FindFirstLocale(langs ...string) *Locale {
	for _, lang := range langs {
		if locale, ok := t.langs[lang]; ok {
			return locale
		}
	}

	return t.defaultLocale
}

func (t *Translator) RegisteredLocales() []*Locale {
	locs := make([]*Locale, len(t.langs))
	i := 0

	for _, locale := range t.langs {
		locs[i] = locale
		i++
	}

	return locs
}
