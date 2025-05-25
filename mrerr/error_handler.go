package mrerr

import (
	"context"

	"github.com/mondegor/go-sysmess/mrerrors"
)

type (
	// ErrorHandler - обработчик ошибок какого либо сервиса или приложения в целом.
	ErrorHandler struct {
		handler func(ctx context.Context, analyzedKind ErrorKind, err error)
		wrapper unknownErrorWrapper
	}

	unknownErrorWrapper func(err error) error
)

// NewErrorHandler - создаёт объект ErrorHandler.
func NewErrorHandler(handler func(ctx context.Context, analyzedKind ErrorKind, err error), wrapper unknownErrorWrapper) *ErrorHandler {
	if handler == nil {
		handler = func(_ context.Context, _ ErrorKind, _ error) {}
	}

	if wrapper == nil {
		wrapper = func(err error) error {
			return err
		}
	}

	return &ErrorHandler{
		handler: handler,
		wrapper: wrapper,
	}
}

// Handle - версия метода HandleWith без вызова дополнительного обработчика.
func (h *ErrorHandler) Handle(ctx context.Context, err error) {
	h.HandleWith(ctx, err, nil)
}

// HandleWith - анализирует ошибку, если ошибка типа ErrorKindUnknown, то оборачивает её,
// далее вызывает основной обработчик, который был указан в конструкторе, а затем
// обработчик extraHandler. В результате вызова этих обработчиков ошибка может быть,
// например, каким-то способом залогирована / отправлена во внешний источник /
// использована для правильного формирования ответа серверу и т.д.
func (h *ErrorHandler) HandleWith(ctx context.Context, err error, extraHandler func(analyzedKind ErrorKind, err error)) {
	analyzedKind := h.analyzeError(err)

	if analyzedKind == ErrorKindUnknown {
		err = h.wrapper(err)
	}

	h.handler(ctx, analyzedKind, err)

	if extraHandler != nil {
		extraHandler(analyzedKind, err)
	}
}

func (h *ErrorHandler) analyzeError(err error) ErrorKind {
	nestedErr := err
	foundUserError := false

	// вычисляется общий тип ошибки с учётом её вложенных ошибок
	for {
		if e, ok := nestedErr.(*mrerrors.InstantError); ok { //nolint:errorlint
			if e.Kind() != ErrorKindUser {
				return e.Kind()
			}

			foundUserError = true

			// выбирается причина пользовательской ошибки если такая существует
			if nestedErr = e.Unwrap(); nestedErr != nil {
				continue
			}

			// ошибка не содержит других ошибок, значит она пользовательская
			return ErrorKindUser
		}

		break
	}

	if foundUserError {
		return ErrorKindUserWithWrappedError
	}

	if e, ok := err.(*mrerrors.ProtoError); ok { //nolint:errorlint
		return e.Kind()
	}

	return ErrorKindUnknown
}
