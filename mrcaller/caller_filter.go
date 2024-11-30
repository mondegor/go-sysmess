package mrcaller

import "strings"

const (
	lowerBoundPrefix = "runtime." // runtime.main, runtime.goexit
)

// FilterStackTraceTrimUpper - функция срезает верхнюю часть стека вызовов,
// которая не несёт в себе информативности. В массиве bounds указываются
// все названия функций и стек будет срезан по самой нижней из них.
func FilterStackTraceTrimUpper(bounds []string) func(frames []uintptr) []uintptr {
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

			if _, ok := boundMap[item]; ok {
				if i < length-1 {
					return frames[i+1 : length]
				}

				return nil
			}
		}

		return frames[:length]
	}
}
