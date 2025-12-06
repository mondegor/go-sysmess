package mrerrors

import (
	"errors"
)

type (
	// Wrapper - помощник для оборачивания перехваченных ошибок.
	Wrapper struct {
		wrapperInternal *ProtoError
		wrapperSystem   *ProtoError
		exceptions      []*ProtoError
	}
)

// NewWrapper - создаёт объект Wrapper.
func NewWrapper(
	wrapperInternal *ProtoError,
	wrapperSystem *ProtoError,
	exceptions ...*ProtoError,
) (*Wrapper, error) {
	if wrapperInternal.Kind() != ErrorKindInternal {
		return nil, errors.New("wrapperInternal is not an Internal Error")
	}

	exceptions = append(exceptions, wrapperInternal)

	if wrapperSystem != nil {
		if wrapperSystem.Kind() != ErrorKindSystem {
			return nil, errors.New("wrapperSystem is not an System Error")
		}

		exceptions = append(exceptions, wrapperSystem)
	}

	return &Wrapper{
		wrapperInternal: wrapperInternal,
		wrapperSystem:   wrapperSystem,
		exceptions:      exceptions,
	}, nil
}

// WrapError - возвращает ошибку, обёрнутую в wrapperInternal или в wrapperSystem в зависимости от её типа.
// Пользовательские ошибки и ошибки из exceptions не оборачиваются.
// Если указанная ошибка совпадает с wrapperInternal или wrapperSystem, то она только дополняется атрибутами.
func (w *Wrapper) WrapError(err error, attrs ...any) error {
	for _, ex := range w.exceptions {
		// если ошибка не находится в исключениях
		if !ex.Is(err) {
			continue
		}

		// если для ошибки не указаны атрибуты или у неё нет обработчика события о её создании
		if len(attrs) == 0 || ex.onCreated == nil {
			return err
		}

		e, _ := Cast(err)

		return e.WithAttrs(attrs...)
	}

	var kind ErrorKind

	if e, ok := err.(interface{ Kind() ErrorKind }); ok {
		kind = e.Kind()
	}

	if kind == ErrorKindUser {
		return err
	}

	if w.wrapperSystem != nil && kind == ErrorKindSystem {
		return w.wrapperSystem.Wrap(err, attrs...)
	}

	return w.wrapperInternal.Wrap(err, attrs...)
}
