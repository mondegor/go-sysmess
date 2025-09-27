package mrtype

import (
	"io"
	"mime/multipart"

	"github.com/mondegor/go-sysmess/mrdto"
)

type (
	// File - мета-информация файла вместе с источником файла.
	File struct {
		mrdto.FileInfo
		Body io.ReadCloser
	}

	// FileContent - файл с мета-информацией.
	FileContent struct {
		mrdto.FileInfo
		Body []byte
	}

	// FileHeader - мета-информация файла вместе с источником файла (multipart/form-data).
	FileHeader struct {
		mrdto.FileInfo
		Header *multipart.FileHeader
	}
)
