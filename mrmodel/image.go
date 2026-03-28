package mrmodel

import (
	"io"
	"mime/multipart"
)

type (
	// Image - мета-информация изображения вместе с источником изображения.
	Image struct {
		ImageInfo
		Body io.ReadCloser
	}

	// ImageContent - изображение с мета-информацией.
	ImageContent struct {
		ImageInfo
		Body []byte
	}

	// ImageHeader - мета-информация изображения вместе с источником изображения (multipart/form-data).
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
