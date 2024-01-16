package mrlang

import (
	"regexp"
	"strings"
)

const (
	maxAcceptLanguageLen = 256
)

var (
	regexpAcceptLanguage = regexp.MustCompile(`^[a-z]{2}(-[a-zA-Z0-9-]+)?$`)
)

// ParseAcceptLanguage
// Sample Accept-Language: ru;q=0.9, fr-CH, fr;q=0.8, en;q=0.7, *;q=0.5
func ParseAcceptLanguage(s string) []string {
	length := len(s)

	if length > 0 && length <= maxAcceptLanguageLen {
		var langs []string
		var keys map[string]bool

		addLang := func(lang string) bool {
			if keys == nil {
				keys = make(map[string]bool, 1)
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
			} else {
				if addLang(lang) {
					addLang(lang + "_" + strings.ToUpper(lang))
				}
			}
		}

		if len(langs) > 0 {
			return langs
		}
	}

	return []string{}
}
