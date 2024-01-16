package main

import (
	"flag"
	"fmt"

	"github.com/mondegor/go-sysmess/mrlang"
)

var (
	langsDir string
)

func init() {
	flag.StringVar(&langsDir, "langs-dir", "./examples/translator/langs", "Dir to language files")
}

func main() {
	flag.Parse()

	tr, err := mrlang.NewTranslator(
		mrlang.TranslatorOptions{
			DirPath:           langsDir,
			LangCodes:         []string{"en_EN", "ru_RU"},
			DefaultLang:       "ru_RU",
			DictionaryDirPath: langsDir + "/dic",
			Dictionaries:      []string{"category"},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	defaultLoc := tr.DefaultLocale()
	fmt.Printf("DefaultLoc: %s\n", defaultLoc.LangCode())

	for _, localeCode := range tr.RegisteredLocales() {
		locale, _ := tr.LocaleByCode(localeCode)
		fmt.Printf("LangCode: %s\n", localeCode)

		locTest := tr.FindFirstLocale(localeCode)

		fmt.Printf("ID=%d, code=%s\n", locale.LangID(), localeCode)
		fmt.Println(locTest.TranslateMessage("msgExample", "Default message for msgExample"))
		fmt.Println(locTest.TranslateError("errInternal", "Default error message for errInternal"))

		for _, dicName := range tr.RegisteredDictionaries() {
			mdict, _ := tr.Dictionary(dicName)
			fmt.Printf("dictionary=%s\n", mdict.Name())

			dict, _ := mdict.ByLangCode(localeCode)

			for _, itemID := range dict.RegisteredItems() {
				fmt.Printf("  - ID=%s\n", itemID)
				dictItem := dict.ItemByKey(itemID)

				for _, attrName := range dictItem.RegisteredAttrs() {
					fmt.Printf("    - attr=%s, value=%s\n", attrName, dictItem.Attr(attrName, "UNKNOWN"))
				}
			}
		}
	}

	// fr - not found
	locTest := tr.FindFirstLocale("fr")
	fmt.Println(locTest.LangCode())
}
