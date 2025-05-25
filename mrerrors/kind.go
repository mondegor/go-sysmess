package mrerrors

const (
	ErrorKindInternal ErrorKind = iota + 1 // ErrorKindInternal - внутренняя ошибка приложения (например: обращение по nil указателю).
	ErrorKindSystem                        // ErrorKindSystem - системная ошибка приложения (например: проблемы с сетью, с доступом к ресурсу).
	ErrorKindUser                          // ErrorKindUser - пользовательская ошибка (например: значение указанного поля некорректно).
)

type (
	// ErrorKind - вид ошибки.
	ErrorKind int8
)

// String - возвращает тип ошибки в виде строки.
func (k ErrorKind) String() string {
	switch k {
	case ErrorKindInternal:
		return "INTERNAL"
	case ErrorKindSystem:
		return "SYSTEM"
	case ErrorKindUser:
		return "USER"
	}

	return "UNKNOWN"
}
