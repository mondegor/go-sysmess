package templater

import (
	"fmt"
	"html/template"
	"io"
	"strings"
)

const (
	leftDelimDefault  = "{{"
	rightDelimDefault = "}}"
)

type (
	// Templater - шаблонизатор сообщений на основе html/template.
	// Позволяет использовать синтаксис Go-шаблонов для подстановки параметров.
	Templater struct {
		leftDelim  string
		rightDelim string
	}
)

// NewTemplater - создаёт Templater с указанными ограничителями.
// Если leftDelim или rightDelim пусты, используются значения по умолчанию: "{{" и "}}".
func NewTemplater(leftDelim, rightDelim string) *Templater {
	if leftDelim == "" {
		leftDelim = leftDelimDefault
	}

	if rightDelim == "" {
		rightDelim = rightDelimDefault
	}

	return &Templater{
		leftDelim:  leftDelim,
		rightDelim: rightDelim,
	}
}

// Render - формирует сообщение из шаблона, подставляя параметры из data.
// Параметры:
//   - message - шаблон с синтаксисом html/template;
//   - data - карта имён параметров и их строковых значений;
func (p *Templater) Render(message string, data map[string]string) (string, error) {
	var buf strings.Builder

	if err := p.RenderTo(&buf, message, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderTo - записывает сформированное сообщение из шаблона в io.Writer.
// Параметры:
//   - message - шаблон с синтаксисом html/template;
//   - data - карта имён параметров и их строковых значений;
//
// Если сообщение не содержит ограничителей, записывается как есть без парсинга.
func (p *Templater) RenderTo(wr io.Writer, message string, data map[string]string) error {
	if message == "" {
		return nil
	}

	if !strings.Contains(message, p.leftDelim) {
		_, err := wr.Write([]byte(message))

		return err //nolint:wrapcheck
	}

	t, err := template.New("").Delims(p.leftDelim, p.rightDelim).Parse(message)
	if err != nil {
		return fmt.Errorf("parse message '%s': %w", message, err)
	}

	if err = t.Execute(wr, data); err != nil {
		return fmt.Errorf("render message '%s': %w", message, err)
	}

	return nil
}
