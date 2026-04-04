package userfast

import (
	"github.com/mondegor/go-sysmess/errors/kind"
)

const (
	missingCode          = "!MISSINGCODE"
	codeMessageSeparator = " - "
)

type (
	// ProtoError - пользовательская ошибка с поддержкой локализации.
	// Используется в слоях бизнес логики.
	ProtoError interface {
		error

		Code() string
	}

	protoError struct {
		message string
		pos     int
	}
)

// New - создаёт объект ProtoError.
func New(code, message string) ProtoError {
	if code == "" {
		code = missingCode
	}

	return &protoError{
		message: "#" + code + codeMessageSeparator + message,
		pos:     len(code) + 1,
	}
}

// Kind - всегда возвращает kind.User.
func (e *protoError) Kind() kind.Enum {
	return kind.User
}

// Code - возвращает код ошибки.
func (e *protoError) Code() string {
	return e.message[1:e.pos]
}

// Message - возвращает сообщение об ошибке (для поддержки локализации).
func (e *protoError) Message() string {
	return e.message[e.pos+len(codeMessageSeparator):]
}

// Args - всегда возвращает пустой слайс аргументов (для поддержки локализации).
func (e *protoError) Args() []any {
	return nil
}

// Error - возвращает ошибку в виде строки.
func (e *protoError) Error() string {
	return e.message
}
