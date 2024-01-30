package mrmsg

import "strings"

type (
	ErrorTranslator interface {
		HasErrorCode(code string) bool
		TranslateError(code, defaultMessage string, args ...NamedArg) ErrorMessage
	}

	ErrorMessage struct {
		Reason  string   `yaml:"reason"`
		Details []string `yaml:"details"`
	}
)

func (m *ErrorMessage) DetailsToString() string {
	switch len(m.Details) {
	case 0:
		return ""

	case 1:
		return m.Details[0]
	}

	return "- " + strings.Join(m.Details, "\n- ")
}
