package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// go install ./cmd/gotext-catalog-fix

type (
	config struct {
		sourceFile string
		outputFile string
	}
)

func main() {
	cfg := config{}

	flag.StringVar(&cfg.sourceFile, "src", "", "source file")
	flag.StringVar(&cfg.outputFile, "out", "", "output file to write to")
	flag.Parse()

	cfg, err := parseArgs(cfg)
	if err != nil {
		log.Fatalf("parse args: %v", err)
	}

	data, err := os.ReadFile(cfg.sourceFile)
	if err != nil {
		log.Fatalf("read file: %v", err)
	}

	data = replace(
		data,
		"\n\t\"golang.org/x/text/message\"", "",
		"func init() {", "func NewCatalog() catalog.Catalog {",
		"message.DefaultCatalog = cat", "return cat",
	)

	if err = os.WriteFile(cfg.outputFile, data, 0o600); err != nil { //nolint:gosec
		log.Fatalf("write file: %v", err)
	}
}

func parseArgs(cfg config) (config, error) {
	if err := validatePath(cfg.sourceFile); err != nil {
		return config{}, fmt.Errorf("the source file error (-src): %w", err)
	}

	if err := isPathExist(cfg.sourceFile); err != nil {
		return config{}, fmt.Errorf("the source file error (-src): %w", err)
	}

	if err := validatePath(cfg.outputFile); err != nil {
		return config{}, fmt.Errorf("the output file error (-out): %w", err)
	}

	if err := isDir(path.Dir(cfg.outputFile)); err != nil {
		return config{}, fmt.Errorf("the output dir error (-out): %w", err)
	}

	return cfg, nil
}

func validatePath(value string) error {
	if value == "" {
		return errors.New("path is empty")
	}

	if !strings.HasPrefix(value, "./") && !strings.HasPrefix(value, "../") {
		return errors.New("path must be relative and start with './' or '../'")
	}

	return nil
}

func isPathExist(value string) error {
	if _, err := os.Stat(value); err != nil {
		return fmt.Errorf("os.Stat error: %w", err)
	}

	return nil
}

func isDir(value string) error {
	fi, err := os.Stat(value)
	if err != nil {
		return fmt.Errorf("os.Stat error: %w", err)
	}

	if !fi.IsDir() {
		return errors.New("value is not a directory")
	}

	return nil
}

func replace(s []byte, newold ...string) []byte {
	for i := 1; i < len(newold); i += 2 {
		s = bytes.Replace(s, []byte(newold[i-1]), []byte(newold[i]), 1)
	}

	return s
}
