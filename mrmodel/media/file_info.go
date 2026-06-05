package media

import (
	"path"
	"time"
)

type (
	// FileInfo - мета-информация о файле.
	// Используется для описания файловых данных без их содержимого.
	FileInfo struct {
		// ContentType - MIME-тип файла (например: "image/jpeg", "application/pdf").
		ContentType string `json:"content_type,omitempty"`

		// OriginalName - оригинальное имя файла при загрузке.
		OriginalName string `json:"original_name,omitempty"`

		// Name - имя файла в хранилище (может отличаться от оригинального).
		Name string `json:"name,omitempty"`

		// Path - путь к файлу в хранилище (не сериализуется в JSON).
		Path string `json:"-"`

		// URL - публичный URL для доступа к файлу.
		URL string `json:"url,omitempty"`

		// Size - размер файла в байтах.
		Size int64 `json:"size,omitempty"`

		// CreatedAt - время создания файла.
		CreatedAt *time.Time `json:"created_at,omitempty"`

		// UpdatedAt - время последнего обновления файла.
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}
)

// Original - возвращает наиболее подходящее имя файла.
// Ищет в порядке приоритета: OriginalName → Name → базовое имя из Path.
func (f FileInfo) Original() string {
	if f.OriginalName != "" {
		return f.OriginalName
	}

	if f.Name != "" {
		return f.Name
	}

	return path.Base(f.Path)
}
