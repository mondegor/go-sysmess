package errors

var (
	// ErrIncorrectInputData - input data is incorrect (желательно настраивать в слое выше валидатор этих данных).
	ErrIncorrectInputData = NewUserProto("IncorrectInputData", "input data is incorrect: '{Reason}'")

	// ErrAccessForbidden - access forbidden (код 403).
	ErrAccessForbidden = NewUserError("AccessForbidden", "access forbidden")

	// ErrRecordNotFound - record not found (код 404).
	ErrRecordNotFound = NewUserError("RecordNotFound", "record not found")

	// ErrRecordAlreadyExists - record already exists.
	ErrRecordAlreadyExists = NewUserError("RecordAlreadyExists", "record already exists")

	// ErrRecordVersionConflict - record version conflict (код 409).
	ErrRecordVersionConflict = NewUserError("RecordVersionConflict", "record version conflict")

	// ErrSwitchStatusRejected - switching from status to status is rejected.
	ErrSwitchStatusRejected = NewUserProto("SwitchStatusRejected", "switching from '{StatusFrom}' to '{StatusTo}' is rejected")

	// ErrNotImplemented - not Implemented (код 501).
	ErrNotImplemented = NewUserError("NotImplemented", "not implemented")

	// ErrValidateInvalidFile - file is invalid.
	ErrValidateInvalidFile = NewUserError("ValidateInvalidFile", "file is invalid")

	// ErrValidateFileSize - invalid file size.
	ErrValidateFileSize = NewUserError("ValidateFileSize", "invalid file size")

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
	ErrValidateImageSize = NewUserError("ValidateImageSize", "invalid image size (width, height)")

	// ErrValidateImageWidthMax - invalid image width - max.
	ErrValidateImageWidthMax = NewUserProto("ValidateImageWidthMax", "invalid image width, max size = {Value}px")

	// ErrValidateImageHeightMax - invalid image height - max.
	ErrValidateImageHeightMax = NewUserProto("ValidateImageHeightMax", "invalid image height, max size = {Value}px")

	// ErrHttpFileUpload - the file with the specified key was not uploaded.
	ErrHttpFileUpload = NewUserProto("FileUpload", "the file with the specified key '{Key}' was not uploaded")

	// ErrHttpClientUnauthorized - 401. client is unauthorized.
	ErrHttpClientUnauthorized = NewUserError("ClientUnauthorized", "401. client is unauthorized")

	// ErrHttpAccessForbidden - 403. access forbidden.
	ErrHttpAccessForbidden = NewUserError("AccessForbidden", "403. access forbidden")

	// ErrHttpResourceNotFound - 404. resource not found.
	ErrHttpResourceNotFound = NewUserError("ResourceNotFound", "404. resource not found")

	// ErrHttpRequestParseData - request body is not valid (ошибка связанная с неправильным форматом отправленных данных, код 422).
	ErrHttpRequestParseData = NewUserProto("RequestParseData", "request body is not valid: '{Reason}'")

	// ErrHttpTooManyRequests - too many requests (код 429).
	ErrHttpTooManyRequests = NewUserError("TooManyRequests", "too many requests")
)
