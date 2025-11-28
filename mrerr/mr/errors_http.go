package mr

import "github.com/mondegor/go-sysmess/mrerr"

var (
	// ErrHttpResponseParseData - response data is not valid.
	ErrHttpResponseParseData = mrerr.NewKindInternal("response data is not valid")

	// ErrHttpFileUpload - the file with the specified key was not uploaded.
	ErrHttpFileUpload = mrerr.NewKindUser("FileUpload", "the file with the specified key '{Key}' was not uploaded")

	// ErrHttpMultipartFormFile - the file with the specified key cannot be processed.
	ErrHttpMultipartFormFile = mrerr.NewKindSystem("the file with the specified key cannot be processed: '{Key}'")

	// ErrHttpClientUnauthorized - 401. client is unauthorized.
	ErrHttpClientUnauthorized = mrerr.NewKindUser("ClientUnauthorized", "401. client is unauthorized")

	// ErrHttpAccessForbidden - 403. access forbidden.
	ErrHttpAccessForbidden = mrerr.NewKindUser("AccessForbidden", "403. access forbidden")

	// ErrHttpResourceNotFound - 404. resource not found.
	ErrHttpResourceNotFound = mrerr.NewKindUser("ResourceNotFound", "404. resource not found")

	// ErrHttpRequestParseData - 422. request body is not valid (ошибка связанная с неправильным форматом отправленных данных).
	ErrHttpRequestParseData = mrerr.NewKindUser("RequestParseData", "request body is not valid: '{Reason}'")
)
