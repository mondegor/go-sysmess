package hint

type (
	// Option - настройка объекта hint.
	Option func(o *options)

	options struct {
		hint *Hint
	}
)

// WithErrorID - устанавливает обработчик события создания экземпляра ошибки.
func WithErrorID(value string) Option {
	return func(o *options) {
		o.hint.errorID = value
	}
}

// WithStackTrace - устанавливает обработчик события создания экземпляра ошибки.
func WithStackTrace(value stackTrace) Option {
	return func(o *options) {
		o.hint.stackTrace = value
	}
}
