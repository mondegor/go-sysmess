package mrpath

type (
	// Builder - формирует полный путь на основе базового и указанного.
	Builder interface {
		BuildPath(path string) string
	}
)
