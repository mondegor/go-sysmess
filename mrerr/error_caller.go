package mrerr

var (
	caller = NewCaller(
		CallerDeep(2),
		CallerUseShortPath(true),
	)
)

// SetCallerOptions - WARNING: use only when starting the main process
func SetCallerOptions(opts ...CallerOption) {
	caller = NewCaller(opts...)
}
