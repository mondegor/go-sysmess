package mrlang

import (
    "fmt"

    "github.com/mondegor/go-sysmess/mrmsg"
)

type (
    Translator interface {
        LocaleFirstFound(langs ...string) Locale
        RegisteredLocales() []Locale
    }

    Locale interface {
        LangCode() string
        TranslateMessage(id string, defaultMessage string, args ...mrmsg.NamedArg) Message
        TranslateError(id string, defaultMessage string, args ...mrmsg.NamedArg) ErrorMessage
    }

    langMap map[string]*locale

    translator struct {
        langs langMap
        defaultLocale *locale
    }

    TranslatorOptions struct {
        LangByDefault string
        DirPath string
        FileType string
        LangCodes []string
    }
)

func NewTranslator(opt TranslatorOptions) (Translator, error) {
    if opt.LangByDefault == "" {
        return nil, fmt.Errorf("opt.LangByDefault is required")
    }

    var langCodes []string

    if len(opt.LangCodes) > 0 {
        langCodes = opt.LangCodes
    } else {
        langCodes = append(langCodes, opt.LangByDefault)
    }

    tr := translator{
        langs: make(langMap, 2),
    }

    for i, langCode := range opt.LangCodes {
        loc, err := newLocale(langCode, fmt.Sprintf("%s/%s.%s", opt.DirPath, langCode, opt.FileType))

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

func (t *translator) LocaleFirstFound(langs ...string) Locale {
    for _, lang := range langs {
        if loc, ok := t.langs[lang]; ok {
            return loc
        }
    }

    return t.defaultLocale
}

func (t *translator) RegisteredLocales() []Locale {
    var locs []Locale

    for _, loc := range t.langs {
        locs = append(locs, loc)
    }

    return locs
}
