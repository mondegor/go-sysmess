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
		Wrap(err error) error
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

// Message - возвращает сообщение об ошибке (для поддержки локализации).
func (e *protoError) Message() string {
	return e.message[e.pos+len(codeMessageSeparator):]
}

// Args - всегда возвращает пустой слайс аргументов (для поддержки локализации).
func (e *protoError) Args() []any {
	return nil
}

// Code - возвращает код ошибки.
func (e *protoError) Code() string {
	return e.message[1:e.pos]
}

// Error - возвращает ошибку в виде строки.
func (e *protoError) Error() string {
	return e.message
}

// Is - сообщает, имеет ли указанная ошибка тот же
// прототип ошибки (errors.Is использует этот интерфейс).
func (e *protoError) Is(target error) bool {
	if e == target {
		return true
	}

	if t, ok := target.(*wrapError); ok {
		return e == t.proto
	}

	return false
}
