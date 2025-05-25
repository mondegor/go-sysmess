package mrerr

import "github.com/mondegor/go-sysmess/mrerrors"

const (
	ErrorKindUnknown              = ErrorKind(0)               // ErrorKindUnknown - специальный тип ошибки сигнализирующий о её нераспознаности.
	ErrorKindInternal             = mrerrors.ErrorKindInternal // ErrorKindInternal - алиас mrerrors.ErrorKindInternal.
	ErrorKindSystem               = mrerrors.ErrorKindSystem   // ErrorKindSystem - алиас mrerrors.ErrorKindSystem.
	ErrorKindUser                 = mrerrors.ErrorKindUser     // ErrorKindUser - алиас mrerrors.ErrorKindUser.
	ErrorKindUserWithWrappedError = mrerrors.ErrorKindUser + 1 // ErrorKindUserWithWrappedError - пользовательская ошибка в которой находится вложенная ошибка.
)

type (
	// ErrorKind - алиас mrerrors.ErrorKind.
	ErrorKind = mrerrors.ErrorKind

	// StackTracer - алиас mrerrors.StackTracer.
	StackTracer = mrerrors.StackTracer

	// MessageReplacer - алиас mrerrors.MessageReplacer.
	MessageReplacer = mrerrors.MessageReplacer
)
