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
            DirPath: langsDir,
            FileType: "yaml",
            LangCodes: []string{"en", "ru"},
            DefaultLang: "ru",
        },
    )

    if err != nil {
        fmt.Println(err)
        return
    }

    defaultLoc := tr.DefaultLocale()
    fmt.Printf("DefaultLoc: %s\n", defaultLoc.LangCode())

    for _, loc := range tr.RegisteredLocales() {
        fmt.Printf("LangCode: %s\n", loc.LangCode())

        locTest := tr.FindFirstLocale(loc.LangCode())

        fmt.Println(locTest.LangCode())
        fmt.Println(locTest.TranslateMessage("msgExample", "Default message for msgExample"))
        fmt.Println(locTest.TranslateError("errInternal", "Default error message for errInternal"))
    }

    // fr - not found
    locTest := tr.FindFirstLocale("fr")
    fmt.Println(locTest.LangCode())
}
