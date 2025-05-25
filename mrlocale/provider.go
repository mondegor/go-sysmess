package mrlocale

import (
	"golang.org/x/text/language"
)

type (
	// MessageProvider - интерфейс для создания провайдеров локализующих сообщения.
	MessageProvider interface {
		Domains() []string
		Languages() []language.Tag
		Localize(domain string, lang language.Tag, msg string, args []any) string
	}
)
