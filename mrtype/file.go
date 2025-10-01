package mrtype

import (
	"io"
	"mime/multipart"
)

type (
	// File - мета-информация файла вместе с источником файла.
	File struct {
		FileInfo
		Body io.ReadCloser
	}

	// FileContent - файл с мета-информацией.
	FileContent struct {
		FileInfo
		Body []byte
	}

	// FileHeader - мета-информация файла вместе с источником файла (multipart/form-data).
	FileHeader struct {
		FileInfo
		Header *multipart.FileHeader
	}
)
