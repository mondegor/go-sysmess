// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package errcat

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func NewCatalog() catalog.Catalog {
	dict := map[string]catalog.Dictionary{
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
		"ru_RU": &dictionary{index: ru_RUIndex, data: ru_RUData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	return cat
}

var messageKeyToIndex = map[string]int{
	"InternalError":    0,
	"SystemError":      1,
	"entity not found": 2,
}

var en_USIndex = []uint32{ // 4 elements
	0x00000000, 0x00000033, 0x00000086, 0x00000099,
} // Size: 40 bytes

const en_USData string = "" + // Size: 153 bytes
	"\x02Internal server error\x0a\x0aContact support in telegram\x02System s" +
	"erver error\x0a\x0aThe system is temporarily unavailable, please try aga" +
	"in later\x02Resource not found"

var ru_RUIndex = []uint32{ // 4 elements
	0x00000000, 0x0000006d, 0x00000116, 0x00000135,
} // Size: 40 bytes

const ru_RUData string = "" + // Size: 309 bytes
	"\x02Внутренняя ошибка сервера\x0a\x0aОбратитесь за поддержкой в Telegram" +
	"\x02Системная ошибка сервера\x0a\x0aСистема временно недоступна, пожалуй" +
	"ста, повторите попытку позже\x02Ресурс не найден"

	// Total table size 542 bytes (0KiB); checksum: 113CCAE2
