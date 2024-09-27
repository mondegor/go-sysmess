package mrlang

import (
	"regexp"
	"strings"
)

const (
	maxAcceptLanguageLen = 256
)

var regexpAcceptLanguage = regexp.MustCompile(`^[a-z]{2}(-[a-zA-Z0-9-]+)?$`)

// ParseAcceptLanguage - извлечение данных о языках.
// Sample Accept-Language: ru;q=0.9, fr-CH, fr;q=0.8, en;q=0.7, *;q=0.5.
func ParseAcceptLanguage(s string) []string {
	if len(s) == 0 || len(s) > maxAcceptLanguageLen {
		return nil
	}

	var (
		langs []string
		keys  map[string]bool
	)

	addLang := func(lang string) bool {
		if keys == nil {
			keys = make(map[string]bool)
		} else {
			if _, ok := keys[lang]; ok {
				return false
			}
		}

		langs = append(langs, lang)
		keys[lang] = true

		return true
	}

	for _, lang := range strings.Split(s, ",") {
		if index := strings.Index(lang, ";"); index >= 0 {
			lang = lang[:index]
		}

		lang = strings.TrimSpace(lang)

		if !regexpAcceptLanguage.MatchString(lang) {
			continue
		}

		if len(lang) > 2 {
			addLang(lang[0:2] + "_" + lang[3:]) // ru-RU -> ru_RU
		} else if addLang(lang) { // ru + ru_RU
			addLang(lang + "_" + strings.ToUpper(lang))
		}
	}

	if len(langs) > 0 {
		return langs
	}

	return nil
}
