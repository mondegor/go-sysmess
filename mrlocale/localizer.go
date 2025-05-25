package mrlocale

import (
	"golang.org/x/text/language"
)

type (
	// Localizer - загружает и локализует сообщения, ошибки, справочники объектов на различные языки.
	Localizer struct {
		bundle   *Bundle
		language language.Tag
	}
)

// Language - возвращает язык локализации сообщений.
func (l *Localizer) Language() string {
	return l.language.String()
}

// Translate - локализует указанное сообщение подставляя в него указанные аргументы.
func (l *Localizer) Translate(msg string, args ...any) string {
	return l.bundle.localize(l.bundle.messagesDomain, l.language, msg, args)
}

// TranslateError - локализует указанное сообщение об ошибке.
func (l *Localizer) TranslateError(err error) string {
	msg, args := l.bundle.formatError(err)

	return l.bundle.localize(l.bundle.errorsDomain, l.language, msg, args)
}

// TranslateInDomain - локализует в указанном домене указанное сообщение подставляя в него указанные аргументы.
func (l *Localizer) TranslateInDomain(domain, msg string, args ...any) string {
	return l.bundle.localize(domain, l.language, msg, args)
}
