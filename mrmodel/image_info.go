package mrmodel

import (
	"path"
	"time"

	"github.com/mondegor/go-sysmess/util/copyptr"
)

type (
	// ImageInfo - мета-информация об изображении.
	// Расширяет FileInfo дополнительной информацией о размерах изображения.
	ImageInfo struct {
		// ContentType - MIME-тип изображения (например, "image/jpeg", "image/png").
		ContentType string `json:"content_type,omitempty"`

		// OriginalName - оригинальное имя изображения при загрузке.
		OriginalName string `json:"original_name,omitempty"`

		// Name - имя изображения в хранилище (может отличаться от оригинального).
		Name string `json:"name,omitempty"`

		// Path - путь к изображению в хранилище (не сериализуется в JSON).
		Path string `json:"-"`

		// URL - публичный URL для доступа к изображению.
		URL string `json:"url,omitempty"`

		// Width - ширина изображения в пикселях.
		Width int32 `json:"width,omitempty"`

		// Height - высота изображения в пикселях.
		Height int32 `json:"height,omitempty"`

		// Size - размер файла изображения в байтах.
		Size int64 `json:"size,omitempty"`

		// CreatedAt - время создания изображения.
		CreatedAt *time.Time `json:"created_at,omitempty"`

		// UpdatedAt - время последнего обновления изображения.
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}
)

// Original - возвращает наиболее подходящее имя изображения.
// Ищет в порядке приоритета: OriginalName → Name → базовое имя из Path.
func (im ImageInfo) Original() string {
	if im.OriginalName != "" {
		return im.OriginalName
	}

	if im.Name != "" {
		return im.Name
	}

	return path.Base(im.Path)
}

// ToFileInfo - возвращает мета-информацию изображения преобразованное в файловую структуру
// (с потерей дополнительной информации об изображении).
func (im ImageInfo) ToFileInfo() FileInfo {
	return FileInfo{
		ContentType:  im.ContentType,
		OriginalName: im.OriginalName,
		Name:         im.Name,
		Path:         im.Path,
		URL:          im.URL,
		Size:         im.Size,
		CreatedAt:    copyptr.Time(im.CreatedAt),
		UpdatedAt:    copyptr.Time(im.UpdatedAt),
	}
}
