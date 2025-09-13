package mr

const (
	// DefaultErrorCodeInternal - обобщённый код ошибки: внутренняя ошибка приложения.
	DefaultErrorCodeInternal = "Internal"

	// DefaultErrorCodeSystem - обобщённый код ошибки: системная ошибка приложения.
	DefaultErrorCodeSystem = "System"

	// ErrorCodeUnexpectedInternal - внутренняя ошибка, содержащая ошибку, которую следовало обернуть ранее.
	ErrorCodeUnexpectedInternal = "UnexpectedInternal"

	// ErrorCodeTemporarilyUnavailable - системная ошибка, временной недоступности какого-либо ресурса.
	ErrorCodeTemporarilyUnavailable = "TemporarilyUnavailable"
)
