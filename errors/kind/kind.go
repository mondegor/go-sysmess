package kind

import "errors"

// Перечисление типов ошибок.
const (
	Internal Enum = iota + 1 // внутренняя ошибка приложения (например: обращение по nil указателю)
	System                   // системная ошибка приложения (например: проблемы с сетью, с доступом к ресурсу)
	User                     // пользовательская ошибка (например: значение указанного поля некорректно)
)

type (
	// Enum - тип ошибки.
	Enum uint8
)

// String - возвращает тип ошибки в виде строки.
func (e Enum) String() string {
	switch e {
	case Internal:
		return "INTERNAL"
	case System:
		return "SYSTEM"
	case User:
		return "USER"
	}

	return "UNKNOWN"
}

// Extract - возвращает тип ошибки.
func Extract(err error) (userKind Enum) {
	if e, ok := err.(interface{ Kind() Enum }); ok {
		return e.Kind()
	}

	return 0
}

// Analyze - возвращает тип ошибки с учётом её вложенных ошибок.
// Если пользовательская ошибка содержит вложенную ошибку
// типа Internal или System, то вернётся этот тип ошибки.
func Analyze(err error) (userKind Enum) {
	for {
		if e, ok := err.(interface{ Kind() Enum }); ok {
			if e.Kind() != User {
				return e.Kind()
			}

			// выбирается причина пользовательской ошибки если такая существует
			if err = errors.Unwrap(err); err != nil {
				// запоминается, что пользовательская ошибка найдена
				userKind = User

				continue
			}

			// ошибка не содержит других ошибок, значит она пользовательская
			return User
		}

		return userKind
	}
}
