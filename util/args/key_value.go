package args

// const (
// 	emptyKey     = "!EMPTYKEY"
// 	badKey       = "!BADKEY"
// 	missingValue = "!MISSINGVALUE"
// )

// // PopKeyValue - возвращает извлечённый из массива два первых элемента:
// // ключ и его значение, а также сам массив с оставшимися элементами.
// func PopKeyValue(args []any) (key string, value any, rest []any) {
// 	switch v := args[0].(type) {
// 	case string:
// 		if v == "" {
// 			v = emptyKey
// 		}
//
// 		if len(args) == 1 {
// 			return v, missingValue, nil
// 		}
//
// 		return v, args[1], args[2:]
// 	default:
// 		return badKey, v, args[1:]
// 	}
// }

// // PushKeyValue - возвращает массив со вставленным в его конец пары ключ/значение.
// // В случае, если ключ уже имеется в массиве, то его значение может быть заменено
// // на значение, которое будет возвращено в функции valueFunc, при этом в index
// // будет указан индекс текущего элемента массива расположенного за ключом,
// // а в foundValue само значение этого элемента. Если ключ не найден, то index будет равен -1.
// func PushKeyValue(args []any, key string, valueFunc func(index int, foundValue any) (newValue any)) []any {
// 	for i := 0; i < len(args); i += 2 {
// 		k, ok := args[i].(string)
// 		if !ok || k != key {
// 			continue
// 		}
//
// 		// если сразу же после последнего ключа нет значения
// 		if i+1 >= len(args) {
// 			args = append(args, missingValue)
//
// 			break
// 		}
//
// 		args[i+1] = valueFunc(i+1, args[i+1])
//
// 		return args
// 	}
//
// 	// если искомый ключ не найден, то он добавляется вместе с его значением
// 	return append(args, key, valueFunc(-1, nil))
// }

// // ReplaceStrategy - возвращает функцию, используемую в PushKeyValue.
// // Реализована самая простая стратегия - в случае существования ключа,
// // его значение будет заменено на новое.
// func ReplaceStrategy(value any) func(index int, foundValue any) (newValue any) {
// 	return func(_ int, _ any) (newValue any) {
// 		return value
// 	}
// }

// // KeyValueToString - возвращает преобразованный массив any в массив string.
// // В указанном массиве последовательно должны располагаться ключ и значение,
// // причём ключи должны быть строкового типа. А все значения будут преобразованы в строковый тип.
// func KeyValueToString(kv []any) []string {
// 	if len(kv) == 0 {
// 		return nil
// 	}
//
// 	var (
// 		key   string
// 		value any
// 	)
//
// 	args := make([]string, 0, countPairs(kv)*2)
//
// 	for len(kv) > 0 {
// 		key, value, kv = PopKeyValue(kv)
// 		args = append(args, key, ToString(value))
// 	}
//
// 	return args
// }

// // KeyValueToValue - удаляет из указанного массива все ключи и возвращает только значения.
// // В указанном массиве последовательно должны располагаться ключ и значение,
// // причём ключи должны быть строкового типа.
// func KeyValueToValue(kv []any) []any {
// 	if len(kv) == 0 {
// 		return nil
// 	}
//
// 	var value any
//
// 	args := make([]any, 0, countPairs(kv))
//
// 	for len(kv) > 0 {
// 		_, value, kv = PopKeyValue(kv)
// 		args = append(args, value)
// 	}
//
// 	return args
// }

// // StringMapToKeyValue - возвращает преобразованную map массив с последовательным расположением - ключ и значение.
// func StringMapToKeyValue(m map[string]string) []any {
// 	for len(m) == 0 {
// 		return nil
// 	}
//
// 	data := make([]any, 0, len(m)*2)
//
// 	for k, v := range m {
// 		data = append(data, k, v)
// 	}
//
// 	return data
// }

// // KeyValueToStringMap - возвращает преобразованный массив в map.
// // В указанном массиве последовательно должны располагаться ключ и значение,
// // причём ключи должны быть строкового типа. А все значения будут преобразованы в строковый тип.
// // Если аргументы не указаны, то вернётся неинициализированная map.
// func KeyValueToStringMap(kv []any) map[string]string {
// 	if len(kv) == 0 {
// 		return nil
// 	}
//
// 	data := make(map[string]string, countPairs(kv))
//
// 	var (
// 		key   string
// 		value any
// 	)
//
// 	for len(kv) > 0 {
// 		key, value, kv = PopKeyValue(kv)
// 		data[key] = ToString(value)
// 	}
//
// 	return data
// }

// // KeyValueToAnyMap - возвращает преобразованный массив в map.
// // В указанном массиве последовательно должны располагаться ключ и значение,
// // причём ключи должны быть строкового типа.
// // Если аргументы не указаны, то вернётся неинициализированная map.
// func KeyValueToAnyMap(kv []any) map[string]any {
// 	if len(kv) == 0 {
// 		return nil
// 	}
//
// 	data := make(map[string]any, countPairs(kv))
//
// 	var (
// 		key   string
// 		value any
// 	)
//
// 	for len(kv) > 0 {
// 		key, value, kv = PopKeyValue(kv)
// 		data[key] = value
// 	}
//
// 	return data
// }

// func countPairs(args []any) int {
// 	n := 0
//
// 	for i := 0; i < len(args); i++ {
// 		n++
//
// 		if _, ok := args[i].(string); ok {
// 			i++
// 		}
// 	}
//
// 	return n
// }
