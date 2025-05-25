package mrmsg

import (
	"fmt"
	"io"
)

type (
	// FmtReplacer - реплейсер параметров сообщения на основе fmt.
	FmtReplacer struct {
		message      string
		placeholders []string
	}
)

// NewFmtReplacer - создаёт объект FmtReplacer.
func NewFmtReplacer(message string, placeholders []string) *FmtReplacer {
	return &FmtReplacer{
		message:      message,
		placeholders: placeholders,
	}
}

// Replace - возвращает заранее подготовленное сообщение с заменёнными его аргументами на указанные значения.
func (p *FmtReplacer) Replace(args []any) (replacedMessage string, err error) {
	if p.message == "" {
		return "", nil
	}

	if len(p.placeholders) == 0 {
		return p.message, nil
	}

	return fmt.Sprintf(p.message, p.correctArgs(args)...), nil
}

// ReplaceTo - записывает в указанный Writer заранее подготовленное сообщение
// с заменёнными его аргументами на указанные значения.
func (p *FmtReplacer) ReplaceTo(wr io.Writer, args []any) error {
	if p.message == "" {
		return nil
	}

	if len(p.placeholders) == 0 {
		_, err := wr.Write([]byte(p.message))

		return err //nolint:wrapcheck
	}

	_, err := fmt.Fprintf(wr, p.message, p.correctArgs(args)...)

	return err //nolint:wrapcheck
}

// CountArgs - возвращает кол-во аргументов в подготовленном сообщении.
func (p *FmtReplacer) CountArgs() int {
	return len(p.placeholders)
}

func (p *FmtReplacer) correctArgs(args []any) []any {
	if len(args) >= len(p.placeholders) {
		return args[:len(p.placeholders)]
	}

	// len(p.placeholders) > len(args)
	for i := len(args); i < len(p.placeholders); i++ {
		args = append(args, missingArg)
	}

	return args
}
