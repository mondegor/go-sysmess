package mrerr

var (
	caller = NewCaller(
		CallerOptions{
			Deep:         2,
			UseShortPath: true,
		})
)

// SetCallerOptions - WARNING: use only when starting the main process
func SetCallerOptions(opt CallerOptions) {
	caller = NewCaller(opt)
}
