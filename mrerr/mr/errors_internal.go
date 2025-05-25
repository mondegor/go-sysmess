package mr

import "github.com/mondegor/go-sysmess/mrerr"

var (
	// ErrInternal - internal error,
	// обобщённая внутренняя ошибка системы которая может быть решена только силами разработки.
	// Для неё всегда должен формироваться стек вызовов и посылаться событие о её создании.
	ErrInternal = mrerr.NewKindInternal("internal error")

	// ErrUnexpectedInternal - unexpected internal error,
	// особая ошибка, в которую система заворачивает все ошибки отличные от типов ProtoError, InstantError.
	// Для неё не имеет смысла формировать стек вызовов, но всегда должно посылаться событие о её создании.
	// При возникновении этой ошибки нужно найти место причины её возникновения и написать для него обработку с указанием конкретной ошибки.
	ErrUnexpectedInternal = mrerr.NewKindInternal("unexpected internal error", mrerr.WithCode(ErrorCodeUnexpectedInternal), mrerr.WithDisabledCaller())

	// ErrInternalNilPointer - unexpected nil pointer.
	ErrInternalNilPointer = mrerr.NewKindInternal("unexpected nil pointer")

	// ErrInternalCaughtPanic - caught panic.
	ErrInternalCaughtPanic = mrerr.NewKindInternal("{Source}; panic: {Recover}; stackTrace: {StackTrace}")

	// ErrInternalTypeAssertion - invalid type assertion.
	ErrInternalTypeAssertion = mrerr.NewKindInternal("invalid type assertion (type='{Type}', value='{Value}')")

	// ErrInternalInvalidType - invalid type, expected type.
	ErrInternalInvalidType = mrerr.NewKindInternal("invalid type (current='{CurrentType}', expected='{ExpectedType}')")

	// ErrInternalUnexpectedValue - var has invalid value.
	ErrInternalUnexpectedValue = mrerr.NewKindInternal("variable has unexpected value (var='{VarName}', value='{Value}')")

	// ErrInternalUnhandledDefaultCase - unhandled default case.
	ErrInternalUnhandledDefaultCase = mrerr.NewKindInternal("unhandled default case")

	// ErrInternalKeyNotFoundInSource - key is not found in source.
	ErrInternalKeyNotFoundInSource = mrerr.NewKindInternal("key is not found in source (key='{Key}', source='{Source}')")

	// ErrInternalFailedToOpen - failed to open object.
	ErrInternalFailedToOpen = mrerr.NewKindInternal("failed to open object")

	// ErrInternalFailedToClose - failed to close object.
	ErrInternalFailedToClose = mrerr.NewKindInternal("failed to close object")

	// ErrInternalTimeoutPeriodHasExpired - the timeout period has expired.
	ErrInternalTimeoutPeriodHasExpired = mrerr.NewKindSystem("the timeout period has expired")

	// ErrInternalUnexpectedEOF - unexpected EOF.
	ErrInternalUnexpectedEOF = mrerr.NewKindInternal("unexpected EOF")

	// ErrInternalValueLenMax - value has length greater than max characters.
	ErrInternalValueLenMax = mrerr.NewKindInternal("value has length greater then max characters (cur='{CurLength}', max='{MaxLength}')")

	// ErrInternalValueNotMatchRegexpPattern - specified value does not match regexp pattern.
	ErrInternalValueNotMatchRegexpPattern = mrerr.NewKindInternal("specified value does not match regexp pattern (value='{Value}', pattern='{Pattern}')")
)
