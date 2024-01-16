package mrlang

import (
	"fmt"
	"strconv"
)

type (
	Dictionary struct {
		rows DictionaryMap
	}

	DictionaryMap map[string]DictionaryItemAttrs // dictionary-name -> [id -> [attr1 -> text1, attr2 -> text2, ...], ...]
)

func newDictionary(filePath string) (*Dictionary, error) {
	data := make(DictionaryMap, 0)

	if err := parseFile(filePath, &data); err != nil {
		return nil, fmt.Errorf("error parsing dictionary file '%s': %w", filePath, err)
	}

	return &Dictionary{
		rows: data,
	}, nil
}

func (d *Dictionary) ItemByID(id int) DictionaryItemAttrs {
	return d.ItemByKey(strconv.Itoa(id))
}

func (d *Dictionary) ItemByKey(key string) DictionaryItemAttrs {
	if text, ok := d.rows[key]; ok {
		return text
	}

	return DictionaryItemAttrs{}
}

func (d *Dictionary) RegisteredItems() []string {
	keys := make([]string, len(d.rows))
	i := 0

	for key := range d.rows {
		keys[i] = key
		i++
	}

	return keys
}
