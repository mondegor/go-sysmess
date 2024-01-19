package mrlang

import (
	"fmt"
)

type (
	locByIDsMap   map[uint16]*Locale
	locByCodesMap map[string]*Locale

	Translator struct {
		locByIDs      locByIDsMap
		locByCodes    locByCodesMap
		defaultLocale *Locale
		dictionaries  map[string]*MultiLangDictionary
	}

	TranslatorOptions struct {
		DirPath           string
		LangCodes         []string
		DefaultLang       string // optional
		DictionaryDirPath string
		Dictionaries      []string // optional
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
		locByIDs:     make(locByIDsMap, 0),
		locByCodes:   make(locByCodesMap, 0),
		dictionaries: make(map[string]*MultiLangDictionary, 0),
	}

	for _, langCode := range opt.LangCodes {
		if _, ok := tr.locByCodes[langCode]; ok {
			return nil, fmt.Errorf("duplicate locale '%s' detected", langCode)
		}

		locale, err := newLocale(langCode, filePath(opt.DirPath, langCode))

		if err != nil {
			return nil, err
		}

		if test, ok := tr.locByIDs[locale.LangID()]; ok {
			return nil, fmt.Errorf(
				"lang code '%s' with ID=%d is already registered in locale '%s'",
				langCode,
				locale.LangID(),
				test.LangCode(),
			)
		}

		tr.locByIDs[locale.LangID()] = locale
		tr.locByCodes[langCode] = locale

		if opt.DefaultLang == langCode {
			tr.defaultLocale = locale
		}
	}

	for _, dicName := range opt.Dictionaries {
		if _, ok := tr.dictionaries[dicName]; ok {
			return nil, fmt.Errorf("duplicate dictionary '%s' detected", dicName)
		}

		multiDict := MultiLangDictionary{
			name:           dicName,
			dicByLangIDs:   make(dicByLangIDsMap, 0),
			dicByLangCodes: make(dicByLangCodesMap, 0),
		}

		for _, langCode := range opt.LangCodes {
			dict, err := newDictionary(
				filePath(opt.DictionaryDirPath, dicName+"/"+langCode),
			)

			if err != nil {
				return nil, err
			}

			multiDict.dicByLangIDs[tr.locByCodes[langCode].LangID()] = dict
			multiDict.dicByLangCodes[langCode] = dict
		}

		tr.dictionaries[dicName] = &multiDict
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

// FindFirstLocale - if not exists get default locale
func (t *Translator) FindFirstLocale(langs ...string) *Locale {
	for _, lang := range langs {
		if locale, ok := t.locByCodes[lang]; ok {
			return locale
		}
	}

	return t.defaultLocale
}

func (t *Translator) LocaleByID(langID uint16) (*Locale, error) {
	if locale, ok := t.locByIDs[langID]; ok {
		return locale, nil
	}

	return nil, fmt.Errorf("lang with ID=%d is not registered", langID)
}

func (t *Translator) LocaleByCode(lang string) (*Locale, error) {
	if locale, ok := t.locByCodes[lang]; ok {
		return locale, nil
	}

	return nil, fmt.Errorf("lang code '%s' is not registered", lang)
}

func (t *Translator) Dictionary(name string) (*MultiLangDictionary, error) {
	if dictionary, ok := t.dictionaries[name]; ok {
		return dictionary, nil
	}

	return nil, fmt.Errorf("dictionary '%s' is not registered", name)
}

func (t *Translator) RegisteredLocales() []string {
	keys := make([]string, len(t.locByCodes))
	i := 0

	for key := range t.locByCodes {
		keys[i] = key
		i++
	}

	return keys
}

func (t *Translator) RegisteredDictionaries() []string {
	keys := make([]string, len(t.dictionaries))
	i := 0

	for key := range t.dictionaries {
		keys[i] = key
		i++
	}

	return keys
}
