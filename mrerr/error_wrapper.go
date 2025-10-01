package mrerr

type (
	// ErrorWrapper - помощник для оборачивания ошибок.
	ErrorWrapper interface {
		WrapError(err error, attrs ...any) error
	}

	// UseCaseErrorWrapper - помощник для оборачивания ошибок
	// в зависимости от их вида. Используется в UseCase.
	UseCaseErrorWrapper interface {
		IsNotFoundOrNotAffectedError(err error) bool
		WrapErrorFailed(err error, attrs ...any) error
		WrapErrorNotFoundOrFailed(err error, attrs ...any) error
	}
)
