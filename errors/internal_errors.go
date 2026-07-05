package errors

var (
	// ErrInternalNilPointer - неожиданный нулевой указатель.
	ErrInternalNilPointer = NewInternalProto("unexpected nil pointer")

	// ErrInternalCaughtPanic - перехвачена паника (attrs: source, recover, stack_trace).
	ErrInternalCaughtPanic = NewInternalProto("internal panic")

	// ErrInternalTypeAssertion - некорректное приведение типа (attrs: type, value).
	ErrInternalTypeAssertion = NewInternalProto("invalid type assertion")

	// ErrInternalInvalidType - некорректный тип, ожидался другой тип (attrs: type, expected).
	ErrInternalInvalidType = NewInternalProto("invalid type, expected type")

	// ErrInternalKeyNotFoundInSource - ключ не найден в источнике (attrs: key, source).
	ErrInternalKeyNotFoundInSource = NewInternalProto("key is not found in source")

	// ErrInternalIncorrectInputData - входные данные некорректны.
	// Используется, когда нет возможности ответить клиенту, например, данные поступают из очереди,
	// или когда поступающие данные не зависят от клиента, например, приходят из конфига.
	// В остальных случаях лучше использовать ErrIncorrectInputData.
	ErrInternalIncorrectInputData = NewInternalProto("input data is incorrect")

	// ErrInternalStorageConnectionIsAlreadyCreated - подключение уже создано (attr: source).
	ErrInternalStorageConnectionIsAlreadyCreated = NewInternalProto("connection is already created")

	// ErrInternalStorageConnectionIsNotOpened - подключение не открыто (attr: source).
	ErrInternalStorageConnectionIsNotOpened = NewInternalProto("connection is not opened")

	// ErrInternalStorageQueryFailed - запрос не удался (attr: source).
	// По умолчанию, следует оборачивать все ошибки запросов этой ошибкой.
	ErrInternalStorageQueryFailed = NewInternalProto("query failed")

	// ErrInternalStorageFetchDataFailed - получение данных не удалось (attr: source).
	ErrInternalStorageFetchDataFailed = NewInternalProto("fetching data failed")

	// ErrInternalStorageDuplicateKeyViolation - нарушение уникальности ключа (attr: source).
	ErrInternalStorageDuplicateKeyViolation = NewInternalProto("duplicate key violation")

	// ErrInternalServiceOperationFailed - операция сервиса не удалась.
	// По умолчанию, следует оборачивать все ошибки Service/UseCase слоя этой ошибкой.
	ErrInternalServiceOperationFailed = NewInternalProto("service operation failed")

	// ErrInternalHttpResponseParseData - ответ некорректен.
	ErrInternalHttpResponseParseData = NewInternalProto("response data is not valid")
)
