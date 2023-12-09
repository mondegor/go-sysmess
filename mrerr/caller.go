package mrerr

import "runtime"

type (
	Caller struct {
		deep         int
		useShortPath bool
		rootPath     string
	}

	CallerOptions struct {
		Deep         int
		UseShortPath bool
		RootPath     string
	}

	CallStackRow struct {
		File string
		Line int
	}
)

func NewCaller(opt CallerOptions) *Caller {
	if opt.Deep < 0 {
		opt.Deep = 0
	}

	return &Caller{
		deep:         opt.Deep,
		useShortPath: opt.UseShortPath,
		rootPath:     opt.RootPath,
	}
}

func (c *Caller) CallStack(skip int) []CallStackRow {
	cs := make([]CallStackRow, 0, c.deep)

	for i := 0; i < c.deep; i++ {
		_, file, line, ok := runtime.Caller(skip + i + 1)

		if !ok {
			break
		}

		if file == "" {
			file = "???"
		}

		cs = append(cs, CallStackRow{file, line})
	}

	if c.useShortPath {
		pattern := c.rootPath

		for i := range cs {
			if pattern == "" {
				pattern = cs[i].File
				continue
			}

			cs[i].File = c.shortFileName(pattern, cs[i].File)
		}
	}

	return cs
}

func (c *Caller) shortFileName(pattern, file string) string {
	minLen := len(pattern)

	if minLen > len(file) {
		minLen = len(file)
	}

	var (
		i, sepIndex int
	)

	for ; i < minLen; i++ {
		if pattern[i] != file[i] {
			break
		}

		if pattern[i] == '/' || pattern[i] == '\\' {
			sepIndex = i
		}
	}

	if sepIndex > 0 {
		return "..." + file[sepIndex:]
	}

	return file
}
