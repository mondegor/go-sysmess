package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mondegor/go-sysmess/errors/kind"
)

//go:generate mockgen -source=user_error.go -destination=./mock/user_error.go

const (
	missingCode = "!MISSINGCODE"
)

type (
	// ProtoError - пользовательская ошибка.
	// Используется в слоях бизнес логики и представления, поддерживает возможность.
	ProtoError interface {
		error

		New(args ...any) error
		Wrap(err error, args ...any) error
		Code() string
	}

	protoError struct {
		id      *protoError
		code    string
		message string
		args    []any
		err     error
	}
)

// ErrHasNilError - не указана ошибка для враппинга.
var ErrHasNilError = errors.New("user error has a nil wrapped error")

// New - создаёт объект ProtoError.
func New(code, message string) ProtoError {
	p := &protoError{
		code:    code,
		message: message,
	}

	// id для сравнения порождённых ошибок
	p.id = p

	if p.code == "" {
		p.code = missingCode
	}

	return p
}

// New - создаётся новая ошибка на основе текущей ProtoError ошибки.
func (e *protoError) New(args ...any) error {
	return e.newError(nil, args...)
}

// Wrap - создаёт новую ошибку на основе текущей ProtoError и оборачивает в неё указанную.
func (e *protoError) Wrap(err error, args ...any) error {
	if err == nil {
		err = ErrHasNilError
	}

	return e.newError(err, args...)
}

func (e *protoError) newError(err error, args ...any) error {
	// оптимизация, чтобы не создавать лишний экземпляр ошибки
	if err == nil && len(args) == 0 {
		return e
	}

	c := *e
	c.args = args
	c.err = err

	return &c
}

// Kind - возвращает всегда kind.User.
func (e *protoError) Kind() kind.Enum {
	return kind.User
}

// Code - возвращает код ошибки.
func (e *protoError) Code() string {
	return e.code
}

// Message - возвращает оригинальное сообщение без подстановки аргументов.
func (e *protoError) Message() string {
	return e.message
}

// Args - возвращает аргументы используемые в сообщении ошибки.
//
// Args возвращает оригинал слайса, не копию!
func (e *protoError) Args() []any {
	return e.args
}

// Error - возвращает ошибку в виде строки.
func (e *protoError) Error() string {
	var buf strings.Builder

	buf.WriteByte('#')
	buf.WriteString(e.code)
	buf.WriteString(" - ")
	buf.WriteString(e.message)

	if len(e.args) > 0 {
		buf.WriteString(" [")

		for i, arg := range e.args {
			if i > 0 {
				buf.WriteString(", ")
			}

			buf.WriteString(fmt.Sprintf("$arg%d=%v", i+1, arg))
		}

		buf.WriteByte(']')
	}

	if e.err != nil {
		buf.WriteString(": ")
		buf.WriteString(e.err.Error())
	}

	return buf.String()
}

// Is - сообщает, имеет ли указанная ошибка тот же
// прототип ошибки (errors.Is использует этот интерфейс).
func (e *protoError) Is(target error) bool {
	if e == target {
		return true
	}

	if t, ok := target.(*protoError); ok {
		return e.id == t.id
	}

	return false
}

// Unwrap - возвращает вложенную ошибку (errors.Is использует этот интерфейс).
func (e *protoError) Unwrap() error {
	return e.err
}
