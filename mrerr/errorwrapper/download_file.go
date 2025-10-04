package errorwrapper

import (
	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/mr"
)

type (
	// DownloadUserFile - помощник для оборачивания пользовательских ошибок связанных с загрузкой файла.
	DownloadUserFile struct{}
)

// NewDownloadUserFile - создаёт объект DownloadUserFile.
func NewDownloadUserFile() *DownloadUserFile {
	return &DownloadUserFile{}
}

// WrapError - оборачивает ошибки связанные с парсингом файла.
func (w *DownloadUserFile) WrapError(err error, name string) error {
	if mr.ErrValidateFileSizeMin.Is(err) { // вложенные ошибки не учитываются
		return mrerr.NewCustomError(name, err)
	}

	if mr.ErrValidateFileSizeMax.Is(err) { // вложенные ошибки не учитываются
		return mrerr.NewCustomError(name, err)
	}

	if mr.ErrValidateFileExtension.Is(err) { // вложенные ошибки не учитываются
		return mrerr.NewCustomError(name, err)
	}

	if mr.ErrValidateFileContentType.Is(err) { // вложенные ошибки не учитываются
		return mrerr.NewCustomError(name, err)
	}

	if mr.ErrValidateFileUnsupportedType.Is(err) { // вложенные ошибки не учитываются
		return mrerr.NewCustomError(name, err)
	}

	return err
}
