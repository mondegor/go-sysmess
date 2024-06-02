package mrcaller

type (
	// CallerOption - настройка объекта Caller.
	CallerOption func(c *Caller)
)

// DepthOption - устанавливает опцию depth для Caller.
// Если value меньше 1 или больше callStackMaxDepth, то устанавливается 1.
func DepthOption(value int) CallerOption {
	return func(c *Caller) { c.depth = value }
}

// FilterStackTraceOption - функцию фильтрации стека вызовов.
func FilterStackTraceOption(fn func(frames []uintptr) []uintptr) CallerOption {
	return func(c *Caller) { c.filterStackTraceFunc = fn }
}
