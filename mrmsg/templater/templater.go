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
	// Templater - шаблонизатор на основе html/template для рендеринга сообщений с параметрами.
	Templater struct {
		leftDelim  string
		rightDelim string
	}
)

// NewTemplater - создаёт объект Templater.
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

// Render - возвращает сформированное сообщение с заменёнными его аргументами на указанные значения.
func (p *Templater) Render(message string, data map[string]string) (string, error) {
	var buf strings.Builder

	if err := p.RenderTo(&buf, message, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderTo - записывает в указанный Writer сформированное сообщение
// с заменёнными его аргументами на указанные значения.
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
		return fmt.Errorf("parse message '%s' error: %w", message, err)
	}

	if err = t.Execute(wr, data); err != nil {
		return fmt.Errorf("render message '%s' error: %w", message, err)
	}

	return nil
}
