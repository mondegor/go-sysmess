package mrmsg

import (
	"io"
	"strings"

	"github.com/mondegor/go-core/util/conv"
)

const (
	missingArg = "!MISSINGARG"
)

type (
	// PlaceholderReplacer - замена именованных плейсхолдеров в сообщении через strings.Replacer.
	// Плейсхолдеры заменяются на строковые представления аргументов по порядку.
	PlaceholderReplacer struct {
		message      string
		placeholders []string
	}
)

// NewPlaceholderReplacer - создаёт PlaceholderReplacer для замены плейсхолдеров.
// Параметры:
//   - message - сообщение с именованными плейсхолдерами (например: "Hello {name}, you are {age}");
//   - placeholders - список плейсхолдеров для замены в порядке их появления.
func NewPlaceholderReplacer(message string, placeholders []string) *PlaceholderReplacer {
	return &PlaceholderReplacer{
		message:      message,
		placeholders: placeholders,
	}
}

// NewMessageReplacer - создаёт PlaceholderReplacer, автоматически извлекая плейсхолдеры из сообщения.
// Параметры:
//   - leftDelim, rightDelim - ограничители плейсхолдеров (например: "{" и "}");
//   - message - сообщение с плейсхолдерами для последующей замены.
func NewMessageReplacer(leftDelim, rightDelim, message string) *PlaceholderReplacer {
	return NewPlaceholderReplacer(
		message,
		NewPlaceholderExtractor(leftDelim, rightDelim).Extract(message),
	)
}

// Replace - подставляет аргументы в подготовленное сообщение, заменяя плейсхолдеры.
// Аргументы преобразуются в строки через conv.String.
// Если аргументов меньше, чем плейсхолдеров, недостающие заменяются на "!MISSINGARG".
func (p *PlaceholderReplacer) Replace(args []any) (replacedMessage string, err error) {
	if p.message == "" {
		return "", nil
	}

	if len(p.placeholders) == 0 {
		return p.message, nil
	}

	return strings.NewReplacer(p.merge(args)...).Replace(p.message), nil
}

// ReplaceTo - записывает сообщение с подставленными аргументами в io.Writer.
// Поведение обработки недостающих/лишних аргументов аналогично Replace.
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

// CountArgs - возвращает количество плейсхолдеров, ожидаемых сообщением.
func (p *PlaceholderReplacer) CountArgs() int {
	return len(p.placeholders)
}

func (p *PlaceholderReplacer) merge(args []any) (keyValues []string) {
	keyValues = make([]string, 0, len(p.placeholders)*2)

	if len(args) >= len(p.placeholders) {
		for i := range p.placeholders {
			keyValues = append(keyValues, p.placeholders[i], conv.String(args[i]))
		}

		return keyValues
	}

	i := 0

	// len(p.placeholders) > len(args)
	for range args {
		keyValues = append(keyValues, p.placeholders[i], conv.String(args[i]))
		i++
	}

	for j := i; j < len(p.placeholders); j++ {
		keyValues = append(keyValues, p.placeholders[j], missingArg)
	}

	return keyValues
}
