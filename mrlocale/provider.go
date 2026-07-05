package mrlocale

import (
	"golang.org/x/text/language"
)

type (
	// MessageProvider - интерфейс провайдера локализации сообщений.
	// Реализация должна хранить переводы для набора языков и доменов.
	MessageProvider interface {
		Domains() []string
		Languages() []language.Tag
		Localize(domain string, lang language.Tag, msg string, args []any) string
	}
)
