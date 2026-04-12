package mrmodel

import (
	"io"
	"mime/multipart"
)

type (
	// File - файл с мета-информацией и потоком данных для чтения.
	// Body представляет io.ReadCloser для чтения содержимого файла.
	File struct {
		FileInfo

		Body io.ReadCloser
	}

	// FileContent - файл с мета-информацией и содержимым в памяти.
	// Body содержит все данные файла в виде байтового слайса.
	FileContent struct {
		FileInfo

		Body []byte
	}

	// FileHeader - файл с мета-информацией из multipart/form-data запроса.
	// Header содержит заголовок multipart-файла для доступа к оригинальным метаданным.
	FileHeader struct {
		FileInfo

		Header *multipart.FileHeader
	}
)
