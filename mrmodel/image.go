package mrmodel

import (
	"io"
	"mime/multipart"
)

type (
	// Image - изображение с мета-информацией и потоком данных для чтения.
	// Body представляет io.ReadCloser для чтения содержимого изображения.
	Image struct {
		ImageInfo
		Body io.ReadCloser
	}

	// ImageContent - изображение с мета-информацией и содержимым в памяти.
	// Body содержит все данные изображения в виде байтового слайса.
	ImageContent struct {
		ImageInfo
		Body []byte
	}

	// ImageHeader - изображение с мета-информацией из multipart/form-data запроса.
	// Header содержит заголовок multipart-файла для доступа к оригинальным метаданным.
	ImageHeader struct {
		ImageInfo
		Header *multipart.FileHeader
	}
)

// ToFile - возвращает изображение преобразованное в файловую структуру
// (с потерей дополнительной информации об изображении).
func (im Image) ToFile() File {
	return File{
		FileInfo: im.ImageInfo.ToFileInfo(),
		Body:     im.Body,
	}
}

// ToFileContent - возвращает изображение преобразованное в файловую структуру
// (с потерей дополнительной информации об изображении).
func (im ImageContent) ToFileContent() FileContent {
	return FileContent{
		FileInfo: im.ImageInfo.ToFileInfo(),
		Body:     im.Body,
	}
}

// ToFileHeader - возвращает изображение преобразованное в файловую структуру
// (с потерей дополнительной информации об изображении).
func (im ImageHeader) ToFileHeader() FileHeader {
	return FileHeader{
		FileInfo: im.ImageInfo.ToFileInfo(),
		Header:   im.Header,
	}
}
