package mrlang

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	langFileType = "yaml"
)

func parseFile(path string, data any) error {
	f, err := os.OpenFile(path, os.O_RDONLY, 0)

	if err != nil {
		return err
	}

	defer f.Close()

	return yaml.NewDecoder(f).Decode(data)
}

func filePath(dirPath, name string) string {
	// dir/lang.ext, ./translate/en.yaml, ./translate/dic/category/en.yaml
	return strings.TrimRight(dirPath, "/") + "/" + strings.Trim(name, "/") + "." + langFileType
}
