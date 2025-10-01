package crypt

type (
	// IDGenerator - генератор уникальных идентификаторов.
	IDGenerator interface {
		GenerateID() string
	}

	// IDGeneratorFunc - реализация интерфейса IDGenerator.
	IDGeneratorFunc func() string
)

// GenerateID - реализация интерфейса IDGenerator.
func (f IDGeneratorFunc) GenerateID() string {
	return f()
}
