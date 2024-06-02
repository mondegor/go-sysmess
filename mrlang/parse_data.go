package mrlang

import (
	"log"
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

	defer func() {
		if err := f.Close(); err != nil {
			log.Println("parseFile: error when closing file:", err)
		}
	}()

	return yaml.NewDecoder(f).Decode(data)
}

func getFilePath(dirPath, name string) string {
	// dir/lang.ext: ./translate/en_EN.yaml, ./translate/dic/category/en_EN.yaml
	return strings.TrimRight(dirPath, "/") + "/" + strings.Trim(name, "/") + "." + langFileType
}
