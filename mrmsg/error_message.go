package mrmsg

import "strings"

type (
	// ErrorMessage - сообщение об ошибке с её деталями.
	ErrorMessage struct {
		Reason  string   `yaml:"reason"`
		Details []string `yaml:"details"`
	}
)

// DetailsToString - преобразовывает список деталей ошибки в строку.
func (m *ErrorMessage) DetailsToString() string {
	switch len(m.Details) {
	case 0:
		return ""

	case 1:
		return m.Details[0]
	}

	return "- " + strings.Join(m.Details, "\n- ")
}
