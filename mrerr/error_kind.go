package mrerr

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
		GenerateIDFunc IDGeneratorFunc
		CallerFunc     CallerFunc
	}

	// ErrorOptions - опции для создания ErrorType.
	ErrorOptions struct {
		Kind            ErrorKind
		WithIDGenerator bool
		WithCaller      bool
	}

	// IDGeneratorFunc - генерация уникальных ID ошибок.
	IDGeneratorFunc func() string

	// CallerFunc - функция получения стека вызовов.
	CallerFunc func() StackTracer

	// StackTracer - предоставляет доступ к стеку вызовов.
	StackTracer interface {
		Count() int
		FileLine(index int) (file string, line int)
	}
)

var (
	// GlobalIDGeneratorFunc - глобальный объект для генерации ID экземпляра ошибки,
	// который может быть переопределён.
	GlobalIDGeneratorFunc IDGeneratorFunc = generateInstanceID

	// GlobalCallerFunc - глобальный объект для формирования текущего стека вызовов,
	// который может быть переопределён.
	GlobalCallerFunc CallerFunc = func() StackTracer { return newStackTrace() }

	// PrepareErrorTypeFunc - формирует тип ошибки на основе указанных для неё опций.
	PrepareErrorTypeFunc = func(opts ErrorOptions) ErrorType {
		etype := ErrorType{
			Kind: opts.Kind,
		}

		if opts.WithIDGenerator {
			etype.GenerateIDFunc = func() string {
				return GlobalIDGeneratorFunc()
			}
		}

		if opts.WithCaller {
			etype.CallerFunc = func() StackTracer {
				return GlobalCallerFunc()
			}
		}

		return etype
	}

	// ErrorTypeInternal - базовые настройки для типа "внутренняя ошибка приложения".
	ErrorTypeInternal = PrepareErrorTypeFunc(ErrorOptions{
		Kind:            ErrorKindInternal,
		WithIDGenerator: true,
		WithCaller:      true,
	})

	// ErrorTypeInternalNotice - настройки для типа "внутреннее предупреждение",
	// которое, в некоторых случаях, может стать поводом для реальной ошибки.
	ErrorTypeInternalNotice = PrepareErrorTypeFunc(ErrorOptions{
		Kind: ErrorKindInternal,
	})

	// ErrorTypeSystem - базовые настройки для типа "системная ошибка приложения".
	ErrorTypeSystem = PrepareErrorTypeFunc(ErrorOptions{
		Kind:            ErrorKindSystem,
		WithIDGenerator: true,
		WithCaller:      true,
	})

	// ErrorTypeUser - базовые настройки для типа "пользовательская ошибка".
	ErrorTypeUser = PrepareErrorTypeFunc(ErrorOptions{
		Kind: ErrorKindUser,
	})

	// ErrorTypeUserWithCaller - настройки для типа "пользовательская ошибка",
	// при котором требуется иноформация для фиксации ошибки.
	ErrorTypeUserWithCaller = PrepareErrorTypeFunc(ErrorOptions{
		Kind:            ErrorKindUser,
		WithIDGenerator: true,
		WithCaller:      true,
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
