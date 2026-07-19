package mrlocale

import (
	"golang.org/x/text/language"
)

type (
	// MessageProvider - интерфейс провайдера локализации сообщений.
	// Реализация должна хранить переводы для набора языков и доменов.
	//
	// Список языков провайдер не публикует: языки задаются при создании Bundle
	// и хранятся им, поэтому спрашивать их у провайдера незачем. Localize
	// вызывается только с языком из этого списка, а неизвестный язык
	// реализация обязана обслужить языком по умолчанию, а не отказом.
	MessageProvider interface {
		Domains() []string
		Localize(domain string, lang language.Tag, msg string, args []any) string
	}
)
