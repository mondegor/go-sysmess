package baggage

type (
	// Option - настройка объекта baggage.
	Option func(o *options)

	options struct {
		baggage *Baggage
	}
)

// WithErrorID - устанавливает обработчик события создания экземпляра ошибки.
func WithErrorID(value string) Option {
	return func(o *options) {
		o.baggage.errorID = value
	}
}

// WithStackTrace - устанавливает обработчик события создания экземпляра ошибки.
func WithStackTrace(value stackTrace) Option {
	return func(o *options) {
		o.baggage.stackTrace = value
	}
}
