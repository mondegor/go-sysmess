package mrerr

type (
	// ProtoOption - настройка объекта ProtoAppError.
	ProtoOption func(e *protoAppError)
)

// WithProtoCode - устанавливает код ошибки.
func WithProtoCode(value string) ProtoOption {
	return func(e *protoAppError) {
		if e.changedCode || value == "" {
			return
		}

		e.changedCode = true
		e.p.code = value
	}
}

// WithProtoCaller - устанавливает формирование callstack при создании экземпляра ошибки.
func WithProtoCaller(value func() StackTracer) ProtoOption {
	return func(e *protoAppError) {
		if e.changedCaller {
			return
		}

		e.changedCaller = true
		e.p.caller = value
	}
}

// WithProtoOnCreated - устанавливает обработчик события создания экземпляра ошибки.
func WithProtoOnCreated(value func(err *AppError) (instanceID string)) ProtoOption {
	return func(e *protoAppError) {
		if e.changedOnCreated {
			return
		}

		e.changedOnCreated = true
		e.p.onCreated = value
	}
}
