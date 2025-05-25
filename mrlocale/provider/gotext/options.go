package gotext

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

type (
	// Option - настройка объекта Provider.
	Option func(o *options)
)

// WithCatalog - добавляет справочник фраз привязывая его к указанному домену.
func WithCatalog(domain string, cat catalog.Catalog) Option {
	return func(o *options) {
		if cat == nil {
			return
		}

		o.domains = append(o.domains, domain)
		o.catalogs = append(o.catalogs, cat)
	}
}

// WithLanguages - добавляет языки для локализации сообщений.
func WithLanguages(values ...language.Tag) Option {
	return func(o *options) {
		if len(values) > 0 {
			o.languages = append(o.languages, values...)
		}
	}
}
