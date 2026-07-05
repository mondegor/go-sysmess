package mime

import (
	"errors"
	"fmt"
	"strings"
)

type (
	// TypeList - хранит соответствие расширений их типам файлов (в обе стороны).
	TypeList struct {
		contentTypeMap map[string]string
		extensionMap   map[string]string
	}

	// Type - хранит расширение и соответствующий ему тип файла.
	Type struct {
		ContentType string `yaml:"type"`
		Extension   string `yaml:"ext"`
	}
)

// NewTypeList - создаёт объект TypeList на основе списка соответствий расширений и файлов.
func NewTypeList(items []Type) *TypeList {
	mimeMap := make(map[string]string, len(items))
	extMap := make(map[string]string, len(items))

	for _, item := range items {
		item.Extension = strings.TrimPrefix(item.Extension, ".")
		extMap[item.Extension] = item.ContentType

		// т.к. у одного типа может быть несколько расширений,
		// то в индекс попадает только первый зарегистрированный
		if _, ok := mimeMap[item.ContentType]; !ok {
			mimeMap[item.ContentType] = item.Extension
		}
	}

	return &TypeList{
		contentTypeMap: mimeMap,
		extensionMap:   extMap,
	}
}

// TypesByExts - возвращает Type массив, в который войдут указанные расширения,
// если хотя бы одно расширение не зарегистрировано в списке, то будет выдана ошибка.
func (mt *TypeList) TypesByExts(values []string) ([]Type, error) {
	mime := make([]Type, len(values))

	for i, ext := range values {
		contentType, err := mt.ContentTypeByExt(ext)
		if err != nil {
			return nil, err
		}

		mime[i] = Type{
			ContentType: contentType,
			Extension:   ext,
		}
	}

	return mime, nil
}

// ContentTypeByExt - возвращает тип файла по указанному расширению,
// если тип не найден, то возвращается пустая строка.
func (mt *TypeList) ContentTypeByExt(value string) (string, error) {
	if value == "" || len(value) == 1 && value[0] == '.' {
		return "", errors.New("arg 'value' is empty")
	}

	if value[0] == '.' { // если указано расширение с точкой в начале
		value = value[1:]
	}

	if ext, ok := mt.extensionMap[value]; ok {
		return ext, nil
	}

	return "", fmt.Errorf("mime not found for arg '%s'", value)
}

// ExtByContentType - возвращает расширение по указанному типу файла,
// если расширение не найдено, то возвращается пустая строка.
func (mt *TypeList) ExtByContentType(value string) (string, error) {
	if value == "" {
		return "", errors.New("arg 'value' is empty")
	}

	if contentType, ok := mt.contentTypeMap[value]; ok {
		return contentType, nil
	}

	return "", fmt.Errorf("ext not found for arg '%s'", value)
}
