package runtime

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mondegor/go-sysmess/errors/kind"
)

//go:generate mockgen -source=runtime_error.go -destination=./mock/runtime_error.go

type (
	// ProtoError - внутренняя/системная ошибка.
	ProtoError interface {
		error

		Kind() kind.Enum
		New(attrs ...any) error
		WithDetails(details string, attrs ...any) error
		Wrap(err error, attrs ...any) error
		WithError(err error, details string, attrs ...any) error
	}

	protoError struct {
		id       *protoError
		kind     kind.Enum
		text     string
		onCreate func(kindErr kind.Enum, wrappedErr error) (bag any)
		attrs    []any
		hint     any
		err      error
	}
)

// ErrHasNilError - не указана ошибка для враппинга.
var ErrHasNilError = errors.New("runtime error has a nil wrapped error")

// New - создаёт объект ProtoError.
func New(errKind kind.Enum, text string, opts ...Option) ProtoError {
	return newProto(errKind, text, opts)
}

func newProto(errKind kind.Enum, text string, opts []Option) *protoError {
	p := &protoError{
		kind: errKind,
		text: text,
	}

	for _, opt := range opts {
		opt(p)
	}

	// id для сравнения порождённых ошибок
	p.id = p

	return p
}

// New - возвращает новую ошибку на основе текущей ProtoError ошибки.
func (e *protoError) New(attrs ...any) error {
	return e.newError(nil, "", attrs...)
}

// WithDetails - возвращает новую ошибку на основе текущей ProtoError ошибки с дополненным сообщением.
func (e *protoError) WithDetails(details string, attrs ...any) error {
	return e.newError(nil, details, attrs...)
}

// Wrap - возвращает новую ошибку на основе текущей ProtoError с обёрнутой указанной ошибкой.
func (e *protoError) Wrap(err error, attrs ...any) error {
	if err == nil {
		err = ErrHasNilError
	}

	return e.newError(err, "", attrs...)
}

// WithError - возвращает новую ошибку на основе текущей ProtoError с дополненным сообщением и обёрнутой указанной ошибкой.
func (e *protoError) WithError(err error, details string, attrs ...any) error {
	if err == nil {
		err = ErrHasNilError
	}

	return e.newError(err, details, attrs...)
}

func (e *protoError) newError(err error, details string, attrs ...any) error {
	// оптимизация, чтобы не создавать лишний экземпляр ошибки
	if err == nil && details == "" && len(attrs) == 0 && e.onCreate == nil {
		return e
	}

	c := *e
	c.err = err
	c.attrs = attrs

	if details != "" {
		c.text += ": " + details
	}

	if e.onCreate != nil {
		c.onCreate = nil
		c.hint = e.onCreate(e.kind, err)
	}

	return &c
}

// Kind - возвращает тип ошибки.
func (e *protoError) Kind() kind.Enum {
	return e.kind
}

// Attrs - возвращает попарно атрибуты прикреплённые к ошибке: ключ/значение.
//
// Attrs возвращает оригинал слайса, не копию!
func (e *protoError) Attrs() []any {
	return e.attrs
}

// Hint - возвращает дополнительные данные ассоциированные с ошибкой.
// Они устанавливаются во время создания экземпляра ошибки при вызове обработчика onCreate().
func (e *protoError) Hint() any {
	return e.hint
}

// Error - возвращает ошибку в виде строки.
func (e *protoError) Error() string {
	var key string

	attrExists := false
	attrs := e.attrs

	// необходимо убедиться, что имеется хотя бы один атрибут для вывода
	for len(attrs) > 0 {
		key, _, attrs = popKeyValue(attrs)

		// если найден внутренний атрибут
		if strings.IndexByte(key, '.') == -1 {
			attrExists = true

			break
		}
	}

	var buf strings.Builder

	buf.WriteString(e.text)

	if e.err != nil {
		buf.WriteString(": ")
		buf.WriteString(e.err.Error())
	}

	if attrExists {
		var value any

		attrs = e.attrs
		next := false

		buf.WriteString(" [")

		for len(attrs) > 0 {
			key, value, attrs = popKeyValue(attrs)

			// внешние атрибуты пропускаются
			if strings.IndexByte(key, '.') >= 0 {
				continue
			}

			if next {
				buf.WriteString(", ")
			} else {
				next = true
			}

			buf.WriteString(key)
			buf.WriteByte('=')
			buf.WriteString(fmt.Sprintf("%v", value))
		}

		buf.WriteByte(']')
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
