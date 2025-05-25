package mrerrors

import "io"

type (
	// MessageReplacer - интерфейс для замены аргументов содержащихся
	// в сообщении на конкретные значения переданные при создании экземпляра ошибки.
	MessageReplacer interface {
		ReplaceTo(wr io.Writer, args []any) error
		CountArgs() int
	}

	// pureError - базовая часть ошибки хранящая её неизменяемые свойства,
	// и по которой определяется уникальность ошибки этого типа.
	pureError struct {
		kind            ErrorKind
		code            string
		message         string
		messageReplacer MessageReplacer
	}
)

// Kind - возвращает тип ошибки.
func (e *pureError) Kind() ErrorKind {
	return e.kind
}

// Code - возвращает код ошибки.
func (e *pureError) Code() string {
	return e.code
}

// Message - возвращает оригинальное сообщение без подстановки аргументов.
func (e *pureError) Message() string {
	return e.message
}
