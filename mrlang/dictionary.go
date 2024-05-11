package mrlang

import (
	"fmt"
	"strconv"
)

type (
	// Dictionary - справочник объектов указанного типа на конкретном языке.
	Dictionary struct {
		rows DictionaryMap
	}

	// DictionaryMap - атрибуты объекта на конкретном языке.
	// dictionary-name -> [id -> [attr1 -> text1, attr2 -> text2, ...], ...]
	DictionaryMap map[string]DictionaryItemAttrs
)

func newDictionary(filePath string) (*Dictionary, error) {
	data := make(DictionaryMap, 0)

	if err := parseFile(filePath, &data); err != nil {
		return nil, fmt.Errorf("error parsing dictionary file '%s': %w (see registered dicts: config.yaml:translation/dictionaries/list)", filePath, err)
	}

	return &Dictionary{
		rows: data,
	}, nil
}

// ItemByID - возвращает объект с его атрибутами по ID.
func (d *Dictionary) ItemByID(id int) DictionaryItemAttrs {
	return d.ItemByKey(strconv.Itoa(id))
}

// ItemByKey - возвращает объект с его атрибутами по ключу (строковому ID).
func (d *Dictionary) ItemByKey(key string) DictionaryItemAttrs {
	if text, ok := d.rows[key]; ok {
		return text
	}

	return DictionaryItemAttrs{}
}

// RegisteredItems - возвращает список ключей зарегистрированных объектов.
func (d *Dictionary) RegisteredItems() []string {
	keys := make([]string, 0, len(d.rows))

	for key := range d.rows {
		keys = append(keys, key)
	}

	return keys
}
