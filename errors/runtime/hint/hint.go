package hint

import (
	"github.com/mondegor/go-core/errors/kind"
)

type (
	// Hint - дополнительные данные, ассоциированные с runtime-ошибкой.
	// Содержит тип ошибки, уникальный ID и стек вызовов.
	// Используется для диагностики и логирования.
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

// New - создаёт Hint с указанными параметрами.
// Параметры:
//   - errorKind - тип ошибки (Internal, System);
//   - errorID - уникальный идентификатор экземпляра ошибки;
//   - stackTrace - стек вызовов на момент создания ошибки (может быть nil).
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

// ErrorKind - возвращает тип ошибки, для которой был создан этот Hint.
func (h *hint) ErrorKind() kind.Enum {
	return h.errorKind
}

// ErrorID - возвращает уникальный идентификатор экземпляра ошибки.
func (h *hint) ErrorID() string {
	return h.errorID
}

// StackTraceIterator - возвращает итератор по стеку вызовов.
// Если стек не был передан, возвращает итератор заглушку.
func (h *hint) StackTraceIterator() func() (index int, name, file string, line int) {
	if h.stackTrace == nil {
		return func() (index int, name, file string, line int) {
			return -1, "", "", 0
		}
	}

	return h.stackTrace.Iterator()
}
