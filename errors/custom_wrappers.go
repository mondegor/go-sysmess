package errors

import (
	"github.com/mondegor/go-core/errors/custom"
	"github.com/mondegor/go-core/errors/wrap"
)

type (
	// CustomWrapper - помощник для оборачивания пользовательских ошибок в ошибки типа Custom.
	CustomWrapper = wrap.CustomErrorWrapper
)

// NewCustomWrapper - создаёт обёртку с маппингом кодов пользовательских ошибок.
// Параметр codeCustom - плоский слайс пар "исходныйКод/кастомныйКод".
// Пример: NewCustomWrapper("userErr1", "fieldEmail").
func NewCustomWrapper(codeCustom ...string) CustomWrapper {
	return wrap.NewCustomErrorWrapper(
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

// NopCustomWrapper - создаёт CustomWrapper-заглушку.
func NopCustomWrapper() CustomWrapper {
	return wrap.NopCustomErrorWrapper()
}
