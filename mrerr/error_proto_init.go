package mrerr

import "sync"

var proto = struct { //nolint:gochecknoglobals
	mu sync.Mutex

	// delayed - используется в момент запуска приложения, когда
	// создаются глобальные Proto переменные типа 'var ErrInternal = mrerr.NewProto()'
	// для того чтобы собрать весь список созданных таким образом ошибок и
	// проинициализировать его опциями по умолчанию с помощью функции InitDefaultOptions(),
	// далее система работает в обычном режиме.
	delayed []protoAppError

	// defaultProtoOptions - глобальный обработчик формирования списка опций по умолчанию
	// для создаваемых прототипов ошибок (при вызове NewProto()).
	defaultOptions ProtoOptionsHandler
}{}

// InitDefaultOptions - инициализация опциями по умолчанию глобальных Proto ошибок,
// созданных в момент запуска приложения. Если обработчик формирования опций по умолчанию не указан,
// то отложенный список Proto ошибок очищается и далее опции по умолчанию использоваться не будут.
// Функция может принимать от 1 до 2-х необязательных аргументов:
//   - первый - обработчик формирования опций по умолчанию применяемый к отложенному списку Proto ошибок;
//   - второй - обработчик формирования опций по умолчанию устанавливаемый в качестве глобального
//     и формирует персональный список опций по умолчанию при каждом последующем вызове функции NewProto(),
//     но если он не указан, то вместо него эту функцию будет выполнять обработчик из первого аргумента.
func InitDefaultOptions(handler ...ProtoOptionsHandler) {
	proto.mu.Lock()
	defer proto.mu.Unlock()

	if proto.defaultOptions != nil {
		return
	}

	if len(handler) > 0 {
		proto.defaultOptions = handler[0]
	}

	if proto.defaultOptions != nil {
		// устанавливаются опции по умолчанию для каждой созданной ошибки,
		// но только если они не были явно установлены ранее
		for _, wp := range proto.delayed {
			for _, opt := range proto.defaultOptions.Options(wp.p.code, wp.p.kind) {
				opt(&wp)
			}
		}
	}

	proto.delayed = nil

	// если указан второй аргумент
	if len(handler) > 1 {
		proto.defaultOptions = handler[1]

		return
	}

	// если ни один обработчик не указан, то устанавливается заглушка,
	// которая также блокирует повторный вызов функции InitDefaultOptions()
	proto.defaultOptions = ProtoOptionsHandlerFunc(
		func(_ string, _ ErrorKind) []ProtoOption {
			return nil
		},
	)
}
