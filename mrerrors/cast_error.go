package mrerrors

// Cast - приводит указанную ошибку к InstantError или если это невозможно,
// то вызывает функцию по умолчанию, указанную в параметре defFunc, затем возвращает результат.
func Cast(err error) (ie *InstantError, ok bool) {
	if err == nil {
		return nil, true
	}

	switch e := err.(type) { //nolint:errorlint
	case *InstantError:
		return e, true
	case *ProtoError:
		return CastProto(e), true
	}

	return nil, false
}

// CastOrWrap - приводит указанную ошибку к InstantError или если это невозможно,
// то вызывает функцию по умолчанию, указанную в параметре defFunc, затем возвращает результат.
func CastOrWrap(err error, wrapper func(err error) *InstantError) *InstantError {
	if e, ok := Cast(err); ok {
		return e
	}

	return wrapper(err)
}

// CastLessVerboseError - преобразует в ошибку InstantError без вызова
// обработчиков создания ошибки и формирования стека вызовов.
func CastLessVerboseError(err *InstantError) error {
	if err == nil {
		return nil
	}

	return (*wrappedError)(err)
}

// CastProto - преобразует в ошибку InstantError без вызова
// обработчиков создания ошибки и формирования стека вызовов.
func CastProto(err *ProtoError) *InstantError {
	if err == nil {
		return nil
	}

	return &InstantError{
		pureError: err.pureError,
	}
}
