package mrpath

type (
	// Builder - интерфейс для построения полных путей из базового шаблона и переданного значения.
	// Реализации могут вставлять значение в середину пути (placeholder)
	// или добавлять его в конец (tail).
	Builder interface {
		BuildPath(path string) string
	}
)
