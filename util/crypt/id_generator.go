package crypt

type (
	// IDGenerator - генератор уникальных идентификаторов.
	IDGenerator interface {
		GenerateID() string
	}

	// IDGeneratorFunc - генератор уникальных идентификаторов в виде функции.
	IDGeneratorFunc func() string
)

// GenerateID - реализует интерфейс IDGenerator, вызывая саму функцию f.
// Позволяет использовать обычную функцию как генератор идентификаторов.
func (f IDGeneratorFunc) GenerateID() string {
	return f()
}
