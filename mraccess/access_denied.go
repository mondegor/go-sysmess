package mraccess

type (
	// accessDenied - заглушка, используемая когда доступ явно запрещён.
	// Применяется как объект по умолчанию для неизвестных групп ролей.
	accessDenied struct{}
)

// Has - всегда сообщает, что право отсутствует.
func (accessDenied) Has(_ string) bool {
	return false
}
