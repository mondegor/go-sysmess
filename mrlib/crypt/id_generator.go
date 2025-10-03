package crypt

type (
	// IDGenerator - генератор уникальных идентификаторов.
	IDGenerator interface {
		GenerateID() string
	}

	// IDGeneratorFunc - генератор уникальных идентификаторов в виде функции.
	IDGeneratorFunc func() string
)

// GenerateID - реализация интерфейса IDGenerator в виде функции.
func (f IDGeneratorFunc) GenerateID() string {
	return f()
}
