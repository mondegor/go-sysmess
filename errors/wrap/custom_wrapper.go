package wrap

import (
	"github.com/mondegor/go-core/errors/kind"
)

type (
	// CustomErrorWrapper - помощник для оборачивания пользовательских ошибок
	// в ошибки типа Custom. Позволяет маппить коды стандартных пользовательских
	// ошибок в кастомные коды для слоя представления.
	CustomErrorWrapper interface {
		Wrap(err error) error
	}

	customErrorWrapper struct {
		wrapFunc    func(err error, code string) error
		code2custom map[string]string
	}
)

// NewCustomErrorWrapper - создаёт обёртку, транслирующую коды пользовательских ошибок в кастомные.
// Параметры:
//   - wrapperCustom - функция, создающая кастомную ошибку из исходной ошибки и кастомного кода;
//   - codeCustom - плоский слайс пар "исходныйКод(string)/кастомныйКод(string)",
//     например: ["userErr1", "fieldEmail", "userErr2", ""];
//
// Если количество элементов нечётное, последний элемент транслируется в пустой кастомный код.
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

// Wrap - оборачивает пользовательские ошибки (kind.User) в кастомные ошибки.
// Если код ошибки найден в таблице маппинга, вызывает wrapFunc с кастомным кодом.
// Ошибки других типов (System, Internal, без Kind) возвращаются без изменений.
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
	// nopCustomErrorWrapper - заглушка, реализующая интерфейс CustomErrorWrapper.
	// Возвращает переданную ошибку без изменений.
	nopCustomErrorWrapper struct{}
)

// NopCustomErrorWrapper - создаёт CustomErrorWrapper-заглушку.
func NopCustomErrorWrapper() CustomErrorWrapper {
	return nopCustomErrorWrapper{}
}

// Wrap - возвращает указанную ошибку, реализуя CustomErrorWrapper интерфейс.
func (t nopCustomErrorWrapper) Wrap(err error) error { return err }
