package mrprocess

import (
	"errors"
	"fmt"
	"runtime/debug"
)

// ErrSystemTemporaryProblemHasOccurred - сигнализирует о временной проблеме, операцию можно повторить.
var ErrSystemTemporaryProblemHasOccurred = errors.New("temporary problem has occurred")

type (
	// CaughtPanicError - ошибка перехваченной в воркере паники.
	// Содержит источник возникновения, значение recover() и стек вызовов на момент перехвата.
	CaughtPanicError struct {
		Source    string
		Recovered any
		Stack     string
	}
)

// NewCaughtPanicError - создаёт ошибку перехваченной паники с указанием источника,
// значения recover() и текущего стека вызовов.
func NewCaughtPanicError(source string, recovered any) error {
	return &CaughtPanicError{
		Source:    source,
		Recovered: recovered,
		Stack:     string(debug.Stack()),
	}
}

// Error - возвращает строковое представление ошибки.
func (e *CaughtPanicError) Error() string {
	return fmt.Sprintf(
		"caught panic [source=%s, recover=%v, stack_trace=%s]",
		e.Source, e.Recovered, e.Stack,
	)
}
