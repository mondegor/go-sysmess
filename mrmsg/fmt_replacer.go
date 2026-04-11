package mrmsg

import (
	"fmt"
	"io"
)

type (
	// FmtReplacer - замена плейсхолдеров в сообщении через fmt.Sprintf.
	// Использует формат fmt, где плейсхолдеры заменяются позиционными аргументами.
	FmtReplacer struct {
		message      string
		placeholders []string
	}
)

// NewFmtReplacer - создаёт FmtReplacer для замены плейсхолдеров в сообщении.
// Параметры:
//   - message - строка формата для fmt.Sprintf (например: "Hello %[1]s, you are %[2]d");
//   - placeholders - список плейсхолдеров, которые будут заменены аргументами.
func NewFmtReplacer(message string, placeholders []string) *FmtReplacer {
	return &FmtReplacer{
		message:      message,
		placeholders: placeholders,
	}
}

// Replace - подставляет аргументы в подготовленное сообщение через fmt.Sprintf.
// Если аргументов меньше, чем плейсхолдеров, недостающие заменяются на "!MISSINGARG".
// Если аргументов больше, лишние игнорируются.
func (p *FmtReplacer) Replace(args []any) (replacedMessage string, err error) {
	if p.message == "" {
		return "", nil
	}

	if len(p.placeholders) == 0 {
		return p.message, nil
	}

	return fmt.Sprintf(p.message, p.correctArgs(args)...), nil
}

// ReplaceTo - записывает сообщение с подставленными аргументами в io.Writer.
// Поведение обработки недостающих/лишних аргументов аналогично Replace.
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

// CountArgs - возвращает количество плейсхолдеров, ожидаемых сообщением.
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
