package config

type (
	// FileType - ограничения по загрузке файлов определённого типа.
	FileType struct {
		MinSize                 uint64   `yaml:"min_size"`
		MaxSize                 uint64   `yaml:"max_size"`
		MaxFiles                int      `yaml:"max_files"`
		CheckRequestContentType bool     `yaml:"check_request_content_type"`
		Extensions              []string `yaml:"extensions"`
	}

	// ImageType - ограничения по загрузке изображений определённого типа.
	ImageType struct {
		MaxWidth  uint64   `yaml:"max_width"`
		MaxHeight uint64   `yaml:"max_height"`
		CheckBody bool     `yaml:"check_body"`
		File      FileType `yaml:"file"`
	}
)
