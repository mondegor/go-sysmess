package mrmsg

import (
	"io"
	"strings"

	"github.com/mondegor/go-sysmess/mrargs"
)

const (
	missingArg = "!MISSINGARG"
)

type (
	// PlaceholderReplacer - реплейсер параметров сообщения на основе strings.Replacer.
	PlaceholderReplacer struct {
		message      string
		placeholders []string
	}
)

// NewPlaceholderReplacer - создаёт объект PlaceholderReplacer.
func NewPlaceholderReplacer(message string, placeholders []string) *PlaceholderReplacer {
	return &PlaceholderReplacer{
		message:      message,
		placeholders: placeholders,
	}
}

// NewMessageReplacer - создаёт объект PlaceholderReplacer.
// Из указанного сообщения предварительно извлекаются аргументы помеченные указанными разделителями.
func NewMessageReplacer(leftDelim, rightDelim, message string) *PlaceholderReplacer {
	return NewPlaceholderReplacer(
		message,
		NewPlaceholderExtractor(leftDelim, rightDelim).Extract(message),
	)
}

// Replace - возвращает заранее подготовленное сообщение с заменёнными его аргументами на указанные значения.
func (p *PlaceholderReplacer) Replace(args []any) (replacedMessage string, err error) {
	if p.message == "" {
		return "", nil
	}

	if len(p.placeholders) == 0 {
		return p.message, nil
	}

	return strings.NewReplacer(p.merge(args)...).Replace(p.message), nil
}

// ReplaceTo - записывает в указанный Writer заранее подготовленное сообщение
// с заменёнными его аргументами на указанные значения.
func (p *PlaceholderReplacer) ReplaceTo(wr io.Writer, args []any) error {
	if p.message == "" {
		return nil
	}

	if len(p.placeholders) == 0 {
		_, err := wr.Write([]byte(p.message))

		return err //nolint:wrapcheck
	}

	_, err := wr.Write([]byte(strings.NewReplacer(p.merge(args)...).Replace(p.message)))

	return err //nolint:wrapcheck
}

// CountArgs - возвращает кол-во аргументов в подготовленном сообщении.
func (p *PlaceholderReplacer) CountArgs() int {
	return len(p.placeholders)
}

func (p *PlaceholderReplacer) merge(args []any) (keyValues []string) {
	keyValues = make([]string, 0, len(p.placeholders)*2)

	if len(args) >= len(p.placeholders) {
		for i := range p.placeholders {
			keyValues = append(keyValues, p.placeholders[i], mrargs.ToString(args[i]))
		}

		return keyValues
	}

	i := 0

	// len(p.placeholders) > len(args)
	for range args {
		keyValues = append(keyValues, p.placeholders[i], mrargs.ToString(args[i]))
		i++
	}

	for j := i; j < len(p.placeholders); j++ {
		keyValues = append(keyValues, p.placeholders[j], missingArg)
	}

	return keyValues
}
