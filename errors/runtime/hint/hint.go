package hint

import (
	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// Hint - содержит дополнительные данные, которые ассоциированы с ошибкой.
	Hint interface {
		ErrorKind() kind.Enum
		ErrorID() string
		StackTraceIterator() func() (index int, name, file string, line int)
	}

	hint struct {
		errorKind  kind.Enum
		errorID    string
		stackTrace stackTrace
	}

	stackTrace interface {
		Iterator() func() (index int, name, file string, line int)
	}
)

// New - создаёт объект Hint.
func New(
	errorKind kind.Enum,
	errorID string,
	stackTrace stackTrace,
) Hint {
	return &hint{
		errorKind:  errorKind,
		errorID:    errorID,
		stackTrace: stackTrace,
	}
}

// ErrorKind - возвращает тип ошибки.
func (h *hint) ErrorKind() kind.Enum {
	return h.errorKind
}

// ErrorID - возвращает ID ошибки.
func (h *hint) ErrorID() string {
	return h.errorID
}

// StackTraceIterator - возвращает итератор стектрейса.
func (h *hint) StackTraceIterator() func() (index int, name, file string, line int) {
	if h.stackTrace == nil {
		return func() (index int, name, file string, line int) {
			return -1, "", "", 0
		}
	}

	return h.stackTrace.Iterator()
}
