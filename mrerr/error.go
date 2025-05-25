package mrerr

import (
	"sync"

	"github.com/mondegor/go-sysmess/mrerrors"
)

type (
	// options - объект используемый только в момент создания mrerrors.ProtoError ошибки
	// для того чтобы явно заданные опции не сбрасывались в значения по умолчанию.
	options struct {
		proto               *mrerrors.ProtoError
		changedArgsReplacer bool
		changedCaller       bool
		changedOnCreated    bool
	}
)

var defopts = struct { //nolint:gochecknoglobals
	mu sync.Mutex

	// delayed - массив собирает все ошибки созданные через New(), до момента вызова InitDefaultOptions().
	// Это позволяет, в момент запуска приложения, отследить создание всех таких ошибок в переменных
	// вида 'var Err* = mrerr.New(...)' для того чтобы проинициализировать их опциями по умолчанию
	// при запуске приложения с помощью метода InitDefaultOptions(). Далее все таки ошибки инициализируются
	// опциями по умолчанию сразу же при их создании, а этот массив более не используется.
	delayed []options

	// defaultProtoOptions - глобальный обработчик формирования списка опций по умолчанию
	// для создаваемых прототипов ошибок (при вызове New()).
	handler OptionsHandler
}{}

// NewKindInternal - создаёт фабрику ProtoError для создания ошибок типа Internal с указанными опциями.
func NewKindInternal(message string, opts ...Option) *mrerrors.ProtoError {
	return New(ErrorKindInternal, message, opts...)
}

// NewKindSystem - создаёт фабрику ProtoError для создания ошибок типа System с указанными опциями.
func NewKindSystem(message string, opts ...Option) *mrerrors.ProtoError {
	return New(ErrorKindSystem, message, opts...)
}

// NewKindUser - создаёт фабрику ProtoError для создания ошибок типа User с указанными опциями.
func NewKindUser(code, message string, opts ...Option) *mrerrors.ProtoError {
	return New(ErrorKindUser, message, append(opts, WithCode(code))...)
}

// New - создаёт фабрику ProtoError для создания ошибок указанного типа с указанными опциями.
func New(kind ErrorKind, message string, opts ...Option) *mrerrors.ProtoError {
	wp := options{
		proto: mrerrors.NewProto(message, mrerrors.WithProtoKind(kind)),
	}

	for _, opt := range opts {
		opt(&wp)
	}

	defopts.mu.Lock()
	defer defopts.mu.Unlock()

	// сначала происходит сбор создаваемых глобальных Proto ошибок в момент запуска приложения
	// чтобы их инициализировать нужными опциями, которые определяются приложением позже
	// это будет происходить до тех пор, пока не будет вызвана функция InitDefaultOptions()
	// далее опции по умолчанию применяются сразу
	if defopts.handler == nil {
		defopts.delayed = append(defopts.delayed, wp)

		return wp.proto
	}

	// устанавливаются опции по умолчанию,
	// но только если они не были явно переданы в данный метод
	for _, opt := range defopts.handler.Options(kind, wp.proto.Code(), wp.proto.Message()) {
		opt(&wp)
	}

	return wp.proto
}

// InitDefaultOptions - с помощью указанного обработчика одноразово присваивает опции
// по умолчанию всем созданным через New() ошибкам в момент инициализации приложения,
// при этом, не изменяет опции, которые были явно переданы в конструктор такой ошибки.
// После этого, этот обработчик сохраняется и начинает вызываться каждый раз в момент
// создания очередной такой ошибки.
func InitDefaultOptions(handler OptionsHandler) {
	defopts.mu.Lock()
	defer defopts.mu.Unlock()

	if defopts.handler != nil {
		return
	}

	// если обработчик не указан, то устанавливается заглушка,
	// которая блокирует повторный вызов функции InitDefaultOptions()
	if handler == nil {
		defopts.handler = OptionsHandlerFunc(
			func(_ ErrorKind, _, _ string) []Option {
				return nil
			},
		)

		defopts.delayed = nil

		return
	}

	// устанавливаются опции по умолчанию для каждой созданной ошибки, в момент инициализации приложения,
	// но только если эти опции уже не были установлены в момент создания этих ошибок
	for _, wp := range defopts.delayed {
		for _, opt := range handler.Options(wp.proto.Kind(), wp.proto.Code(), wp.proto.Message()) {
			opt(&wp)
		}
	}

	defopts.handler = handler
	defopts.delayed = nil
}
