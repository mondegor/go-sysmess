package userfast

import (
	"github.com/mondegor/go-core/errors/kind"
)

const (
	missingCode          = "!MISSINGCODE"
	codeMessageSeparator = " - "
)

type (
	// ProtoError - быстрый прототип пользовательской ошибки с поддержкой локализации.
	// Хранит код ошибки и сообщение в единой строке формата "#CODE - message".
	ProtoError interface {
		error

		Code() string
		Wrap(err error) error
	}

	// protoError - внутренняя реализация ProtoError.
	// Хранит код и сообщение в компактном формате: "#CODE - message".
	// pos - позиция конца кода в строке (для быстрого извлечения).
	protoError struct {
		message string
		pos     int
	}
)

// New - создаёт прототип пользовательской ошибки с указанным кодом и сообщением.
// Если code пустой, подставляется missingCode.
func New(code, message string) ProtoError {
	if code == "" {
		code = missingCode
	}

	return &protoError{
		message: "#" + code + codeMessageSeparator + message,
		pos:     len(code) + 1,
	}
}

// Wrap - оборачивает указанную ошибку в прототип пользовательской ошибки.
// Если err == nil, возвращает сам прототип.
func (e *protoError) Wrap(err error) error {
	if err == nil {
		return e
	}

	return &wrapError{
		proto: e,
		err:   err,
	}
}

// Kind - всегда возвращает kind.User.
func (e *protoError) Kind() kind.Enum {
	return kind.User
}

// Message - возвращает текст ошибки (без кода) для локализации.
func (e *protoError) Message() string {
	return e.message[e.pos+len(codeMessageSeparator):]
}

// Args - возвращает пустой слайс аргументов.
func (e *protoError) Args() []any {
	return nil
}

// Code - возвращает код ошибки для локализации.
func (e *protoError) Code() string {
	return e.message[1:e.pos]
}

// Error - возвращает строковое представление ошибки.
func (e *protoError) Error() string {
	return e.message
}

// Is - реализует интерфейс errors.Is.
// Сравнивает прототипы ошибок по указателю.
func (e *protoError) Is(target error) bool {
	if e == target {
		return true
	}

	if t, ok := target.(*wrapError); ok {
		return e == t.proto
	}

	return false
}
