package mrerr

import "github.com/mondegor/go-sysmess/mrcaller"

const (
	// ErrorKindInternal - внутренняя ошибка приложения (например: обращение по nil указателю).
	ErrorKindInternal ErrorKind = iota

	// ErrorKindSystem - системная ошибка приложения (например: проблемы с сетью, с доступом к файлу).
	ErrorKindSystem

	// ErrorKindUser - пользовательская ошибка (например: значение указанного поля некорректно).
	ErrorKindUser
)

type (
	// ErrorKind - вид ошибки.
	ErrorKind int8

	// ErrorType - тип ошибки (вид ошибки + дополнительные свойства).
	ErrorType struct {
		Kind           ErrorKind
		GenerateIDFunc func() string
		CallerFunc     func(skip int) mrcaller.CallStack
	}

	// ErrorOptions - опции используемые для создания типа ошибки.
	ErrorOptions struct {
		Kind           ErrorKind
		HasIDGenerator bool
		HasCaller      bool
	}
)

var (
	// GlobalIDGenerator - глобальный объект для генерации ID экземпляра ошибки.
	GlobalIDGenerator ErrorIDGenerator = NewIDGenerator()

	// GlobalCaller - глобальный объект для формирования CallStack.
	GlobalCaller = mrcaller.New()

	// PrepareErrorTypeFunc - формирует тип ошибки на основе указанных для неё опций.
	PrepareErrorTypeFunc = func(opts ErrorOptions) ErrorType {
		etype := ErrorType{
			Kind: opts.Kind,
		}

		if opts.HasIDGenerator {
			etype.GenerateIDFunc = func() string {
				return GlobalIDGenerator.GenerateID()
			}
		}

		if opts.HasCaller {
			etype.CallerFunc = func(skip int) mrcaller.CallStack {
				return GlobalCaller.CallStack(skip + 1)
			}
		}

		return etype
	}

	// ErrorTypeInternal - базовые настройки для типа "внутренняя ошибка приложения".
	ErrorTypeInternal = PrepareErrorTypeFunc(ErrorOptions{
		Kind:           ErrorKindInternal,
		HasIDGenerator: true,
		HasCaller:      true,
	})

	// ErrorTypeInternalNotice - настройки для типа "внутреннее предупреждение",
	// которое, в некоторых случаях, может стать поводом для реальной ошибки.
	ErrorTypeInternalNotice = PrepareErrorTypeFunc(ErrorOptions{
		Kind: ErrorKindInternal,
	})

	// ErrorTypeSystem - базовые настройки для типа "системная ошибка приложения".
	ErrorTypeSystem = PrepareErrorTypeFunc(ErrorOptions{
		Kind:           ErrorKindSystem,
		HasIDGenerator: true,
		HasCaller:      true,
	})

	// ErrorTypeUser - базовые настройки для типа "пользовательская ошибка".
	ErrorTypeUser = PrepareErrorTypeFunc(ErrorOptions{
		Kind: ErrorKindUser,
	})

	// ErrorTypeUserWithCaller - настройки для типа "пользовательская ошибка",
	// при котором требуется иноформация для фиксации ошибки.
	ErrorTypeUserWithCaller = PrepareErrorTypeFunc(ErrorOptions{
		Kind:           ErrorKindUser,
		HasIDGenerator: true,
		HasCaller:      true,
	})
)

// String - возвращает вид ошибки в виде строки.
func (k ErrorKind) String() string {
	switch k {
	case ErrorKindInternal:
		return "Internal"
	case ErrorKindSystem:
		return "System"
	case ErrorKindUser:
		return "User"
	}

	return "Unknown"
}
