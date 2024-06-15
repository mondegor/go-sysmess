package mrerr

import (
	"github.com/mondegor/go-sysmess/mrmsg"
)

const (
	ErrorCodeUnexpectedInternal = "errUnexpectedInternal" // ErrorCodeUnexpectedInternal - обобщённый код ошибки: внутренняя ошибка приложения
	ErrorCodeUnexpectedSystem   = "errUnexpectedSystem"   // ErrorCodeUnexpectedSystem - обобщённый код ошибки: системная ошибка приложения
)

type (
	pureError struct {
		code      string
		kind      ErrorKind
		message   string
		argsNames []string
		args      []any
	}

	// translator - интерфейс для работы с переводом ошибок на различные языки.
	translator interface {
		HasErrorCode(code string) bool
		TranslateError(code, defaultMessage string, args ...mrmsg.NamedArg) mrmsg.ErrorMessage
	}

	codeGetter interface {
		Code() string
	}
)

// Code - возвращает код ошибки.
func (e *pureError) Code() string {
	return e.code
}

// Kind - возвращает тип ошибки.
func (e *pureError) Kind() ErrorKind {
	return e.kind
}

// Is - проверяется что ошибка с указанным кодом (для возможности использования errors.Is).
func (e *pureError) Is(err error) bool {
	if v, ok := err.(codeGetter); ok {
		return e.code == v.Code()
	}

	return false
}

// Translate - возвращает сформированное сообщение предназначенное для пользователя.
func (e *pureError) Translate(t translator) mrmsg.ErrorMessage {
	if e.kind == ErrorKindUser || t.HasErrorCode(e.code) {
		return t.TranslateError(e.code, e.message, e.getNamedArgs()...)
	}

	// если об ошибке с указанным кодом ничего не знает translator,
	// то берётся текст соответствующей стандартной ошибки
	if e.kind == ErrorKindSystem {
		return t.TranslateError(ErrorCodeUnexpectedSystem, ErrorCodeUnexpectedSystem)
	}

	return t.TranslateError(ErrorCodeUnexpectedInternal, ErrorCodeUnexpectedInternal)
}

func (e *pureError) getNamedArgs() []mrmsg.NamedArg {
	namedArgs := make([]mrmsg.NamedArg, len(e.argsNames))

	for i, argName := range e.argsNames {
		namedArgs[i] = mrmsg.NamedArg{
			Name:  argName,
			Value: e.args[i],
		}
	}

	return namedArgs
}
