package mrerr

const (
	ErrorKindInternal ErrorKind = iota // внутренняя ошибка + traceID + call stack (by default)
	ErrorKindSystem                    // системная ошибка + traceID + call stack (by default)
	ErrorKindUser                      // пользовательская ошибка
)

type (
	ErrorKind int8
)

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
