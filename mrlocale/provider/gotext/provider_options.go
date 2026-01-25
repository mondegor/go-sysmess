package gotext

import (
	"golang.org/x/text/message/catalog"
)

type (
	// Option - настройка объекта Provider.
	Option func(o *options)

	options struct {
		domains  []string
		catalogs []catalog.Catalog
	}
)

// WithDomainCatalog - добавляет справочник фраз привязывая его к указанному домену.
func WithDomainCatalog(domain string, cat catalog.Catalog) Option {
	return func(o *options) {
		if cat == nil {
			return
		}

		o.domains = append(o.domains, domain)
		o.catalogs = append(o.catalogs, cat)
	}
}
