package mrlang

import (
	"fmt"
)

type (
	dicByLangIDsMap   map[uint16]*Dictionary
	dicByLangCodesMap map[string]*Dictionary

	MultiLangDictionary struct {
		name           string
		dicByLangIDs   dicByLangIDsMap
		dicByLangCodes dicByLangCodesMap
	}
)

func (d *MultiLangDictionary) Name() string {
	return d.name
}

func (d *MultiLangDictionary) ByLangID(id uint16) (*Dictionary, error) {
	if dict, ok := d.dicByLangIDs[id]; ok {
		return dict, nil
	}

	return nil, fmt.Errorf("language with ID=%d is not registered", id)
}

func (d *MultiLangDictionary) ByLangCode(code string) (*Dictionary, error) {
	if dict, ok := d.dicByLangCodes[code]; ok {
		return dict, nil
	}

	return nil, fmt.Errorf("language with code '%s' is not registered", code)
}

func (d *MultiLangDictionary) RegisteredLangs() []string {
	langs := make([]string, len(d.dicByLangCodes))
	i := 0

	for key := range d.dicByLangCodes {
		langs[i] = key
		i++
	}

	return langs
}
