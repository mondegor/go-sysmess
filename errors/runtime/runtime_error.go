package runtime

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// ProtoError - прототип runtime-ошибки (внутренней или системной).
	// Служит фабрикой для создания конкретных экземпляров ошибок с атрибутами,
	// обёрнутыми ошибками и дополнительными деталями.
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
		onCreate func(kindErr kind.Enum, wrappedErr error) (hint any)
		attrs    []any
		hint     any
		err      error
	}
)

// ErrHasNilError - не указана ошибка для враппинга.
var ErrHasNilError = errors.New("runtime error has a nil wrapped error")

// New создаёт прототип runtime-ошибки указанного типа с заданным текстом.
// opts - дополнительные опции настройки (например, WithOnCreate для генерации hint).
// Возвращённый прототип можно использовать как фабрику конкретных ошибок.
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

// New - создаёт новый экземпляр ошибки на основе прототипа.
// Параметр attrs - пары ключ(string)/значение(any), прикрепляемые к ошибке.
func (e *protoError) New(attrs ...any) error {
	return e.newError(nil, "", attrs...)
}

// WithDetails - создаёт экземпляр ошибки, дополняя текст прототипа указанным описанием.
// Параметр attrs - пары ключ(string)/значение(any), прикрепляемые к ошибке.
func (e *protoError) WithDetails(details string, attrs ...any) error {
	return e.newError(nil, details, attrs...)
}

// Wrap - создаёт экземпляр ошибки, обёртывающий указанную ошибку.
// Если err == nil, подставляется ErrHasNilError.
// Параметр attrs - пары ключ(string)/значение(any), прикрепляемые к ошибке.
func (e *protoError) Wrap(err error, attrs ...any) error {
	if err == nil {
		err = ErrHasNilError
	}

	return e.newError(err, "", attrs...)
}

// WithError - создаёт экземпляр ошибки, обёртывающий указанную ошибку и дополняющий
// текст прототипа описанием. Если err == nil, подставляется ErrHasNilError.
// Параметр attrs - пары ключ(string)/значение(any), прикрепляемые к ошибке.
func (e *protoError) WithError(err error, details string, attrs ...any) error {
	if err == nil {
		err = ErrHasNilError
	}

	return e.newError(err, details, attrs...)
}

func (e *protoError) newError(err error, details string, attrs ...any) error {
	// оптимизация: чтобы не создавать лишний экземпляр ошибки
	if err == nil && details == "" && len(attrs) == 0 && e.onCreate == nil {
		return e
	}

	text := e.text
	if details != "" {
		text += ": " + details
	}

	var hint any
	if e.onCreate != nil {
		hint = e.onCreate(e.kind, err)
	}

	return &protoError{
		id:    e.id,
		kind:  e.kind,
		text:  text,
		attrs: attrs,
		hint:  hint,
		err:   err,
	}
}

// Kind - возвращает тип runtime-ошибки (Internal или System).
func (e *protoError) Kind() kind.Enum {
	return e.kind
}

// Attrs - возвращает копию атрибутов ошибки в виде плоского слайса пар ключ/значение.
// Возвращает nil, если атрибутов нет.
func (e *protoError) Attrs() []any {
	if len(e.attrs) == 0 {
		return nil
	}

	attrs := make([]any, len(e.attrs))
	copy(attrs, e.attrs)

	return attrs
}

// Hint - возвращает дополнительные данные, ассоциированные с ошибкой.
// Устанавливаются обработчиком onCreate() при создании экземпляра ошибки.
func (e *protoError) Hint() any {
	return e.hint
}

// Error - возвращает строковое представление ошибки.
func (e *protoError) Error() string {
	var key string

	attrExists := false
	attrs := e.attrs
	estimatedLen := len(e.text)

	// необходимо убедиться, что имеется хотя бы один атрибут для вывода
	for len(attrs) > 0 {
		key, _, attrs = popKeyValue(attrs)

		// если найден внутренний атрибут
		if strings.IndexByte(key, '.') == -1 {
			attrExists = true
			estimatedLen += 16 * len(e.attrs) // средний размер атрибутов

			break
		}
	}

	if e.err != nil {
		estimatedLen += 64 // средний размер вложенной ошибки
	}

	var buf strings.Builder

	buf.Grow(estimatedLen)
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
			buf.WriteString(formatAttr(value))
		}

		buf.WriteByte(']')
	}

	return buf.String()
}

// Is - реализует интерфейс errors.Is.
// Сравнивает прототипы ошибок по указателю id (общий прототип = одинаковые ошибки).
func (e *protoError) Is(target error) bool {
	if e == target {
		return true
	}

	if t, ok := target.(*protoError); ok {
		return e.id == t.id
	}

	return false
}

// Unwrap - реализует интерфейс errors.Unwrap.
// Возвращает вложенную ошибку или nil.
func (e *protoError) Unwrap() error {
	return e.err
}

// formatAttr - функция для небольшой оптимизации.
func formatAttr(value any) string {
	switch val := value.(type) {
	case string:
		return val
	case int:
		return strconv.FormatInt(int64(val), 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case bool:
		return strconv.FormatBool(val)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case nil:
		return "<NIL>"
	case error:
		return val.Error()
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%+v", val)
	}
}
