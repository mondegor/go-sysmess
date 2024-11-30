package mrerr

import (
	"github.com/mondegor/go-sysmess/mrmsg"
)

//go:generate mockgen -source=error_proto.go -destination=./mock/error_proto.go

type (
	// ProtoAppError - прототип ошибки с поддержкой параметров, ID экземпляра ошибки и стека вызовов.
	ProtoAppError struct {
		pureError
		caller    func() StackTracer
		onCreated func(err *AppError) (instanceID string)
	}

	// StackTracer - предоставляет доступ к стеку вызовов.
	StackTracer interface {
		Count() int
		Item(index int) (name, file string, line int)
	}

	// protoAppError - объект используемый только в момент создания Proto ошибки
	// для того чтобы явно заданные опции не сбрасывались в значения по умолчанию.
	protoAppError struct {
		p                *ProtoAppError
		changedCode      bool
		changedCaller    bool
		changedOnCreated bool
	}
)

// ErrErrorIsNilPointer - указанная ошибка - nil pointer.
var ErrErrorIsNilPointer = NewProto(
	"errErrorIsNilPointer", ErrorKindInternal, "specified error is nil")

// NewProto - создаёт объект ProtoAppError.
func NewProto(code string, kind ErrorKind, message string, opts ...ProtoOption) *ProtoAppError {
	argsNames := mrmsg.ParseArgsNames(message)

	wp := protoAppError{
		p: &ProtoAppError{
			pureError: pureError{
				code:      code,
				kind:      kind,
				message:   message,
				argsNames: argsNames,

				// параметризованные сообщения по умолчанию
				// используют фиктивные значения параметров
				args: makeArgs(nil, argsNames),
			},
		},
	}

	for _, opt := range opts {
		opt(&wp)
	}

	proto.mu.Lock()
	defer proto.mu.Unlock()

	// сначала происходит сбор создаваемых глобальных Proto ошибок в момент запуска приложения
	// чтобы их инициализировать нужными опциями, которые определяются приложением позже
	// это будет происходить до тех пор, пока не будет вызвана функция InitDefaultOptions()
	// далее опции по умолчанию применяются сразу, но, как правило, функция NewProto() уже не вызывается
	if proto.defaultOptions == nil {
		proto.delayed = append(proto.delayed, wp)
	} else {
		// устанавливаются опции по умолчанию,
		// но только если они не были явно установлены ранее
		for _, opt := range proto.defaultOptions.Options(code, kind) {
			opt(&wp)
		}
	}

	return wp.p
}

// New - всегда создаёт новую копию текущего объекта,
// при этом вызываются функции caller и onCreated если они были установлены.
func (e *ProtoAppError) New(args ...any) *AppError {
	c := &AppError{
		pureError: e.pureError,
	}

	if len(args) > 0 {
		// если аргументов передано больше c.argsNames,
		// то при вызове Error() ошибки будет выведено предупреждение
		c.args = makeArgs(args, c.argsNames)
	}

	if e.caller != nil {
		c.stackTrace.val = e.caller()
		c.stackTrace.has = true
	}

	if e.onCreated != nil {
		c.instanceID = e.onCreated(c)
	}

	return c
}

// Wrap - создаёт новую ошибку на основе прототипа и оборачивает в неё указанную.
// Если указанная ошибка типа AppError, то проверяется, был ли у этой ошибки
// сгенерированы ID и стек, и если да, то у новой ошибки эти параметры не генерятся,
// даже если соответствующие генераторы установлены для этой ошибки.
func (e *ProtoAppError) Wrap(err error, args ...any) *AppError {
	if err == nil {
		err = ErrErrorIsNilPointer.New()
	}

	c := &AppError{
		pureError: e.pureError,
		err:       err,
	}

	if len(args) > 0 {
		// если аргументов передано больше c.argsNames,
		// то при вызове Error() ошибки будет выведено предупреждение
		c.args = makeArgs(args, c.argsNames)
	}

	if wrappedErr, ok := c.err.(*AppError); ok { //nolint:errorlint
		// instanceID is raising to the top
		if wrappedErr.errInstanceID != nil {
			c.errInstanceID = wrappedErr.errInstanceID
		} else if wrappedErr.instanceID != "" {
			c.errInstanceID = &wrappedErr.instanceID
		}

		// если стек был установлен во вложенном объекте,
		// то запрещается генерация стека текущему объекту
		if wrappedErr.stackTrace.has {
			c.stackTrace.has = true
		}
	}

	if e.caller != nil && !c.stackTrace.has {
		c.stackTrace.val = e.caller()
		c.stackTrace.has = true
	}

	if e.onCreated != nil && c.errInstanceID == nil {
		c.instanceID = e.onCreated(c)
	}

	return c
}

// Error - возвращает ошибку в виде строки.
func (e *ProtoAppError) Error() string {
	return mrmsg.MustRenderWithNamedArgs(e.message, e.getNamedArgs())
}
