package mrlang

import (
	"fmt"
)

type (
	// MultiLangDictionary - мультиязычный справочник объектов одного типа.
	MultiLangDictionary struct {
		name           string
		dicByLangIDs   dicByLangIDsMap
		dicByLangCodes dicByLangCodesMap
	}

	dicByLangIDsMap   map[uint16]*Dictionary
	dicByLangCodesMap map[string]*Dictionary
)

// Name - название справочника.
func (d *MultiLangDictionary) Name() string {
	return d.name
}

// ByLangID - возвращает справочник на указанном языке (ID) или ошибку если язык не был найден.
func (d *MultiLangDictionary) ByLangID(id uint16) (*Dictionary, error) {
	if dict, ok := d.dicByLangIDs[id]; ok {
		return dict, nil
	}

	return nil, fmt.Errorf("language with ID=%d is not registered", id)
}

// ByLangCode - возвращает справочник на указанном языке (code) или ошибку если язык не был найден.
func (d *MultiLangDictionary) ByLangCode(code string) (*Dictionary, error) {
	if dict, ok := d.dicByLangCodes[code]; ok {
		return dict, nil
	}

	return nil, fmt.Errorf("language with code '%s' is not registered", code)
}

// RegisteredLangs - возвращает список ключей зарегистрированных языков справочника.
func (d *MultiLangDictionary) RegisteredLangs() []string {
	langs := make([]string, 0, len(d.dicByLangCodes))

	for key := range d.dicByLangCodes {
		langs = append(langs, key)
	}

	return langs
}
