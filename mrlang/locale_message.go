package mrlang

import (
    "strings"
)

type (
    ErrorMessage struct {
        Reason string `yaml:"reason"`
        Details []string `yaml:"details"`
    }
)

func (em *ErrorMessage) DetailsToString() string {
    switch len(em.Details) {
    case 0:
        return ""

    case 1:
        return em.Details[0]
    }

    return "- " + strings.Join(em.Details, "\n- ")
}
