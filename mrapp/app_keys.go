package mrapp

// Названия ключей атрибутов, под которыми приложение (сервис) пишет
// свои сведения в лог. Держатся в одном месте, чтобы имя ключа задавалось
// один раз, а не повторялось строковым литералом в каждом месте записи:
// по этим же именам настраивается раскраска атрибутов (см. wire/mrlog/slog).
const (
	// KeyAppEnvironment - название ключа окружения приложения (сервиса).
	KeyAppEnvironment = "app_env"

	// KeyAppVersion - название ключа версии приложения (сервиса).
	KeyAppVersion = "app_ver"

	// KeyErrorID - название ключа ID ошибки.
	KeyErrorID = "error_id"

	// KeyStackTrace - название ключа стека вызовов формируемого при ошибке.
	KeyStackTrace = "stack_trace"

	// KeyLangCode - название ключа кода языка клиента (пользователя).
	KeyLangCode = "lang"

	// KeyUserID - название ключа ID пользователя.
	// KeyUserID = "user_id".
)
