package errors

var (
	// ErrSystemTimeoutPeriodHasExpired - период ожидания процесса/соединения истёк.
	ErrSystemTimeoutPeriodHasExpired = NewSystemProto("the timeout period has expired")

	// ErrSystemStorageConnectionFailed - подключение к хранилищу не удалось (attr: source).
	ErrSystemStorageConnectionFailed = NewSystemProto("connection failed")

	// ErrSystemStorageUnexpectedEOF - неожиданный конец файла (attr: source).
	ErrSystemStorageUnexpectedEOF = NewSystemProto("unexpected EOF")

	// ErrSystemStorageFailedToClose - не удалось закрыть источник (attr: source).
	ErrSystemStorageFailedToClose = NewSystemProto("failed to close source")

	// ErrSystemServiceTemporarilyUnavailable - сервис временно недоступен.
	// Системная ошибка, которая сообщает о сетевых проблемах, о работоспособности внешних ресурсов (БД, API, FileSystem).
	ErrSystemServiceTemporarilyUnavailable = NewSystemProto("service is temporarily unavailable")

	// ErrSystemHttpMultipartFormFile - файл с указанным ключом не может быть обработан (attr: key).
	ErrSystemHttpMultipartFormFile = NewSystemProto("the file with the specified key cannot be processed")
)
