package gotext

import (
	"errors"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type (
	// Provider - локализатор сообщений построенный на библиотеке gotext.
	Provider struct {
		domain2langCatalog map[string]map[language.Tag]*message.Printer
		defaultPrinter     *message.Printer
		domains            []string
		languages          []language.Tag
	}

	options struct {
		domains   []string
		catalogs  []catalog.Catalog
		languages []language.Tag
	}
)

// NewProvider - создаёт объект Provider.
func NewProvider(opts ...Option) (*Provider, error) {
	o := options{}

	for _, opt := range opts {
		opt(&o)
	}

	if len(o.languages) == 0 {
		return nil, errors.New("gotext provider:opts.Languages is required")
	}

	if len(o.catalogs) == 0 {
		return nil, errors.New("gotext provider:opts.Catalogs is required")
	}

	p := &Provider{
		domain2langCatalog: make(map[string]map[language.Tag]*message.Printer, len(o.domains)),
		defaultPrinter:     message.NewPrinter(o.languages[0]),
	}

	for i, domain := range o.domains {
		p.domain2langCatalog[domain] = make(map[language.Tag]*message.Printer, len(o.languages))

		for _, lang := range o.languages {
			p.domain2langCatalog[domain][lang] = message.NewPrinter(lang, message.Catalog(o.catalogs[i]))
		}
	}

	// uniq domains
	for domain := range p.domain2langCatalog {
		p.domains = append(p.domains, domain)
	}

	// uniq languages
	for lang := range p.domain2langCatalog[o.domains[0]] {
		p.languages = append(p.languages, lang)
	}

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
