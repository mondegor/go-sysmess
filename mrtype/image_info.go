package mrtype

import (
	"path"
	"time"

	"github.com/mondegor/go-sysmess/mrlib/copyptr"
)

type (
	// ImageInfo - мета-информация об изображении.
	ImageInfo struct {
		ContentType  string     `json:"contentType,omitempty"`
		OriginalName string     `json:"originalName,omitempty"`
		Name         string     `json:"name,omitempty"`
		Path         string     `json:"-"`
		URL          string     `json:"url,omitempty"`
		Width        uint64     `json:"width,omitempty"`
		Height       uint64     `json:"height,omitempty"`
		Size         uint64     `json:"size,omitempty"`
		CreatedAt    *time.Time `json:"createdAt,omitempty"`
		UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
	}
)

// Original - возвращает оригинальное имя изображение (как оно было названо в первоисточнике).
func (i *ImageInfo) Original() string {
	if i.OriginalName != "" {
		return i.OriginalName
	}

	if i.Name != "" {
		return i.Name
	}

	return path.Base(i.Path)
}

// ToFileInfo - возвращает мета-информацию изображения преобразованное в файловую структуру
// (с потерей дополнительной информации об изображении).
func (i *ImageInfo) ToFileInfo() FileInfo {
	return FileInfo{
		ContentType:  i.ContentType,
		OriginalName: i.OriginalName,
		Name:         i.Name,
		Path:         i.Path,
		URL:          i.URL,
		Size:         i.Size,
		CreatedAt:    copyptr.Time(i.CreatedAt),
		UpdatedAt:    copyptr.Time(i.UpdatedAt),
	}
}
