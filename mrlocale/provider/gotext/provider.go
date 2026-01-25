package gotext

import (
	"errors"
	"maps"
	"slices"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type (
	// Provider - локализатор сообщений построенный на библиотеке gotext.
	Provider struct {
		defaultPrinter     *message.Printer
		domains            []string
		languages          []language.Tag
		domain2langCatalog map[string]map[language.Tag]*message.Printer
	}
)

// NewProvider - создаёт объект Provider.
func NewProvider(
	languages []language.Tag,
	opts ...Option,
) (*Provider, error) {
	o := options{}

	for _, opt := range opts {
		opt(&o)
	}

	if len(languages) == 0 {
		return nil, errors.New("gotext provider:languages is required")
	}

	p := &Provider{
		defaultPrinter: message.NewPrinter(languages[0]),
	}

	if len(o.domains) == 0 {
		return p, nil
	}

	p.domain2langCatalog = make(map[string]map[language.Tag]*message.Printer, len(o.domains))

	for i, domain := range o.domains {
		p.domain2langCatalog[domain] = make(map[language.Tag]*message.Printer, len(languages))

		for _, lang := range languages {
			p.domain2langCatalog[domain][lang] = message.NewPrinter(lang, message.Catalog(o.catalogs[i]))
		}
	}

	// uniq domains
	p.domains = slices.Collect(maps.Keys(p.domain2langCatalog))

	// uniq languages
	p.languages = slices.Collect(maps.Keys(p.domain2langCatalog[o.domains[0]]))

	return p, nil
}

// Domains - возвращает список используемых доменов.
func (p *Provider) Domains() []string {
	return p.domains
}

// Languages - возвращает список используемых языков.
func (p *Provider) Languages() []language.Tag {
	return p.languages
}

// Localize - возвращает локализованное сообщение с подставленными аргументами.
func (p *Provider) Localize(domain string, lang language.Tag, msg string, args []any) string {
	printer := p.defaultPrinter

	if pr, ok := p.domain2langCatalog[domain][lang]; ok {
		printer = pr
	}

	return printer.Sprintf(msg, args...)
}
