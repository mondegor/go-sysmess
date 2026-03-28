package mrmodel

import (
	"path"
	"time"

	"github.com/mondegor/go-sysmess/util/copyptr"
)

type (
	// ImageInfo - мета-информация об изображении.
	ImageInfo struct {
		ContentType  string     `json:"content_type,omitempty"`
		OriginalName string     `json:"original_name,omitempty"`
		Name         string     `json:"name,omitempty"`
		Path         string     `json:"-"`
		URL          string     `json:"url,omitempty"`
		Width        int32      `json:"width,omitempty"`
		Height       int32      `json:"height,omitempty"`
		Size         int64      `json:"size,omitempty"`
		CreatedAt    *time.Time `json:"created_at,omitempty"`
		UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	}
)

// Original - возвращает оригинальное имя изображение (как оно было названо в первоисточнике).
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
