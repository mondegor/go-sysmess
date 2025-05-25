package mr

import "github.com/mondegor/go-sysmess/mrerr"

type (
	errorMessage interface {
		Kind() mrerr.ErrorKind
		Code() string
		Message() string
		Args() []any
	}
)

// ErrorToMessage - возвращает функцию для преобразования ошибки в сообщение для отображения её пользователю.
func ErrorToMessage(displayedCodes ...string) func(err error) (message string, args []any) {
	if len(displayedCodes) == 0 {
		displayedCodes = append(
			displayedCodes,
			ErrorCodeUnexpectedInternal,
			ErrorCodeTemporarilyUnavailable,
		)
	}

	code2ok := make(map[string]bool, len(displayedCodes))

	for _, code := range displayedCodes {
		code2ok[code] = true
	}

	return func(err error) (message string, args []any) {
		if e, ok := err.(errorMessage); ok {
			// если детали ошибки можно отобразить пользователю
			if e.Kind() == mrerr.ErrorKindUser || code2ok[e.Code()] {
				return e.Message(), e.Args()
			}

			if e.Kind() == mrerr.ErrorKindSystem {
				return DefaultErrorCodeSystem, nil
			}
		}

		return DefaultErrorCodeInternal, nil
	}
}
