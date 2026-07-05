package config

type (
	// FileType - ограничения по загрузке файлов определённого типа.
	FileType struct {
		MinSize                 uint32   `yaml:"min_size"`
		MaxSize                 uint32   `yaml:"max_size"`
		MaxFiles                uint16   `yaml:"max_files"`
		CheckRequestContentType bool     `yaml:"check_request_content_type"`
		Extensions              []string `yaml:"extensions"`
	}

	// ImageType - ограничения по загрузке изображений определённого типа.
	ImageType struct {
		MaxWidth  uint16   `yaml:"max_width"`
		MaxHeight uint16   `yaml:"max_height"`
		CheckBody bool     `yaml:"check_body"`
		File      FileType `yaml:"file"`
	}
)
