package stacktrace

import (
	"runtime"
	"strings"
)

const (
	defaultDepth       = 1
	defaultShowFunc    = true
	stackTraceMaxDepth = 32
	upperBoundPrefix   = "main." // main.main, main.*
)

type (
	// Caller - сборщик стека вызовов функций на основе runtime.Callers.
	// Позволяет настраивать глубину, отображение имён функций и границы стека.
	Caller interface {
		Call() StackTrace
	}

	// StackTrace - готовый стек вызовов функций.
	// Предоставляет итератор для последовательного чтения кадров стека.
	StackTrace interface {
		Iterator() func() (index int, name, file string, line int)
	}

	caller struct {
		depth               int                        // максимальное кол-во элементов в стеке вызовов
		showFuncName        bool                       // отображение названий функций в стеке вызовов
		findBottomBoundFunc func(funcName string) bool // функция распознавания границы стека вызовов, после которой содержится полезная информация
	}
)

// NewCaller - создаёт объект Caller.
func NewCaller(opts ...Option) Caller {
	c := &caller{
		showFuncName: defaultShowFunc,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.depth < 1 || c.depth > stackTraceMaxDepth {
		c.depth = defaultDepth
	}

	if c.findBottomBoundFunc == nil {
		c.findBottomBoundFunc = func(_ string) bool {
			return false
		}
	}

	return c
}

// Call - собирает стек вызовов функций начиная с места вызова этой функции.
// Автоматически отсекает кадры выше функции main (upperBoundPrefix)
// и кадры ниже нижней границы (findBottomBoundFunc).
// Возвращает не более depth кадров.
func (c *caller) Call() StackTrace {
	var pcs [stackTraceMaxDepth]uintptr

	top, n := 0, runtime.Callers(2, pcs[:])

	// TODO: здесь достаточно проверить 3-5 элементов стека: i >= max(0, n - 1 - 3)
	for i := n - 1; i >= 0; i-- {
		if strings.HasPrefix(runtimeFrame(pcs[i]).Name(), upperBoundPrefix) {
			n = i + 1

			break
		}
	}

	if n == 0 {
		return stackTrace(nil)
	}

	// найденная upperBoundPrefix граница сверху уже исключена (i + 1 < n)
	for i := 1; i < n; i++ {
		if c.findBottomBoundFunc(runtimeFrame(pcs[i-1]).Name()) {
			top = i

			continue
		}

		if top > 0 {
			break
		}
	}

	// здесь всегда n >= 1
	n -= top

	if n > c.depth {
		n = c.depth
	}

	items := make([]source, n)

	for i := 0; i < len(items); i++ {
		function, file, line := runtimeFrame(pcs[i+top]).Source()

		items[i] = source{
			file: file,
			line: line,
		}

		if c.showFuncName {
			items[i].function = function
		}
	}

	return stackTrace(items)
}
