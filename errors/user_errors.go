package errors

var (
	// ErrUseCaseIncorrectInputData - input data is incorrect (желательно настраивать в слое выше валидатор этих данных).
	ErrUseCaseIncorrectInputData = NewUserProto("IncorrectInputData", "input data is incorrect: '{Reason}'")

	// ErrUseCaseAccessForbidden - access forbidden (код 401 или 403).
	ErrUseCaseAccessForbidden = NewUserProto("AccessForbidden", "access forbidden")

	// ErrUseCaseEntityNotFound - entity not found (код 404).
	ErrUseCaseEntityNotFound = NewUserProto("EntityNotFound", "entity not found")

	// ErrUseCaseEntityAlreadyExists - entity already exists.
	ErrUseCaseEntityAlreadyExists = NewUserProto("EntityAlreadyExists", "entity already exists")

	// ErrUseCaseEntityVersionConflict - entity version conflict (код 409).
	ErrUseCaseEntityVersionConflict = NewUserProto("EntityVersionConflict", "entity version conflict")

	// ErrUseCaseSwitchStatusRejected - switching from status to status is rejected.
	ErrUseCaseSwitchStatusRejected = NewUserProto("SwitchStatusRejected", "switching from '{StatusFrom}' to '{StatusTo}' is rejected")

	// ErrUseCaseInvalidFile - file is invalid.
	ErrUseCaseInvalidFile = NewUserProto("InvalidFile", "file is invalid")

	// ErrValidateFileSize - invalid file size.
	ErrValidateFileSize = NewUserProto("ValidateFileSize", "invalid file size")

	// ErrValidateFileSizeMin - invalid file size - min.
	ErrValidateFileSizeMin = NewUserProto("ValidateFileSizeMin", "invalid file size, min size = {Value}b")

	// ErrValidateFileSizeMax - invalid file size - max.
	ErrValidateFileSizeMax = NewUserProto("ValidateFileSizeMax", "invalid file size, max size = {Value}b")

	// ErrValidateFileExtension - invalid file extension.
	ErrValidateFileExtension = NewUserProto("ValidateFileExtension", "invalid file extension: {Value}")

	// ErrValidateFileTotalSizeMax - invalid file total size - max.
	ErrValidateFileTotalSizeMax = NewUserProto("ValidateFileTotalSizeMax", "invalid file total size, max total size = {Value}b")

	// ErrValidateFileContentType - the content type does not match the detected type.
	ErrValidateFileContentType = NewUserProto("ValidateFileContentType", "the content type '{Value}' does not match the detected type")

	// ErrValidateFileUnsupportedType - unsupported file type.
	ErrValidateFileUnsupportedType = NewUserProto("ValidateFileUnsupportedType", "unsupported file type '{Value}'")

	// ErrValidateImageSize - invalid image size (width, height).
	ErrValidateImageSize = NewUserProto("ValidateImageSize", "invalid image size (width, height)")

	// ErrValidateImageWidthMax - invalid image width - max.
	ErrValidateImageWidthMax = NewUserProto("ValidateImageWidthMax", "invalid image width, max size = {Value}px")

	// ErrValidateImageHeightMax - invalid image height - max.
	ErrValidateImageHeightMax = NewUserProto("ValidateImageHeightMax", "invalid image height, max size = {Value}px")

	// ErrHttpFileUpload - the file with the specified key was not uploaded.
	ErrHttpFileUpload = NewUserProto("FileUpload", "the file with the specified key '{Key}' was not uploaded")

	// ErrHttpClientUnauthorized - 401. client is unauthorized.
	ErrHttpClientUnauthorized = NewUserProto("ClientUnauthorized", "401. client is unauthorized")

	// ErrHttpAccessForbidden - 403. access forbidden.
	ErrHttpAccessForbidden = NewUserProto("AccessForbidden", "403. access forbidden")

	// ErrHttpResourceNotFound - 404. resource not found.
	ErrHttpResourceNotFound = NewUserProto("ResourceNotFound", "404. resource not found")

	// ErrHttpResourceVersionInvalid - resource version is invalid (код 409).
	ErrHttpResourceVersionInvalid = NewUserProto("ResourceVersionInvalid", "resource version is invalid")

	// ErrHttpRequestParseData - request body is not valid (ошибка связанная с неправильным форматом отправленных данных, код 422).
	ErrHttpRequestParseData = NewUserProto("RequestParseData", "request body is not valid: '{Reason}'")

	// ErrHttpTooManyRequests - too many requests (код 429).
	ErrHttpTooManyRequests = NewUserProto("TooManyRequests", "too many requests")
)
