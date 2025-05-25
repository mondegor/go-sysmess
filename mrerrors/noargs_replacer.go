package mrerrors

import (
	"io"
)

type (
	// noArgsRenderer - заглушка для имитации,
	// что в сообщении нет аргументов для рендеринга.
	noArgsRenderer struct {
		message []byte
	}
)

func newNoArgsRenderer(message string) *noArgsRenderer {
	return &noArgsRenderer{
		message: []byte(message),
	}
}

// ReplaceTo - записывает без каких либо изменений.
func (r *noArgsRenderer) ReplaceTo(wr io.Writer, _ []any) error {
	_, err := wr.Write(r.message)

	return err //nolint:wrapcheck
}

// CountArgs - возвращает всегда 0.
func (r *noArgsRenderer) CountArgs() int {
	return 0
}
