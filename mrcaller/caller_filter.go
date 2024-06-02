package mrcaller

// FilterStackTraceTrimTop - функция срезает верхнюю часть стека вызовов,
// которая не несёт в себе информативности. В массиве borders указываются
// все названия функций и стек будет срезан по самой нижней из них.
func FilterStackTraceTrimTop(borders []string) func(frames []uintptr) []uintptr {
	borderMap := make(map[string]bool, len(borders))
	for _, item := range borders {
		borderMap[item] = true
	}

	return func(frames []uintptr) []uintptr {
		for i := len(frames) - 1; i >= 0; i-- {
			item := runtimeFrame(frames[i]).Name()

			if _, ok := borderMap[item]; ok {
				if i < len(frames)-1 {
					return frames[i+1:]
				}

				return nil
			}
		}

		return frames
	}
}
