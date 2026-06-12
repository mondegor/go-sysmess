package mraccess

type (
	// User - представляет пользователя с привязанными к нему
	// привилегиями и разрешениями.
	User interface {
		// RightsChecker - встраивает методы проверки прав доступа.
		RightsChecker

		// ID - возвращает уникальный идентификатор пользователя (UUID v4).
		ID() [16]byte

		// Group - возвращает имя группы ролей пользователя.
		Group() string

		// LangCode - возвращает код языка интерфейса пользователя.
		LangCode() string
	}

	// entryUser - внутренняя реализация интерфейса User.
	entryUser struct {
		id       [16]byte
		group    string
		langCode string
		rights   RightsChecker
	}
)

// NewUser - создаёт объект User с указанными параметрами.
// Права доступа определяются через RightsGetter для указанной группы.
func NewUser(id [16]byte, group, langCode string, rights RightsGetter) User {
	return &entryUser{
		id:       id,
		group:    group,
		langCode: langCode,
		rights:   rights.Rights(group),
	}
}

// ID - возвращает уникальный идентификатор пользователя.
func (u *entryUser) ID() [16]byte {
	return u.id
}

// Group - возвращает имя группы ролей пользователя.
func (u *entryUser) Group() string {
	return u.group
}

// LangCode - возвращает код языка интерфейса пользователя.
func (u *entryUser) LangCode() string {
	return u.langCode
}

// Has - сообщает о наличии указанного права у пользователя.
func (u *entryUser) Has(name string) bool {
	return u.rights.Has(name)
}
