package errors

var (
	// ErrSystemTimeoutPeriodHasExpired - the timeout period has expired.
	ErrSystemTimeoutPeriodHasExpired = NewSystemProto("the timeout period has expired")

	// ErrSystemStorageConnectionFailed - connection failed (attr: source).
	ErrSystemStorageConnectionFailed = NewSystemProto("connection failed")

	// ErrSystemStorageUnexpectedEOF - unexpected EOF (attr: source).
	ErrSystemStorageUnexpectedEOF = NewSystemProto("unexpected EOF")

	// ErrSystemStorageFailedToClose - failed to close source (attr: source).
	ErrSystemStorageFailedToClose = NewSystemProto("failed to close source")

	// ErrSystemServiceTemporarilyUnavailable - service is temporarily unavailable.
	// Системная ошибка, которая сообщает о сетевых проблемах, о работоспособности внешних ресурсов (БД, API, FileSystem).
	ErrSystemServiceTemporarilyUnavailable = NewSystemProto("service is temporarily unavailable")

	// ErrSystemHttpMultipartFormFile - the file with the specified key cannot be processed (attr: key).
	ErrSystemHttpMultipartFormFile = NewSystemProto("the file with the specified key cannot be processed")
)
