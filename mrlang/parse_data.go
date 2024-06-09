package mrlang

import (
	"fmt"
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
		return fmt.Errorf("parse file error: %w", err)
	}

	defer f.Close()

	err = yaml.NewDecoder(f).Decode(data)
	if err != nil {
		return fmt.Errorf("decode data error: %w", err)
	}

	return nil
}

func getFilePath(dirPath, name string) string {
	// dir/lang.ext: ./translate/en_EN.yaml, ./translate/dic/category/en_EN.yaml
	return strings.TrimRight(dirPath, "/") + "/" + strings.Trim(name, "/") + "." + langFileType
}
