package mrtype

import (
	"io"
	"mime/multipart"

	"github.com/mondegor/go-sysmess/mrdto"
)

type (
	// Image - мета-информация изображения вместе с источником изображения.
	Image struct {
		mrdto.ImageInfo
		Body io.ReadCloser
	}

	// ImageContent - изображение с мета-информацией.
	ImageContent struct {
		mrdto.ImageInfo
		Body []byte
	}

	// ImageHeader - мета-информация изображения вместе с источником изображения (multipart/form-data).
	ImageHeader struct {
		mrdto.ImageInfo
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
