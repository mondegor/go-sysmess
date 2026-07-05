package userfast

import (
	"errors"

	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// wrapError - обёртка пользовательской ошибки вокруг другой ошибки.
	// Сохраняет тот же прототип (код и сообщение), но добавляет вложенную ошибку.
	wrapError struct {
		proto *protoError
		err   error
	}
)

// Wrap - оборачивает указанную ошибку в тот же прототип.
// Если err == nil, возвращает ту же обёртку.
// Метод добавлен для того, чтобы был полностью реализован интерфейс ProtoError.
// Это позволяет errors.As находить ошибку по ProtoError в цепочке обёрток.
func (e *wrapError) Wrap(err error) error {
	if err == nil {
		return e
	}

	return &wrapError{
		proto: e.proto,
		err:   err,
	}
}

// Kind - всегда возвращает kind.User.
func (e *wrapError) Kind() kind.Enum {
	return kind.User
}

// Message - возвращает текст ошибки из прототипа (для локализации).
func (e *wrapError) Message() string {
	return e.proto.Message()
}

// Args - возвращает пустой слайс аргументов.
func (e *wrapError) Args() []any {
	return nil
}

// Code - возвращает код ошибки из прототипа.
func (e *wrapError) Code() string {
	return e.proto.Code()
}

// Error - возвращает строковое представление обёрнутой ошибки.
func (e *wrapError) Error() string {
	// e.err никогда не будет nil, т.к. wrapError
	// создаётся только с реальной ошибкой.
	return e.proto.message + ": " + e.err.Error()
}

// Is - реализует интерфейс errors.Is.
// Сравнивает: саму обёртку, прототип, другой wrapError с тем же прототипом,
// а также рекурсивно проверяет вложенную ошибку.
func (e *wrapError) Is(target error) bool {
	if e == target || e.proto == target {
		return true
	}

	if t, ok := target.(*wrapError); ok && e.proto == t.proto {
		return true
	}

	return errors.Is(e.err, target)
}

// Unwrap - реализует интерфейс errors.Unwrap.
// Возвращает вложенную ошибку.
func (e *wrapError) Unwrap() error {
	return e.err
}
