package helper

import (
	"errors"

	"github.com/mondegor/go-sysmess/errors/kind"
)

const (
	// ErrorMessageKindInternal - сообщение для пользователя: внутренняя ошибка приложения.
	ErrorMessageKindInternal = "internal error"

	// ErrorMessageKindSystem - сообщение для пользователя: системная ошибка приложения.
	ErrorMessageKindSystem = "system error"

	// ErrorMessageKindUnexpected - сообщение для пользователя: внутренняя необработанная ошибка приложения.
	ErrorMessageKindUnexpected = "unexpected internal error"
)

// ExtractMessageForLocalization - возвращает сообщение об ошибки и переданные ей аргументы
// для дальнейшего её локализации и отображения пользователю.
func ExtractMessageForLocalization(err error) (message string, args []any) {
	// сначала предполагается, что передана пользовательская ошибка
	if e, ok := err.(interface {
		Kind() kind.Enum
		Message() string
		Args() []any
	}); ok && e.Kind() == kind.User {
		return e.Message(), e.Args()
	}

	// для ошибок типа System и Internal и для необработанных ошибок
	// пользователю формируются фиксированные сообщения
	switch kind.Extract(err) {
	case kind.System:
		return ErrorMessageKindSystem, nil
	case kind.Internal:
		return ErrorMessageKindInternal, nil
	default:
		return ErrorMessageKindUnexpected, nil
	}
}

// ExtractAttrs - возвращает попарно все атрибуты (ключ/значение)
// прикреплённые данной ошибке и её вложенным ошибкам.
func ExtractAttrs(err error, filter func(key string) bool) []any {
	var n int

	unwrappedErr := err

	for {
		if e, ok := unwrappedErr.(interface{ Attrs() []any }); ok {
			n += len(e.Attrs())
		}

		if unwrappedErr = errors.Unwrap(unwrappedErr); unwrappedErr != nil {
			continue
		}

		break
	}

	if n == 0 {
		return nil
	}

	attrs := make([]any, 0, n)

	for {
		if e, ok := err.(interface{ Attrs() []any }); ok {
			errAttrs := e.Attrs()

			for len(errAttrs) > 1 {
				if key, ok := errAttrs[0].(string); ok {
					if filter(key) {
						attrs = append(attrs, errAttrs[0], errAttrs[1])
					}

					errAttrs = errAttrs[2:]
				} else {
					errAttrs = errAttrs[1:]
				}
			}
		}

		if err = errors.Unwrap(err); err != nil {
			continue
		}

		break
	}

	return attrs
}
