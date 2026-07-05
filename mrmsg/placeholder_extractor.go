package mrmsg

import (
	"regexp"
	"slices"
	"strings"
)

const (
	leftDelimDefault  = "{"
	rightDelimDefault = "}"
)

type (
	// PlaceholderExtractor - извлекает именованные плейсхолдеры из текста сообщения.
	// Плейсхолдер должен соответствовать формату: {Name}, где Name начинается с буквы
	// и содержит только буквы и цифры.
	PlaceholderExtractor struct {
		leftDelim  string
		rightDelim string
	}
)

var regexpPlaceholderBody = regexp.MustCompile(`^[A-Za-z][A-Za-z0-9]*$`)

// NewPlaceholderExtractor - создаёт PlaceholderExtractor с указанными ограничителями.
// Если leftDelim или rightDelim пусты, используются значения по умолчанию: "{" и "}".
func NewPlaceholderExtractor(leftDelim, rightDelim string) *PlaceholderExtractor {
	if leftDelim == "" {
		leftDelim = leftDelimDefault
	}

	if rightDelim == "" {
		rightDelim = rightDelimDefault
	}

	return &PlaceholderExtractor{
		leftDelim:  leftDelim,
		rightDelim: rightDelim,
	}
}

// Extract - извлекает все уникальные плейсхолдеры из сообщения.
// Плейсхолдер валидируется: тело должно соответствовать ^[A-Za-z][A-Za-z0-9]*$.
// Возвращает nil, если плейсхолдеры не найдены или невалидны.
// Плейсхолдеры возвращаются в порядке первого появления, дубликаты пропускаются.
func (p *PlaceholderExtractor) Extract(message string) (placeholders []string) {
	n := strings.Count(message, p.leftDelim)

	if n == 0 {
		return nil
	}

	placeholders = make([]string, 0, n)

	for {
		pos1 := strings.Index(message, p.leftDelim)

		if pos1 < 0 {
			break
		}

		messageWithLeftDelim := message[pos1:]
		message = message[pos1+len(p.leftDelim):]
		pos2 := strings.Index(message, p.rightDelim)

		if pos2 <= 0 {
			if pos2 < 0 {
				break
			}

			// skip: p.leftDelim + p.rightDelim
			continue
		}

		placeholder := messageWithLeftDelim[:pos2+len(p.leftDelim)+len(p.rightDelim)]
		placeholderBody := message[:pos2]

		if !regexpPlaceholderBody.MatchString(placeholderBody) {
			continue
		}

		message = message[pos2+len(p.rightDelim):]

		if p.placeholderIn(placeholder, placeholders) {
			continue
		}

		placeholders = append(placeholders, placeholder)
	}

	// из-за отсутствия p.rightDelim или невыполнения regexpPlaceholderBody
	// может ни одного плейсхолдера не найтись
	if len(placeholders) == 0 {
		return nil
	}

	return slices.Clip(placeholders)
}

func (p *PlaceholderExtractor) placeholderIn(placeholder string, placeholders []string) bool {
	for i := 0; i < len(placeholders); i++ {
		if placeholder == placeholders[i] {
			return true
		}
	}

	return false
}
