package mapping

import (
	"strconv"
	"strings"

	"github.com/mondegor/go-sysmess/mrerrors"
)

// StackTraceToStrings - функция срезает верхнюю часть стека вызовов,
// которая не несёт в себе информативности. В массиве bounds указываются
// все названия функций и стек будет срезан по самой нижней из них.
func StackTraceToStrings(stack mrerrors.StackTracer) []string {
	if stack == nil {
		return nil
	}

	var buf strings.Builder

	cnt := stack.Count()
	list := make([]string, 0, cnt)

	for i := 0; i < cnt; i++ {
		name, file, line := stack.Source(i)

		if name != "" {
			buf.WriteByte('[')
			buf.WriteString(name)
			buf.WriteString("] ")
		}

		buf.WriteString(file)
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(line))

		list = append(list, buf.String())
		buf.Reset()
	}

	return list
}
