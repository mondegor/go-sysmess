package mrerr

const (
	ErrorKindInternal       ErrorKind = iota // внутренняя ошибка + traceID + call stack
	ErrorKindInternalNotice                  // внутреннее предупреждение, которое, в некоторых случаях, может стать поводом для реальной ошибки
	ErrorKindSystem                          // системная ошибка + traceID + call stack
	ErrorKindUser                            // пользовательская ошибка
)

type (
	ErrorKind int
)

func (k ErrorKind) String() string {
	switch k {
	case ErrorKindInternal:
		return "Internal"
	case ErrorKindInternalNotice:
		return "InternalNotice"
	case ErrorKindSystem:
		return "System"
	case ErrorKindUser:
		return "User"
	}

	return "Unknown"
}
