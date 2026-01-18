package errors

var (
	// ErrInternalNilPointer - unexpected nil pointer.
	ErrInternalNilPointer = NewInternalProto("unexpected nil pointer")

	// ErrInternalCaughtPanic - caught panic (attrs: source, recover, stack_trace).
	ErrInternalCaughtPanic = NewInternalProto("internal panic")

	// ErrInternalTypeAssertion - invalid type assertion (attrs: type, value).
	ErrInternalTypeAssertion = NewInternalProto("invalid type assertion")

	// ErrInternalInvalidType - invalid type, expected type (attrs: type, expected).
	ErrInternalInvalidType = NewInternalProto("invalid type, expected type")

	// ErrInternalKeyNotFoundInSource - key is not found in source (attrs: key, source).
	ErrInternalKeyNotFoundInSource = NewInternalProto("key is not found in source")

	// ErrInternalIncorrectInputData - input data is incorrect
	// (когда нет возможности ответить клиенту, например, данные поступают из очереди,
	// или когда поступаемые данные не зависят от клиента, например, приходят из конфига,
	// в остальных случаях лучше использовать ErrUseCaseIncorrectInputData).
	ErrInternalIncorrectInputData = NewInternalProto("input data is incorrect")

	// ErrInternalStorageConnectionIsAlreadyCreated - connection is already created (attr: source).
	ErrInternalStorageConnectionIsAlreadyCreated = NewInternalProto("connection is already created")

	// ErrInternalStorageConnectionIsNotOpened - connection is not opened (attr: source).
	ErrInternalStorageConnectionIsNotOpened = NewInternalProto("connection is not opened")

	// ErrInternalStorageQueryFailed - query failed (attr: source).
	// По умолчанию, следует оборачивать все ошибки запросов этой ошибкой.
	ErrInternalStorageQueryFailed = NewInternalProto("query failed")

	// ErrInternalStorageFetchDataFailed - fetching data failed (attr: source).
	ErrInternalStorageFetchDataFailed = NewInternalProto("fetching data failed")

	// ErrInternalServiceOperationFailed - service operation failed.
	// По умолчанию, следует оборачивать все ошибки Service слоя этой ошибкой.
	ErrInternalServiceOperationFailed = NewInternalProto("service operation failed")

	// ErrInternalUseCaseOperationFailed - operation failed.
	// По умолчанию, следует оборачивать все ошибки UseCase слоя этой ошибкой.
	ErrInternalUseCaseOperationFailed = NewInternalProto("operation failed")

	// ErrInternalHttpResponseParseData - response data is not valid.
	ErrInternalHttpResponseParseData = NewInternalProto("response data is not valid")
)
