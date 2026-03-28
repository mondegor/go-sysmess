package mrmodel

import (
	"path"
	"time"
)

type (
	// FileInfo - мета-информация о файле.
	FileInfo struct {
		ContentType  string     `json:"content_type,omitempty"`
		OriginalName string     `json:"original_name,omitempty"`
		Name         string     `json:"name,omitempty"`
		Path         string     `json:"-"`
		URL          string     `json:"url,omitempty"`
		Size         int64      `json:"size,omitempty"`
		CreatedAt    *time.Time `json:"created_at,omitempty"`
		UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	}
)

// Original - возвращает оригинальное имя файла (как оно было названо в первоисточнике).
func (f FileInfo) Original() string {
	if f.OriginalName != "" {
		return f.OriginalName
	}

	if f.Name != "" {
		return f.Name
	}

	return path.Base(f.Path)
}
