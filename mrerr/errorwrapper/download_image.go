package errorwrapper

import (
	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/mr"
)

type (
	// DownloadUserImage - помощник для оборачивания пользовательских ошибок связанных с загрузкой файла изображения.
	DownloadUserImage struct {
		file *DownloadUserFile
	}
)

// NewDownloadUserImage - создаёт объект DownloadUserImage.
func NewDownloadUserImage() *DownloadUserImage {
	return &DownloadUserImage{
		file: &DownloadUserFile{},
	}
}

// WrapError - оборачивает ошибки связанные с парсингом файла изображения.
func (w *DownloadUserImage) WrapError(err error, name string) error {
	if mr.ErrValidateImageWidthMax.Is(err) { // вложенные ошибки не учитываются
		return mrerr.NewCustomError(name, err)
	}

	if mr.ErrValidateImageHeightMax.Is(err) { // вложенные ошибки не учитываются
		return mrerr.NewCustomError(name, err)
	}

	return w.file.WrapError(err, name)
}
