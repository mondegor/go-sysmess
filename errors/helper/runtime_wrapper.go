package helper

import (
	"errors"

	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// RuntimeErrorWrapper - помощник для оборачивания перехваченных ошибок.
	RuntimeErrorWrapper struct {
		internalError runtimeError
		systemError   runtimeError
		prewrapFunc   func(err error, attrs []any) (wrappedErr error, ok bool)
	}

	runtimeError interface {
		error
		Kind() kind.Enum
		Wrap(err error, attrs ...any) error
	}
)

// NewRuntimeErrorWrapper - создаёт объект RuntimeErrorWrapper.
// Ошибки совпадающие с internalError, systemError, пользовательские ошибки
// не оборачиваются и переданные атрибуты для них игнорируются.
func NewRuntimeErrorWrapper(
	internalError runtimeError,
	systemError runtimeError,
	prewrapFunc func(err error, attrs []any) (wrappedErr error, ok bool),
) (*RuntimeErrorWrapper, error) {
	if internalError.Kind() != kind.Internal {
		return nil, errors.New("internalError is not an Internal Error")
	}

	if systemError != nil {
		if systemError.Kind() != kind.System {
			return nil, errors.New("systemError is not a System Error")
		}
	}

	if prewrapFunc == nil {
		prewrapFunc = func(err error, _ []any) (wrappedErr error, ok bool) {
			return err, false
		}
	}

	return &RuntimeErrorWrapper{
		internalError: internalError,
		systemError:   systemError,
		prewrapFunc:   prewrapFunc,
	}, nil
}

// Wrap - возвращает ошибку, обёрнутую в internalError или в systemError в зависимости от её типа.
// Ошибки совпадающие с internalError, systemError, пользовательские ошибки
// не оборачиваются и переданные атрибуты для них игнорируются.
// Проверка вложенных ошибок не осуществляется.
func (w *RuntimeErrorWrapper) Wrap(err error, attrs ...any) error {
	if err, ok := w.prewrapFunc(err, attrs); ok {
		return err
	}

	switch kind.Extract(err) {
	case kind.User:
		return err
	case kind.System:
		if w.systemError != nil {
			// проверяются только ошибки первого уровня (необёрнутые)
			if w.systemError == err { //nolint:errorlint
				return err
			}

			return w.systemError.Wrap(err, attrs...)
		}
	default:
	}

	// проверяются только ошибки первого уровня (необёрнутые)
	if w.internalError == err { //nolint:errorlint
		return err
	}

	return w.internalError.Wrap(err, attrs...)
}
