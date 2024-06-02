package mrlang

import (
	"fmt"
)

type (
	// Translator - загружает и транслирует сообщения, ошибки,
	// справочники объектов на различные языки.
	Translator struct {
		locByIDs      locByIDsMap
		locByCodes    locByCodesMap
		defaultLocale *Locale
		dictionaries  map[string]*MultiLangDictionary
	}

	// TranslatorOptions - опции для создания Translator.
	TranslatorOptions struct {
		DirPath           string
		LangCodes         []string
		DefaultLang       string // optional
		DictionaryDirPath string
		Dictionaries      []string // optional
	}

	locByIDsMap   map[uint16]*Locale
	locByCodesMap map[string]*Locale
)

// NewTranslator - создаётся объект Translator.
func NewTranslator(opts TranslatorOptions) (*Translator, error) {
	if len(opts.LangCodes) == 0 {
		return nil, fmt.Errorf("opts.LangCodes is required")
	}

	if opts.DefaultLang == "" {
		opts.DefaultLang = opts.LangCodes[0]
	} else if !defaultLangInArray(opts.DefaultLang, opts.LangCodes) {
		return nil, fmt.Errorf("opts.DefaultLang='%s' not found in opts.LangCodes", opts.DefaultLang)
	}

	tr := Translator{
		locByIDs:     make(locByIDsMap, 0),
		locByCodes:   make(locByCodesMap, 0),
		dictionaries: make(map[string]*MultiLangDictionary, 0),
	}

	for _, langCode := range opts.LangCodes {
		if _, ok := tr.locByCodes[langCode]; ok {
			return nil, fmt.Errorf("duplicate locale '%s' detected", langCode)
		}

		locale, err := newLocale(langCode, getFilePath(opts.DirPath, langCode))
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

		if opts.DefaultLang == langCode {
			tr.defaultLocale = locale
		}
	}

	for _, dicName := range opts.Dictionaries {
		if _, ok := tr.dictionaries[dicName]; ok {
			return nil, fmt.Errorf("duplicate dictionary '%s' detected", dicName)
		}

		multiDict := MultiLangDictionary{
			name:           dicName,
			dicByLangIDs:   make(dicByLangIDsMap, 0),
			dicByLangCodes: make(dicByLangCodesMap, 0),
		}

		for _, langCode := range opts.LangCodes {
			dict, err := newDictionary(
				getFilePath(opts.DictionaryDirPath, dicName+"/"+langCode),
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

// DefaultLocale - возвращает языковой объект по умолчанию.
func (t *Translator) DefaultLocale() *Locale {
	return t.defaultLocale
}

// FindFirstLocale - возвращает первый языковой объект по указанным языковым кодам.
// Если ни один языковой код не зарегистрирован, то возвращает языковой объект по умолчанию.
func (t *Translator) FindFirstLocale(langs ...string) *Locale {
	for _, lang := range langs {
		if locale, ok := t.locByCodes[lang]; ok {
			return locale
		}
	}

	return t.defaultLocale
}

// LocaleByID - возвращает языковой объект по его ID.
func (t *Translator) LocaleByID(langID uint16) (*Locale, error) {
	if locale, ok := t.locByIDs[langID]; ok {
		return locale, nil
	}

	return nil, fmt.Errorf("lang with ID=%d is not registered", langID)
}

// LocaleByCode - возвращает языковой объект по его коду.
func (t *Translator) LocaleByCode(lang string) (*Locale, error) {
	if locale, ok := t.locByCodes[lang]; ok {
		return locale, nil
	}

	return nil, fmt.Errorf("lang code '%s' is not registered", lang)
}

// Dictionary - возвращает мультиязычный справочник по его имени.
func (t *Translator) Dictionary(name string) (*MultiLangDictionary, error) {
	if dictionary, ok := t.dictionaries[name]; ok {
		return dictionary, nil
	}

	return nil, fmt.Errorf("dictionary '%s' is not registered", name)
}

// RegisteredLocales - возвращает список кодов зарегистрированных языков.
func (t *Translator) RegisteredLocales() []string {
	keys := make([]string, 0, len(t.locByCodes))

	for key := range t.locByCodes {
		keys = append(keys, key)
	}

	return keys
}

// RegisteredDictionaries - возвращает список имён зарегистрированных мультиязычных справочников.
func (t *Translator) RegisteredDictionaries() []string {
	keys := make([]string, 0, len(t.dictionaries))

	for key := range t.dictionaries {
		keys = append(keys, key)
	}

	return keys
}
