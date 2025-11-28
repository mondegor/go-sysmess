package config

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrlib/extfile"
)

// ValidateMimeTypes - валидирует указанный список типов файлов.
func ValidateMimeTypes(mimeTypes []extfile.MimeType) error {
	uniqExtensions := make(map[string]struct{}, len(mimeTypes))

	for _, mime := range mimeTypes {
		if mime.Extension == "" {
			return fmt.Errorf("mimeType extension is required for content type '%s'", mime.ContentType)
		}

		if _, ok := uniqExtensions[mime.Extension]; ok {
			return fmt.Errorf("duplicate mimeType extension for content type (ext='%s', type='%s')", mime.Extension, mime.ContentType)
		}

		uniqExtensions[mime.Extension] = struct{}{}
	}

	return nil
}
