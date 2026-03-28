package wrap

import (
	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// CustomErrorWrapper - помощник для оборачивания пользовательских ошибок в ошибки типа Custom.
	CustomErrorWrapper interface {
		Wrap(err error) error
	}

	customErrorWrapper struct {
		wrapFunc    func(err error, code string) error
		code2custom map[string]string
	}
)

// NewCustomErrorWrapper - создаёт объект CustomErrorWrapper.
func NewCustomErrorWrapper(
	wrapperCustom func(err error, code string) error,
	codeCustom ...string,
) CustomErrorWrapper {
	code2custom := make(map[string]string, len(codeCustom)/2)

	for i := 0; i < len(codeCustom); i += 2 {
		code2custom[codeCustom[i]] = codeCustom[i+1]
	}

	if len(codeCustom)%2 != 0 {
		code2custom[codeCustom[len(codeCustom)-1]] = ""
	}

	return &customErrorWrapper{
		wrapFunc:    wrapperCustom,
		code2custom: code2custom,
	}
}

// Wrap - оборачивает пользовательские ошибки в ошибки типа Custom.
// Остальные ошибки оставляет без изменения.
func (w *customErrorWrapper) Wrap(err error) error {
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

type (
	nopCustomErrorWrapper struct{}
)

// NopCustomErrorWrapper - создаёт объект ErrorWrapper, который возвращает переданную ему ошибку как есть.
func NopCustomErrorWrapper() CustomErrorWrapper {
	return nopCustomErrorWrapper{}
}

// Wrap - возвращает указанную ошибку, реализуя CustomErrorWrapper интерфейс.
func (t nopCustomErrorWrapper) Wrap(err error) error { return err }
