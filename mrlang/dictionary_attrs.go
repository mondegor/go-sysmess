package mrlang

type (
	// DictionaryItemAttrs - список имён атрибутов и их значений.
	DictionaryItemAttrs map[string]string
)

// Attr - возвращает значение атрибута по его имени или defaultText если имя не найдено.
func (a DictionaryItemAttrs) Attr(name, defaultText string) string {
	if text, ok := a[name]; ok {
		return text
	}

	return defaultText
}

// RegisteredAttrs - возвращает список имён зарегистрированных атрибутов объекта.
func (a DictionaryItemAttrs) RegisteredAttrs() []string {
	attrs := make([]string, 0, len(a))

	for key := range a {
		attrs = append(attrs, key)
	}

	return attrs
}
