package mr

import (
	"github.com/mondegor/go-sysmess/mrerr"
)

var (
	// ErrValidateFileSize - invalid file size.
	ErrValidateFileSize = mrerr.NewKindUser("ValidateFileSize", "invalid file size")

	// ErrValidateFileSizeMin - invalid file size - min.
	ErrValidateFileSizeMin = mrerr.NewKindUser("ValidateFileSizeMin", "invalid file size, min size = {Value}b")

	// ErrValidateFileSizeMax - invalid file size - max.
	ErrValidateFileSizeMax = mrerr.NewKindUser("ValidateFileSizeMax", "invalid file size, max size = {Value}b")

	// ErrValidateFileExtension - invalid file extension.
	ErrValidateFileExtension = mrerr.NewKindUser("ValidateFileExtension", "invalid file extension: {Value}")

	// ErrValidateFileTotalSizeMax - invalid file total size - max.
	ErrValidateFileTotalSizeMax = mrerr.NewKindUser("ValidateFileTotalSizeMax", "invalid file total size, max total size = {Value}b")

	// ErrValidateFileContentType - the content type does not match the detected type.
	ErrValidateFileContentType = mrerr.NewKindUser("ValidateFileContentType", "the content type '{Value}' does not match the detected type")

	// ErrValidateFileUnsupportedType - unsupported file type.
	ErrValidateFileUnsupportedType = mrerr.NewKindUser("ValidateFileUnsupportedType", "unsupported file type '{Value}'")

	// ErrValidateImageSize - invalid image size (width, height).
	ErrValidateImageSize = mrerr.NewKindUser("ValidateImageSize", "invalid image size (width, height)")

	// ErrValidateImageWidthMax - invalid image width - max.
	ErrValidateImageWidthMax = mrerr.NewKindUser("ValidateImageWidthMax", "invalid image width, max size = {Value}px")

	// ErrValidateImageHeightMax - invalid image height - max.
	ErrValidateImageHeightMax = mrerr.NewKindUser("ValidateImageHeightMax", "invalid image height, max size = {Value}px")
)
