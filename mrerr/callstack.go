package mrerr

type (
	CallStackOptions struct {
		Deep         int
		UseShortPath bool
		RootPath     string
	}
)

var (
	callStackOptions = CallStackOptions{
		Deep:         2,
		UseShortPath: true,
	}
)

// SetCallStackOptions - WARNING: use only when starting the main process
func SetCallStackOptions(opt CallStackOptions) {
	if opt.Deep > 0 {
		callStackOptions.Deep = opt.Deep
	}

	callStackOptions.UseShortPath = opt.UseShortPath
	callStackOptions.RootPath = opt.RootPath
}

func shortFileName(pattern, file string) string {
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
