package mrtype

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
func (i *Image) ToFile() File {
	return File{
		FileInfo: i.ImageInfo.ToFileInfo(),
		Body:     i.Body,
	}
}

// ToFileContent - возвращает изображение преобразованное в файловую структуру
// (с потерей дополнительной информации об изображении).
func (i *ImageContent) ToFileContent() FileContent {
	return FileContent{
		FileInfo: i.ImageInfo.ToFileInfo(),
		Body:     i.Body,
	}
}

// ToFileHeader - возвращает изображение преобразованное в файловую структуру
// (с потерей дополнительной информации об изображении).
func (i *ImageHeader) ToFileHeader() FileHeader {
	return FileHeader{
		FileInfo: i.ImageInfo.ToFileInfo(),
		Header:   i.Header,
	}
}
