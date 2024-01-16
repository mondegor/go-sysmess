package mrlang

type (
	DictionaryItemAttrs map[string]string
)

func (a DictionaryItemAttrs) Attr(name, defaultText string) string {
	if text, ok := a[name]; ok {
		return text
	}

	return defaultText
}

func (a DictionaryItemAttrs) RegisteredAttrs() []string {
	attrs := make([]string, len(a))
	i := 0

	for key := range a {
		attrs[i] = key
		i++
	}

	return attrs
}
