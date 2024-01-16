package mrerr

type (
	CallerOption func(c *Caller)
)

func CallerDeep(value int) CallerOption {
	return func(c *Caller) {
		if value < 0 {
			value = 0
		}

		c.deep = value
	}
}

func CallerUseShortPath(value bool) CallerOption {
	return func(c *Caller) { c.useShortPath = value }
}

func CallerRootPath(value string) CallerOption {
	return func(c *Caller) { c.rootPath = value }
}
