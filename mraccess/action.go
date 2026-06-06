package mraccess

type (
	// Action - обработчик с наделёнными ему правами доступа.
	Action struct {
		// Name - уникальное имя обработчика.
		Name string

		// Privilege - привилегия, необходимая для вызова обработчика.
		Privilege string

		// Permission - дополнительное разрешение для тонкой настройки доступа.
		Permission string
	}

	// ActionGroup - группа (секция) объединяющая несколько обработчиков.
	// Включает базовый путь к этим обработчикам и обладает привилегией доступа к ним.
	ActionGroup struct {
		// Name - уникальное имя группы действий.
		Name string

		// Privilege - привилегия, необходимая для доступа ко всем действиям в группе.
		Privilege string

		// BasePath - базовый путь для всех обработчиков в группе.
		BasePath string
	}
)
