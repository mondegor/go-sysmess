package config

import (
	"fmt"

	"github.com/mondegor/go-sysmess/util/mime"
)

// ValidateMimeTypes - валидирует указанный список типов файлов.
func ValidateMimeTypes(mimeTypes []mime.Type) error {
	uniqExtensions := make(map[string]bool, len(mimeTypes))

	for _, mt := range mimeTypes {
		if mt.Extension == "" {
			return fmt.Errorf("mimeType extension is required for content type '%s'", mt.ContentType)
		}

		if uniqExtensions[mt.Extension] {
			return fmt.Errorf("duplicate mimeType extension for content type (ext='%s', type='%s')", mt.Extension, mt.ContentType)
		}

		uniqExtensions[mt.Extension] = true
	}

	return nil
}
