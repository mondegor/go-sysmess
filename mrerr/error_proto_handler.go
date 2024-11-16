package mrerr

type (
	// ProtoOptionsHandler - обработчик формирования списка опций для указанных кода и типа ошибки.
	ProtoOptionsHandler interface {
		Options(code string, kind ErrorKind) []ProtoOption
	}

	// ProtoOptionsHandlerFunc - обработчик формирования списка опций в виде функции.
	ProtoOptionsHandlerFunc func(code string, kind ErrorKind) []ProtoOption
)

// Options - реализация интерфейса ProtoOptionsHandler для формирования списка опций.
func (f ProtoOptionsHandlerFunc) Options(code string, kind ErrorKind) []ProtoOption {
	return f(code, kind)
}
