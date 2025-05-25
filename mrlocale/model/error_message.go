package model

import (
	"strings"
)

type (
	// ErrorMessage - модель ошибки используемая при её локализации.
	ErrorMessage struct {
		Reason  string
		Details string
	}
)

// String - возвращает причину ошибки и её подробностей в виде строки.
func (m ErrorMessage) String() string {
	if m.Details == "" {
		return m.Reason
	}

	return m.Reason + "\n\n" + m.Details
}

// ParseErrorMessage - парсинг строки хранящейся в справочнике ошибок в ErrorMessage.
func ParseErrorMessage(translation string) ErrorMessage {
	if r, d, ok := strings.Cut(translation, "\n\n"); ok {
		return ErrorMessage{
			Reason:  r,
			Details: d,
		}
	}

	return ErrorMessage{
		Reason: translation,
	}
}
