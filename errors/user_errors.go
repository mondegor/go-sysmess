package errors

var (
	// ErrIncorrectInputData - входные данные некорректны (желательно настраивать в слое выше валидатор этих данных).
	ErrIncorrectInputData = NewUserProto("IncorrectInputData", "input data is incorrect: '{Reason}'")

	// ErrAccessForbidden - доступ запрещён (код 403).
	ErrAccessForbidden = NewUserError("AccessForbidden", "access forbidden")

	// ErrRecordNotFound - запись не найдена (код 404).
	ErrRecordNotFound = NewUserError("RecordNotFound", "record not found")

	// ErrRecordAlreadyExists - запись уже существует.
	ErrRecordAlreadyExists = NewUserError("RecordAlreadyExists", "record already exists")

	// ErrRecordVersionConflict - конфликт версий записи (код 409).
	ErrRecordVersionConflict = NewUserError("RecordVersionConflict", "record version conflict")

	// ErrSwitchStatusRejected - переход между статусами отклонён.
	ErrSwitchStatusRejected = NewUserProto("SwitchStatusRejected", "switching from '{StatusFrom}' to '{StatusTo}' is rejected")

	// ErrNotImplemented - функциональность не реализована (код 501).
	ErrNotImplemented = NewUserError("NotImplemented", "not implemented")

	// ErrValidateInvalidFile - файл некорректен.
	ErrValidateInvalidFile = NewUserError("ValidateInvalidFile", "file is invalid")

	// ErrValidateFileSize - некорректный размер файла.
	ErrValidateFileSize = NewUserError("ValidateFileSize", "invalid file size")

	// ErrValidateFileSizeMin - некорректный размер файла - минимальный.
	ErrValidateFileSizeMin = NewUserProto("ValidateFileSizeMin", "invalid file size, min size = {Value}b")

	// ErrValidateFileSizeMax - некорректный размер файла - максимальный.
	ErrValidateFileSizeMax = NewUserProto("ValidateFileSizeMax", "invalid file size, max size = {Value}b")

	// ErrValidateFileExtension - некорректное расширение файла.
	ErrValidateFileExtension = NewUserProto("ValidateFileExtension", "invalid file extension: {Value}")

	// ErrValidateFileTotalSizeMax - некорректный общий размер файлов - максимальный.
	ErrValidateFileTotalSizeMax = NewUserProto("ValidateFileTotalSizeMax", "invalid file total size, max total size = {Value}b")

	// ErrValidateFileContentType - тип содержимого не совпадает с определённым типом.
	ErrValidateFileContentType = NewUserProto("ValidateFileContentType", "the content type '{Value}' does not match the detected type")

	// ErrValidateFileUnsupportedType - неподдерживаемый тип файла.
	ErrValidateFileUnsupportedType = NewUserProto("ValidateFileUnsupportedType", "unsupported file type '{Value}'")

	// ErrValidateImageSize - некорректный размер изображения (ширина, высота).
	ErrValidateImageSize = NewUserError("ValidateImageSize", "invalid image size (width, height)")

	// ErrValidateImageWidthMax - некорректная ширина изображения - максимальная.
	ErrValidateImageWidthMax = NewUserProto("ValidateImageWidthMax", "invalid image width, max size = {Value}px")

	// ErrValidateImageHeightMax - некорректная высота изображения - максимальная.
	ErrValidateImageHeightMax = NewUserProto("ValidateImageHeightMax", "invalid image height, max size = {Value}px")

	// ErrHttpFileUpload - файл с указанным ключом не был загружен.
	ErrHttpFileUpload = NewUserProto("FileUpload", "the file with the specified key '{Key}' was not uploaded")

	// ErrHttpClientUnauthorized - 401. клиент не авторизован.
	ErrHttpClientUnauthorized = NewUserError("ClientUnauthorized", "401. client is unauthorized")

	// ErrHttpAccessForbidden - 403. доступ запрещён.
	ErrHttpAccessForbidden = NewUserError("AccessForbidden", "403. access forbidden")

	// ErrHttpResourceNotFound - 404. ресурс не найден.
	ErrHttpResourceNotFound = NewUserError("ResourceNotFound", "404. resource not found")

	// ErrHttpRequestParseData - тело запроса некорректно (ошибка связанная с неправильным форматом отправленных данных, код 422).
	ErrHttpRequestParseData = NewUserProto("RequestParseData", "request body is not valid: '{Reason}'")

	// ErrHttpTooManyRequests - слишком много запросов (код 429).
	ErrHttpTooManyRequests = NewUserError("TooManyRequests", "too many requests")
)
