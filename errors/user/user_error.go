package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mondegor/go-core/errors/kind"
)

const (
	missingCode  = "!MISSINGCODE"
	maxErrorArgs = 32
)

type (
	// ProtoError - прототип пользовательской ошибки с поддержкой аргументов и вложенных ошибок.
	// Позволяет создавать конкретные экземпляры с привязанными аргументами для локализации.
	// Используется в слоях бизнес-логики.
	ProtoError interface {
		error

		Code() string
		New(args ...any) error
		Wrap(err error, args ...any) error
	}

	// protoError - внутренняя реализация ProtoError.
	// Хранит код, шаблон сообщения, аргументы и вложенную ошибку.
	protoError struct {
		id      *protoError
		code    string
		message string
		args    []any
		err     error
	}
)

// ErrHasNilError - ошибка, подставляемая вместо nil при обёртывании.
var ErrHasNilError = errors.New("user error has a nil wrapped error")

// New - создаёт прототип пользовательской ошибки с указанным кодом и сообщением.
// Код и сообщение используются для локализации.
// Возвращённый прототип служит фабрикой конкретных ошибок с аргументами.
func New(code, message string) ProtoError {
	p := &protoError{
		code:    code,
		message: message,
	}

	p.id = p // id для сравнения порождённых ошибок

	if p.code == "" {
		p.code = missingCode
	}

	return p
}

// New - создаёт новый экземпляр ошибки на основе прототипа.
// Параметр args - аргументы, подставляемые в сообщение ошибки для локализации.
func (e *protoError) New(args ...any) error {
	return e.newError(nil, args...)
}

// Wrap создаёт экземпляр ошибки, обёртывающий указанную ошибку.
// args - аргументы, подставляемые в сообщение ошибки для локализации.
// Если err == nil, подставляется ErrHasNilError.
func (e *protoError) Wrap(err error, args ...any) error {
	if err == nil {
		err = ErrHasNilError
	}

	return e.newError(err, args...)
}

func (e *protoError) newError(err error, args ...any) error {
	// оптимизация: чтобы не создавать лишний экземпляр ошибки
	if err == nil && len(args) == 0 {
		return e
	}

	if len(args) > maxErrorArgs {
		args = args[:maxErrorArgs]
	}

	return &protoError{
		id:      e.id,
		code:    e.code,
		message: e.message,
		args:    args,
		err:     err,
	}
}

// Kind - возвращает тип ошибки - всегда kind.User.
func (e *protoError) Kind() kind.Enum {
	return kind.User
}

// Code - возвращает код ошибки для локализации.
func (e *protoError) Code() string {
	return e.code
}

// Message - возвращает исходный шаблон сообщения ошибки без подстановки аргументов.
// Используется системой локализации для получения переводимого текста.
func (e *protoError) Message() string {
	return e.message
}

// Args - возвращает копию аргументов ошибки для подстановки в сообщение.
// Используется системой локализации вместе с Message().
func (e *protoError) Args() []any {
	if len(e.args) == 0 {
		return nil
	}

	args := make([]any, len(e.args))
	copy(args, e.args)

	return args
}

// Error - возвращает строковое представление ошибки.
// Метод вызывается только для отладки; для пользовательского вывода
// предназначены Message() и Args().
func (e *protoError) Error() string {
	// оптимизация: когда нет аргументов и вложенной ошибки
	if len(e.args) == 0 && e.err == nil {
		return "#" + e.code + " - " + e.message
	}

	var buf strings.Builder

	estimatedLen := 4 + len(e.code) + len(e.message)

	if len(e.args) > 0 {
		// Формат: " [$arg1=X, $arg2=Y]"
		estimatedLen += 2 + 6*len(e.args) // 6 - средняя длина "$argN="

		for _, arg := range e.args {
			estimatedLen += estimatedArgLen(arg)
		}
	}

	if e.err != nil {
		estimatedLen += 2 + 48 // средний размер вложенной ошибки
	}

	buf.Grow(estimatedLen)

	buf.WriteByte('#')
	buf.WriteString(e.code)
	buf.WriteString(" - ")
	buf.WriteString(e.message)

	if len(e.args) > 0 {
		buf.WriteString(" [")

		var argNumberBuf [2]byte // maxErrorArgs

		for i, arg := range e.args {
			if i > 0 {
				buf.WriteString(", ")
			}

			buf.WriteString("$arg")
			buf.Write(strconv.AppendInt(argNumberBuf[:0], int64(i+1), 10)) // maxErrorArgs
			buf.WriteString(strconv.Itoa(i + 1))
			buf.WriteByte('=')
			buf.WriteString(formatArg(arg))
		}

		buf.WriteByte(']')
	}

	if e.err != nil {
		buf.WriteString(": ")
		buf.WriteString(e.err.Error())
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

func estimatedArgLen(arg any) int {
	switch v := arg.(type) {
	case string:
		return len(v)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return 10
	case bool, nil:
		return 5 // true/false or <nil>
	case float32, float64:
		return 15
	default:
		return 20
	}
}

// formatArg - преобразует аргумент в строковое представление для вывода в Error().
func formatArg(value any) string {
	switch val := value.(type) {
	case string:
		return val
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case fmt.Stringer:
		return val.String()
	default:
		return fmt.Sprintf("%v", value)
	}
}
