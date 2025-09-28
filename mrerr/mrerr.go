package mrerr

import (
	"context"

	"github.com/mondegor/go-sysmess/mrerrors"
)

//go:generate mockgen -source=mrerr.go -destination=./mock/mrerr.go

// Виды ошибок.
const (
	ErrorKindUnknown              = ErrorKind(0)
	ErrorKindInternal             = mrerrors.ErrorKindInternal
	ErrorKindSystem               = mrerrors.ErrorKindSystem
	ErrorKindUser                 = mrerrors.ErrorKindUser
	ErrorKindUserWithWrappedError = mrerrors.ErrorKindUser + 1 // пользовательская ошибка в которой находится вложенная ошибка
)

type (
	// ErrorKind - алиас mrerrors.ErrorKind.
	ErrorKind = mrerrors.ErrorKind

	// StackTracer - алиас mrerrors.StackTracer.
	StackTracer = mrerrors.StackTracer

	// MessageReplacer - алиас mrerrors.MessageReplacer.
	MessageReplacer = mrerrors.MessageReplacer

	// ErrorHandler - обработчик ошибок.
	ErrorHandler interface {
		Handle(ctx context.Context, err error)
		HandleWith(ctx context.Context, err error, extraHandler func(analyzedKind ErrorKind, err error))
	}

	// ErrorWrapper - помощник для оборачивания ошибок.
	ErrorWrapper interface {
		WrapError(err error, attrs ...any) error
	}

	// UseCaseErrorWrapper - помощник для оборачивания UseCase ошибок.
	UseCaseErrorWrapper interface {
		IsNotFoundOrNotAffectedError(err error) bool
		WrapErrorFailed(err error, attrs ...any) error
		WrapErrorNotFoundOrFailed(err error, attrs ...any) error
	}
)
