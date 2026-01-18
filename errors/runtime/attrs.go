package runtime

const (
	emptyAttrKey     = "!EMPTYATTRKEY"
	badAttrKey       = "!BADATTRKEY"
	missingAttrValue = "!MISSINGATTRVALUE"
)

// popKeyValue - возвращает извлечённый из массива атрибутов два первых элемента:
// ключ и его значение, а также сам массив с оставшимися элементами.
func popKeyValue(attrs []any) (key string, value any, rest []any) {
	switch v := attrs[0].(type) {
	case string:
		if v == "" {
			v = emptyAttrKey
		}

		if len(attrs) == 1 {
			return v, missingAttrValue, nil
		}

		return v, attrs[1], attrs[2:]
	default:
		return badAttrKey, v, attrs[1:]
	}
}
