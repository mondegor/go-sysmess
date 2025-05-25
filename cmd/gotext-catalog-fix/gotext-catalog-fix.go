package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

type config struct {
	sourceFile string
	outputFile string
}

func main() {
	cfg, err := parseArgs()
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

	if err = os.WriteFile(cfg.outputFile, data, 0o600); err != nil {
		log.Fatalf("write file: %v", err)
	}
}

func parseArgs() (config, error) {
	cfg := config{}

	flag.StringVar(&cfg.sourceFile, "src", "", "source file")
	flag.StringVar(&cfg.outputFile, "out", "", "output file to write to")
	flag.Parse()

	if cfg.sourceFile == "" {
		return config{}, errors.New("the source file is empty (-src)")
	}

	if cfg.sourceFile[0] != '.' {
		return config{}, errors.New("the source file must be relative and start with '.' (-src)")
	}

	if err := isPathExist(cfg.sourceFile); err != nil {
		return config{}, fmt.Errorf("the source file error (-src): %w", err)
	}

	if cfg.outputFile == "" {
		return config{}, errors.New("the output file is empty (-out)")
	}

	if cfg.outputFile[0] != '.' {
		return config{}, errors.New("the output file must be relative and start with '.' (-out)")
	}

	if err := isDir(path.Dir(cfg.outputFile)); err != nil {
		return config{}, fmt.Errorf("the output dir error (-out): %w", err)
	}

	return cfg, nil
}

func isPathExist(value string) error {
	_, err := os.Stat(value)

	return err //nolint:wrapcheck
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
	for i := 0; i < len(newold); i += 2 {
		s = bytes.Replace(s, []byte(newold[i]), []byte(newold[i+1]), 1)
	}

	return s
}
