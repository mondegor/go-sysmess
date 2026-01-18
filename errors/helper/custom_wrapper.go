package helper

import (
	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// CustomErrorWrapper - помощник для оборачивания
	// пользовательских ошибок в ошибки типа Custom.
	CustomErrorWrapper struct {
		wrapFunc    func(err error, code string) error
		code2custom map[string]string
	}
)

// NewCustomErrorWrapper - создаёт объект CustomErrorWrapper.
func NewCustomErrorWrapper(
	wrapperCustom func(err error, code string) error,
	codeCustom ...string,
) *CustomErrorWrapper {
	code2custom := make(map[string]string, len(codeCustom)/2)

	for i := 0; i < len(codeCustom); i += 2 {
		code2custom[codeCustom[i]] = codeCustom[i+1]
	}

	if len(codeCustom)%2 != 0 {
		code2custom[codeCustom[len(codeCustom)-1]] = ""
	}

	return &CustomErrorWrapper{
		wrapFunc:    wrapperCustom,
		code2custom: code2custom,
	}
}

// Wrap - оборачивает пользовательские ошибки в ошибки типа Custom.
// Остальные ошибки оставляет без изменения.
func (w *CustomErrorWrapper) Wrap(err error) error {
	if err == nil {
		return w.wrapFunc(err, "")
	}

	if e, ok := err.(interface {
		Kind() kind.Enum
		Code() string
	}); ok && e.Kind() == kind.User { // TODO: ОБЯЗАТЕЛЬНО ПРОВЕРИТЬ!!!
		if customCode, ok := w.code2custom[e.Code()]; ok {
			return w.wrapFunc(err, customCode)
		}
	}

	return err
}
