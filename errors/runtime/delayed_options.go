package runtime

import (
	"sync"

	"github.com/mondegor/go-sysmess/errors/kind"
)

type (
	// OptionsHandler - обработчик формирования списка опций для указанных кодов ошибок.
	OptionsHandler interface {
		Options(kindErr kind.Enum, message string) []Option
	}

	// OptionsHandlerFunc - обработчик формирования списка опций в виде функции.
	OptionsHandlerFunc func(kindErr kind.Enum, message string) []Option
)

// Options - реализация интерфейса OptionsHandler в виде функции для формирования списка опций.
func (f OptionsHandlerFunc) Options(kindErr kind.Enum, message string) []Option {
	return f(kindErr, message)
}

var defopts = struct { //nolint:gochecknoglobals
	mu sync.Mutex

	// delayed - массив собирает все ошибки созданные через NewDelayed(), до момента вызова InitDelayedOptions().
	// Это позволяет, в момент запуска приложения, отследить создание всех таких ошибок в переменных
	// вида 'var Err*' для того чтобы проинициализировать их опциями по умолчанию
	// при запуске приложения с помощью метода InitDelayedOptions(). Далее все таки ошибки инициализируются
	// опциями по умолчанию сразу же при их создании, а этот массив более не используется.
	delayed []*protoError

	// defaultProtoOptions - глобальный обработчик формирования списка опций по умолчанию
	// для создаваемых прототипов ошибок (при вызове New()).
	handler OptionsHandler
}{}

// NewDelayed - создаёт объект ProtoError и запоминает его до вызова InitDelayedOptions
// чтобы была возможность изменить его опции.
func NewDelayed(errKind kind.Enum, text string, opts ...Option) ProtoError {
	p := newProto(errKind, text, opts)

	defopts.mu.Lock()
	defer defopts.mu.Unlock()

	// сначала происходит сбор создаваемых глобальных ProtoError ошибок в момент запуска приложения
	// чтобы их инициализировать нужными опциями, которые определяются приложением позже
	// это будет происходить до тех пор, пока не будет вызвана функция InitDelayedOptions()
	// далее опции по умолчанию применяются сразу
	if defopts.handler == nil {
		defopts.delayed = append(defopts.delayed, p)

		return p
	}

	// устанавливаются опции по умолчанию
	for _, opt := range defopts.handler.Options(p.kind, p.text) {
		opt(p)
	}

	return p
}

// InitDelayedOptions - с помощью указанного обработчика одноразово присваивает опции
// по умолчанию всем созданным через NewDelayed() ошибкам в момент инициализации приложения,
// при этом, не изменяет опции, которые были явно переданы в конструктор такой ошибки.
// После этого, этот обработчик запоминается и начинает вызываться каждый раз в момент
// создания очередной такой ошибки (в момент вызова NewDelayed()).
func InitDelayedOptions(handler OptionsHandler) {
	defopts.mu.Lock()
	defer defopts.mu.Unlock()

	if defopts.handler != nil {
		return
	}

	// если обработчик не указан, то устанавливается заглушка,
	// которая блокирует повторный вызов функции InitDelayedOptions()
	if handler == nil {
		defopts.handler = OptionsHandlerFunc(
			func(_ kind.Enum, _ string) []Option {
				return nil
			},
		)

		defopts.delayed = nil

		return
	}

	// устанавливаются опции по умолчанию для каждой созданной ошибки, в момент инициализации приложения,
	// но только если эти опции уже не были установлены в момент создания этих ошибок
	for _, p := range defopts.delayed {
		for _, opt := range handler.Options(p.kind, p.text) {
			opt(p)
		}
	}

	defopts.handler = handler
	defopts.delayed = nil
}
