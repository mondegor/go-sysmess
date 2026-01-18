package baggage

type (
	// Baggage - дополнительные данные, которые привязываются к возникшей ошибке.
	Baggage struct {
		errorID    string
		stackTrace stackTrace
	}

	stackTrace interface {
		Iterator() func() (index int, name, file string, line int)
	}
)

// New - создаёт объект Baggage.
func New(opts ...Option) *Baggage {
	o := options{
		baggage: &Baggage{},
	}

	for _, opt := range opts {
		opt(&o)
	}

	if o.baggage.stackTrace == nil {
		o.baggage.stackTrace = stackTraceNop{}
	}

	return o.baggage
}

// ErrorID - возвращает уникальный код случившейся ошибки.
// Может быть добавлен в атрибуты при логировании, а также возвращён пользователю.
func (b *Baggage) ErrorID() string {
	return b.errorID
}

// StackTraceIterator - возвращает итератор стека вызовов функций.
func (b *Baggage) StackTraceIterator() func() (index int, name, file string, line int) {
	return b.stackTrace.Iterator()
}

type (
	stackTraceNop struct{}
)

// Iterator - всегда возвращает, что стек прочитан полностью.
func (s stackTraceNop) Iterator() func() (index int, name, file string, line int) {
	return func() (index int, name, file string, line int) {
		return -1, "", "", 0
	}
}
