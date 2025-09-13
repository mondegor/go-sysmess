package mrlocale

import (
	"errors"
	"fmt"

	"golang.org/x/text/language"
)

// Предопределённые названия доменов по умолчанию.
const (
	DefaultMessagesDomain = "messages"
	DefaultErrorsDomain   = "errors"
)

type (
	// Bundle - загружает и транслирует сообщения, ошибки,
	// справочники объектов на различные языки.
	Bundle struct {
		provider        MessageProvider
		messagesDomain  string
		errorsDomain    string
		formatMessage   func(msg string, args []any) (newMsg string, newArgs []any)
		formatError     func(err error) (msg string, args []any)
		languageMatcher language.Matcher
		defaultLanguage language.Tag
	}

	bundleOptions struct {
		createProvider  func(langs []language.Tag) (MessageProvider, error)
		formatMessage   func(msg string, args []any) (newMsg string, newArgs []any)
		formatError     func(err error) (msg string, args []any)
		messagesDomain  string
		errorsDomain    string
		languages       []string
		defaultLanguage string
	}
)

// NewBundle - создаёт объект Bundle.
func NewBundle(opts ...BundleOption) (*Bundle, error) {
	o := bundleOptions{
		messagesDomain: DefaultMessagesDomain,
		errorsDomain:   DefaultErrorsDomain,
	}

	for _, opt := range opts {
		opt(&o)
	}

	if len(o.languages) == 0 {
		return nil, errors.New("bundle create: no matching language found")
	}

	languages := make([]language.Tag, len(o.languages))
	defaultLanguage := language.Und

	for i, lang := range o.languages {
		tag, err := language.Parse(lang)
		if err != nil {
			return nil, fmt.Errorf("bundle create: parsing language '%s': %w", lang, err)
		}

		languages[i] = tag

		if o.defaultLanguage == lang {
			defaultLanguage = tag
		}
	}

	if defaultLanguage == language.Und {
		if o.defaultLanguage != "" {
			return nil, fmt.Errorf("bundle create: default language '%s' is not supported", o.defaultLanguage)
		}

		// если в опции явно не указан язык по умолчанию,
		// то по умолчанию используется первый язык в списке
		defaultLanguage = languages[0]
	}

	pr, err := o.createProvider(languages)
	if err != nil {
		return nil, fmt.Errorf("bundle create: %w", err)
	}

	if !isDomainInArray(o.messagesDomain, pr.Domains()) {
		return nil, fmt.Errorf("bundle create: provider has not messages domain '%s'", o.messagesDomain)
	}

	if !isDomainInArray(o.errorsDomain, pr.Domains()) {
		return nil, fmt.Errorf("bundle create: provider has not errors domain '%s'", o.errorsDomain)
	}

	if o.formatMessage == nil {
		o.formatMessage = func(msg string, args []any) (newMsg string, newArgs []any) {
			return msg, args
		}
	}

	if o.formatError == nil {
		o.formatError = func(err error) (string, []any) {
			return err.Error(), nil
		}
	}

	return &Bundle{
		provider:        pr,
		messagesDomain:  o.messagesDomain,
		errorsDomain:    o.errorsDomain,
		formatMessage:   o.formatMessage,
		formatError:     o.formatError,
		languageMatcher: language.NewMatcher(languages),
		defaultLanguage: defaultLanguage,
	}, nil
}

func (b *Bundle) localize(domain string, lang language.Tag, msg string, args []any) string {
	msg, args = b.formatMessage(msg, args)

	return b.provider.Localize(domain, lang, msg, args)
}

func isDomainInArray(domain string, domains []string) bool {
	for _, val := range domains {
		if val == domain {
			return true
		}
	}

	return false
}
