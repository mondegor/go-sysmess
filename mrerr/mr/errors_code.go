package mr

const (
	DefaultErrorCodeInternal = "Internal" // DefaultErrorCodeInternal - обобщённый код ошибки: внутренняя ошибка приложения.
	DefaultErrorCodeSystem   = "System"   // DefaultErrorCodeSystem - обобщённый код ошибки: системная ошибка приложения.

	// ErrorCodeUnexpectedInternal - внутренняя ошибка, содержащая ошибку, которую следовало обернуть ранее.
	ErrorCodeUnexpectedInternal = "UnexpectedInternal"

	// ErrorCodeTemporarilyUnavailable - системная ошибка, временной недоступности какого-либо ресурса.
	ErrorCodeTemporarilyUnavailable = "TemporarilyUnavailable"
)
