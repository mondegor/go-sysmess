package errors

import (
	"github.com/mondegor/go-sysmess/errors/custom"
	"github.com/mondegor/go-sysmess/errors/helper"
)

type (
	// CustomWrapper - помощник для оборачивания пользовательских ошибок.
	CustomWrapper interface {
		Wrap(err error) error
	}
)

// NewCustomWrapper - оборачивает ошибки с указанными кодами.
// Коды передаются попарно: 1 - код пользовательской ошибки, 2 - переопределённый код.
func NewCustomWrapper(codeCustom ...string) CustomWrapper {
	return helper.NewCustomErrorWrapper(
		func(err error, customCode string) error {
			return custom.New(err, customCode)
		},
		codeCustom...,
	)
}

// NewDownloadFileWrapper - оборачивает ошибки связанные с парсингом файла.
func NewDownloadFileWrapper(customCode string) CustomWrapper {
	return NewCustomWrapper(
		ErrValidateFileSizeMin.Code(), customCode,
		ErrValidateFileSizeMax.Code(), customCode,
		ErrValidateFileExtension.Code(), customCode,
		ErrValidateFileContentType.Code(), customCode,
		ErrValidateFileUnsupportedType.Code(), customCode,
	)
}

// NewDownloadImageWrapper - оборачивает ошибки связанные с парсингом файла изображения.
func NewDownloadImageWrapper(customCode string) CustomWrapper {
	return NewCustomWrapper(
		ErrValidateFileSizeMin.Code(), customCode,
		ErrValidateFileSizeMax.Code(), customCode,
		ErrValidateFileExtension.Code(), customCode,
		ErrValidateFileContentType.Code(), customCode,
		ErrValidateFileUnsupportedType.Code(), customCode,
		ErrValidateImageWidthMax.Code(), customCode,
		ErrValidateImageHeightMax.Code(), customCode,
	)
}
