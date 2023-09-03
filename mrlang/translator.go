package mrlang

import "fmt"

const langPathPattern = "%s/%s.%s" // dir/lang.ext, ./translate/en.yaml

type (
    langMap map[string]*Locale

    Translator struct {
        langs langMap
        defaultLocale *Locale
    }

    TranslatorOptions struct {
        LangByDefault string
        DirPath string
        FileType string
        LangCodes []string
    }
)

func NewTranslator(opt TranslatorOptions) (*Translator, error) {
    if opt.LangByDefault == "" {
        return nil, fmt.Errorf("opt.LangByDefault is required")
    }

    var langCodes []string

    if len(opt.LangCodes) > 0 {
        langCodes = opt.LangCodes
    } else {
        langCodes = append(langCodes, opt.LangByDefault)
    }

    tr := Translator{
        langs: make(langMap, 2),
    }

    for i, langCode := range opt.LangCodes {
        loc, err := newLocale(langCode, fmt.Sprintf(langPathPattern, opt.DirPath, langCode, opt.FileType))

        if err != nil {
            return nil, err
        }

        tr.langs[langCode] = loc

        if i == 0 {
            tr.defaultLocale = loc
        }
    }

    return &tr, nil
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
