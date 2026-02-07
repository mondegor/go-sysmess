package stacktrace

import (
	"strconv"
	"strings"
)

const (
	lowerBoundPrefix = "runtime." // runtime.main, runtime.goexit
)

// TrimUpperFilter - функция срезает верхнюю часть стека вызовов,
// которая не несёт в себе полезной информативности. В массиве bounds указываются
// все названия системных функций и стек будет срезан по самой нижней из них.
func TrimUpperFilter(bounds []string) func(frames []uintptr) []uintptr {
	boundMap := make(map[string]bool, len(bounds))
	for _, item := range bounds {
		boundMap[item] = true
	}

	return func(frames []uintptr) []uintptr {
		length := len(frames)

		for i := length - 1; i >= 0; i-- {
			item := runtimeFrame(frames[i]).Name()

			if length == len(frames) && strings.HasPrefix(item, lowerBoundPrefix) {
				length = i // исключая нижнюю границу

				continue
			}

			if boundMap[item] {
				if i < length-1 {
					return frames[i+1 : length]
				}

				return nil
			}
		}

		return frames[:length]
	}
}

// ToStrings - функция возвращает стек вызовов в виде строки.
func ToStrings(stackIterator func() (index int, name, file string, line int)) []string {
	if stackIterator == nil {
		return nil
	}

	var buf strings.Builder

	list := make([]string, 0, stackTraceMaxDepth)

	for {
		index, name, file, line := stackIterator()
		if index < 0 {
			break
		}

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
