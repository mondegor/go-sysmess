package mraccess

const (
	// PrivilegePublic - привилегия для всех.
	PrivilegePublic = "public"

	// PermissionAnyUser - разрешение для любого пользователя.
	PermissionAnyUser = "any-user"

	// PermissionGuestOnly - разрешение только для гостя.
	PermissionGuestOnly = "guest-only"
)

type (
	// RightsChecker - проверяет, обладает ли сущность (пользователь, группа ролей и т.д.)
	// указанным правом. Используется в момент запроса (request-время).
	RightsChecker interface {
		Has(name string) bool
	}

	// RightsRegistry - проверяет, зарегистрировано ли указанное право (привилегия
	// или разрешение) в системе. Используется при инициализации (init-время).
	// Сигнатура намеренно отличается от RightsChecker: это разные вопросы
	// («обладает ли» vs «зарегистрировано ли»), и разные имена методов не дают
	// случайно подменить один контракт другим.
	RightsRegistry interface {
		IsRegistered(name string) bool
	}

	// RightsGetter - возвращает объект проверки прав для указанной группы ролей.
	RightsGetter interface {
		// Rights - возвращает RightsChecker для указанной группы.
		Rights(group string) RightsChecker
	}
)
