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
	// Bundle - управляет локализацией сообщений, ошибок и справочников.
	// Загружает переводы из провайдера, выбирает подходящий язык
	// и форматирует сообщения с учётом переданных аргументов.
	Bundle struct {
		provider        MessageProvider
		messagesDomain  string
		errorsDomain    string
		formatMessage   func(msg string, args []any) (newMsg string, newArgs []any)
		formatError     func(err error) (msg string, args []any)
		languageMatcher language.Matcher
		defaultLanguage language.Tag
	}
)

// NewBundle - создаёт Bundle для локализации сообщений.
// Параметры:
//   - languages - список поддерживаемых языков в формате BCP 47 (например: "en", "ru", "en-US");
//   - opts - дополнительные опции настройки (WithMessageProvider, WithDefaultLanguage и др.);
//
// Первый язык в списке становится языком по умолчанию, если не задан явно через WithDefaultLanguage.
func NewBundle(languages []string, opts ...BundleOption) (*Bundle, error) {
	o := bundleOptions{
		messagesDomain: DefaultMessagesDomain,
		errorsDomain:   DefaultErrorsDomain,
	}

	for _, opt := range opts {
		opt(&o)
	}

	if len(languages) == 0 {
		return nil, errors.New("bundle create: no matching language found")
	}

	languageTags := make([]language.Tag, len(languages))
	defaultLanguage := language.Und

	for i, lang := range languages {
		tag, err := language.Parse(lang)
		if err != nil {
			return nil, fmt.Errorf("bundle create: parsing language '%s': %w", lang, err)
		}

		languageTags[i] = tag

		if o.defaultLanguage == lang {
			defaultLanguage = tag
		}
	}

	if defaultLanguage == language.Und {
		if o.defaultLanguage != "" {
			return nil, fmt.Errorf("bundle create: default language is not supported (lang='%s')", o.defaultLanguage)
		}

		// если в опции явно не указан язык по умолчанию,
		// то по умолчанию используется первый язык в списке
		defaultLanguage = languageTags[0]
	}

	pr, err := o.createProvider(languageTags)
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
		languageMatcher: language.NewMatcher(languageTags),
		defaultLanguage: defaultLanguage,
	}, nil
}

// localize - выполняет локализацию сообщения в указанном домене.
// Предварительно обрабатывает сообщение и аргументы через formatMessage.
func (b *Bundle) localize(domain string, lang language.Tag, msg string, args []any) string {
	msg, args = b.formatMessage(msg, args)

	return b.provider.Localize(domain, lang, msg, args)
}

// isDomainInArray - проверяет наличие домена в массиве доменов.
func isDomainInArray(domain string, domains []string) bool {
	for _, val := range domains {
		if val == domain {
			return true
		}
	}

	return false
}
