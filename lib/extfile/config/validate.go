package config

import (
	"fmt"

	"github.com/mondegor/go-sysmess/lib/extfile"
)

// ValidateMimeTypes - валидирует указанный список типов файлов.
func ValidateMimeTypes(mimeTypes []extfile.MimeType) error {
	uniqExtensions := make(map[string]bool, len(mimeTypes))

	for _, mime := range mimeTypes {
		if mime.Extension == "" {
			return fmt.Errorf("mimeType extension is required for content type '%s'", mime.ContentType)
		}

		if uniqExtensions[mime.Extension] {
			return fmt.Errorf("duplicate mimeType extension for content type (ext='%s', type='%s')", mime.Extension, mime.ContentType)
		}

		uniqExtensions[mime.Extension] = true
	}

	return nil
}
