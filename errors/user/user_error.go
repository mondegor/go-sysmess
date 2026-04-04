package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/errors/kind"
)

const (
	missingCode  = "!MISSINGCODE"
	maxErrorArgs = 32
)

type (
	// ProtoError - пользовательская ошибка с поддержкой аргументов и враппинга ошибки.
	// Используется в слоях бизнес логики.
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

	p.id = p // id для сравнения порождённых ошибок

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

// Kind - возвращает всегда kind.User.
func (e *protoError) Kind() kind.Enum {
	return kind.User
}

// Code - возвращает код ошибки.
func (e *protoError) Code() string {
	return e.code
}

// Message - возвращает оригинальное сообщение
// без подстановки аргументов (для поддержки локализации).
func (e *protoError) Message() string {
	return e.message
}

// Args - возвращает аргументы используемые
// в сообщении ошибки (для поддержки локализации).
func (e *protoError) Args() []any {
	if len(e.args) == 0 {
		return nil
	}

	args := make([]any, len(e.args))
	copy(args, e.args)

	return args
}

// Error - возвращает ошибку в виде строки.
// Сообщение об ошибке формируется налету, потому как
// в пользовательских ошибках метод Error() вызывается только
// для отладки, а ошибки для пользователя формируется с помощью
// методов Message() и Args().
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

// formatArg - функция для небольшой оптимизации.
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
