package mrmsg

import (
	"strings"
)

type (
	// MessageFormatter - предназначен для форматирования аргументов сообщения.
	MessageFormatter struct {
		extractor *PlaceholderExtractor
		formatter func(placeholder string, index int) (newPlaceholder string)
	}
)

// NewMessageFormatter - создаёт объект MessageFormatter.
func NewMessageFormatter(leftDelim, rightDelim string, formatter func(placeholder string, index int) (newPlaceholder string)) *MessageFormatter {
	return &MessageFormatter{
		extractor: NewPlaceholderExtractor(leftDelim, rightDelim),
		formatter: formatter,
	}
}

// Format - возвращает шаблон сообщения с отформатированными аргументами подходящие
// для конкретных реплейсеров аргументов сообщений.
func (p *MessageFormatter) Format(message string) (formattedMessage string, newPlaceholders []string) {
	placeholders := p.extractor.Extract(message)

	if len(placeholders) == 0 || p.formatter == nil {
		return message, placeholders
	}

	oldNew := make([]string, 0, len(placeholders)*2)

	for i := range placeholders {
		oldNew = append(oldNew, placeholders[i], p.formatter(placeholders[i], i))
		placeholders[i] = oldNew[len(oldNew)-1]
	}

	return strings.NewReplacer(oldNew...).Replace(message), placeholders
}
