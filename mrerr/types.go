package mrerr

import (
	"github.com/mondegor/go-sysmess/mrerrors"
)

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
)
