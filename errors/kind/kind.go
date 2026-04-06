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
	// Определяет источник и характер ошибки: внутренняя, системная или пользовательская.
	Enum uint8
)

// String - возвращает строковое представление типа ошибки.
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

// Extract - извлекает тип ошибки из ошибки, реализующей интерфейс Kind().
// Возвращает 0, если ошибка не реализует этот интерфейс.
func Extract(err error) (userKind Enum) {
	if e, ok := err.(interface{ Kind() Enum }); ok {
		return e.Kind()
	}

	return 0
}

// Analyze - возвращает тип ошибки с учётом вложенных ошибок.
// Алгоритм:
//   - Если ошибка имеет тип Internal или System, то вернётся её тип.
//   - Если ошибка имеет тип User, разворачивает её (errors.Unwrap) и анализирует вложенную ошибку.
//   - Если развёрнутая ошибка - Internal или System, возвращает её тип.
//   - Если цепочка разворачивания закончилась без нахождения Internal/System, возвращает User.
//   - Если ошибка не реализует Kind(), возвращает последний найденный User или 0.
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
