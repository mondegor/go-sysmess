package mrerr

import (
	"runtime"
	"strings"
)

const (
	callerStackBreak = "src/runtime/proc.go"
)

type (
	Caller struct {
		deep         int
		useShortPath bool
		rootPath     string
	}

	CallStackRow struct {
		File string
		Line int
	}
)

func NewCaller(opts ...CallerOption) *Caller {
	c := &Caller{}
	c.applyOptions(opts)

	return c
}

func (c *Caller) applyOptions(opts []CallerOption) {
	for _, f := range opts {
		f(c)
	}
}

func (c *Caller) CallStack(skip int) []CallStackRow {
	cs := make([]CallStackRow, 0, c.deep)

	for i := 0; i < c.deep; i++ {
		_, file, line, ok := runtime.Caller(skip + i + 1)

		if !ok || strings.HasSuffix(file, callerStackBreak) {
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

	var i, sepIndex int

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
