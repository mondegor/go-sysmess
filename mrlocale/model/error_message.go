package model

import (
	"strings"
)

type (
	// ErrorMessage - модель ошибки используемая при её локализации.
	ErrorMessage struct {
		Reason     string
		Details    string
		ProblemURL string
	}
)

// String - возвращает причину ошибки и её подробностей в виде строки.
func (m ErrorMessage) String() string {
	if m.ProblemURL == "" {
		if m.Details == "" {
			return m.Reason
		}

		return m.Reason + "\n\n" + m.Details
	}

	return m.Reason + "\n\n" + m.Details + "\n\n" + m.ProblemURL
}

// ParseErrorMessage - парсинг строки хранящейся в справочнике ошибок в ErrorMessage.
func ParseErrorMessage(translation string) ErrorMessage {
	if r, d, ok := strings.Cut(translation, "\n\n"); ok {
		if d, p, ok := strings.Cut(d, "\n\n"); ok {
			return ErrorMessage{
				Reason:     r,
				Details:    d,
				ProblemURL: p,
			}
		}

		return ErrorMessage{
			Reason:  r,
			Details: d,
		}
	}

	return ErrorMessage{
		Reason: translation,
	}
}
