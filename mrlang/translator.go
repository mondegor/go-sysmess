package mrlang

import "fmt"

const (
	langPathPattern = "%s/%s.%s" // dir/lang.ext, ./translate/en.yaml
)

type (
    langMap map[string]*Locale

    Translator struct {
        langs langMap
        defaultLocale *Locale
    }

    TranslatorOptions struct {
        DirPath string
        FileType string
        LangCodes []string
        DefaultLang string // optional
    }
)

func NewTranslator(opt TranslatorOptions) (*Translator, error) {
    if len(opt.LangCodes) == 0 {
        return nil, fmt.Errorf("opt.LangCodes is required")
    }

    langCodes := opt.LangCodes

    if opt.DefaultLang == "" {
        opt.DefaultLang = langCodes[0]
    } else if !defaultLangInArray(opt.DefaultLang, &langCodes) {
        return nil, fmt.Errorf("opt.DefaultLang='%s' is not found in opt.LangCodes", opt.DefaultLang)
    }

    tr := Translator{
        langs: make(langMap, 2),
    }

    for _, langCode := range opt.LangCodes {
        loc, err := newLocale(langCode, fmt.Sprintf(langPathPattern, opt.DirPath, langCode, opt.FileType))

        if err != nil {
            return nil, err
        }

        tr.langs[langCode] = loc

        if opt.DefaultLang == langCode {
            tr.defaultLocale = loc
        }
    }

    return &tr, nil
}

func defaultLangInArray(lang string, langs *[]string) bool {
    for _, value := range *langs {
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
        if loc, ok := t.langs[lang]; ok {
            return loc
        }
    }

    return t.defaultLocale
}

func (t *Translator) RegisteredLocales() []*Locale {
    var locs []*Locale

    for _, loc := range t.langs {
        locs = append(locs, loc)
    }

    return locs
}
