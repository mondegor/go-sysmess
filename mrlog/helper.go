package mrlog

type (
	// Attr - атрибут, который указывается в аргументах
	// сообщения при его логировании.
	Attr struct {
		Key   string
		Value any
	}
)

// Err - обёртка для указания ошибки в качестве аргумента
// сообщения при его логировании.
func Err(err error) Attr {
	return Attr{
		Key:   "error",
		Value: err,
	}
}
