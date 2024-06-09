package mrerr

const (
	ErrorKindInternal ErrorKind = iota // ErrorKindInternal - внутренняя ошибка приложения (например: обращение по nil указателю)
	ErrorKindSystem                    // ErrorKindSystem - системная ошибка приложения (например: проблемы с сетью, с доступом к файлу)
	ErrorKindUser                      // ErrorKindUser - пользовательская ошибка (например: значение указанного поля некорректно)
)

type (
	// ErrorKind - вид ошибки.
	ErrorKind int8
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
