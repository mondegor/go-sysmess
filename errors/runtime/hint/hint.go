package hint

type (
	// Hint - cодержит дополнительные данные, которые ассоциированы с ошибкой.
	Hint struct {
		errorID    string
		stackTrace stackTrace
	}

	stackTrace interface {
		Iterator() func() (index int, name, file string, line int)
	}
)

// New - создаёт объект Hint.
func New(opts ...Option) *Hint {
	o := options{
		hint: &Hint{},
	}

	for _, opt := range opts {
		opt(&o)
	}

	if o.hint.stackTrace == nil {
		o.hint.stackTrace = stackTraceNop{}
	}

	return o.hint
}

// ErrorID - возвращает уникальный код случившейся ошибки.
// Может быть добавлен в атрибуты при логировании, а также возвращён пользователю.
func (b *Hint) ErrorID() string {
	return b.errorID
}

// StackTraceIterator - возвращает итератор стека вызовов функций.
func (b *Hint) StackTraceIterator() func() (index int, name, file string, line int) {
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
