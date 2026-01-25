package main

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/mondegor/go-sysmess/errors"
	"github.com/mondegor/go-sysmess/errors/helper"
	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/examples/mrlocale/internal/dict/errcat"
	"github.com/mondegor/go-sysmess/examples/mrlocale/internal/dict/fruitcat"
	"github.com/mondegor/go-sysmess/examples/mrlocale/internal/dict/msgcat"
	"github.com/mondegor/go-sysmess/mrlocale"
	"github.com/mondegor/go-sysmess/mrlocale/provider/gotext"
)

// main - пример формирование сообщений, ошибок и свойств объектов на указанном языке.
func main() {
	var (
		localeProvider mrlocale.MessageProvider
		err            error
	)

	bundle, err := mrlocale.NewBundle(
		[]string{"ru-RU", "en-US"},
		mrlocale.WithDefaultLanguage("ru-RU"),
		mrlocale.WithFormatError(helper.ExtractMessageForLocalization),
		mrlocale.WithMessageProvider(
			func(languages []language.Tag) (mrlocale.MessageProvider, error) {
				localeProvider, err = gotext.NewProvider(
					languages,
					gotext.WithDomainCatalog(mrlocale.DefaultMessagesDomain, msgcat.NewCatalog()),
					gotext.WithDomainCatalog(mrlocale.DefaultErrorsDomain, errcat.NewCatalog()),
					gotext.WithDomainCatalog("fruitcat", fruitcat.NewCatalog()),
				)

				return localeProvider, err
			},
		),
	)
	if err != nil {
		panic(err)
	}

	pool := mrlocale.NewPool(bundle)
	lz := pool.Localizer(language.MustParse("ru-RU"))

	fmt.Printf("language: %s\n", lz.Language())
	fmt.Printf("message: %s\n", lz.Translate("Message example"))

	fmt.Println("--------------------------------------------------")

	errorMessage := errcat.ParseErrorMessage(lz.TranslateError(errors.NewInternalError("my-error")))
	fmt.Printf("error: %s\n", errorMessage.Reason)
	fmt.Printf("error-details: %s\n", errorMessage.Details)

	fmt.Println("--------------------------------------------------")

	for _, lang := range localeProvider.Languages() {
		lz = pool.Localizer(lang)

		fmt.Printf("language: %s\n", lang)
		fmt.Println("..................................................")

		fmt.Println(lz.Translate("Message example with param %s", "param1"))
		fmt.Println(lz.Translate("Message example with two params %[1]s and %[2]s", "param1", "param2"))
		fmt.Println(lz.Translate("Total %d message(s)", 0))
		fmt.Println(lz.Translate("Total %d message(s)", 1))
		fmt.Println(lz.Translate("Total %d message(s)", 21))
		fmt.Println(lz.Translate("Total %d message(s)", 43))
		fmt.Println(lz.Translate("Total %d message(s)", 67))

		fmt.Println("..................................................")

		errMess := errcat.ParseErrorMessage(lz.TranslateError(errors.ErrInternalNilPointer.New()))
		fmt.Println("reason: '" + errMess.Reason + "'; details: '" + errMess.Details + "'")

		errMess = errcat.ParseErrorMessage(lz.TranslateError(errors.ErrSystemStorageConnectionFailed.New()))
		fmt.Println("reason: '" + errMess.Reason + "'; details: '" + errMess.Details + "'")

		fmt.Println(lz.TranslateError(errors.ErrUseCaseEntityNotFound))

		fmt.Println("..................................................")

		fruitMess := fruitcat.ParseMessage(lz.TranslateInDomain("fruitcat", "id7391"))
		fmt.Println("ID: '" + fruitMess.ID + "'; Name: '" + fruitMess.Name + "'; Description: '" + fruitMess.Description + "'")

		fruitMess = fruitcat.ParseMessage(lz.TranslateInDomain("fruitcat", "id97104"))
		fmt.Println("ID: '" + fruitMess.ID + "'; Name: '" + fruitMess.Name + "'; Description: '" + fruitMess.Description + "'")

		fmt.Println("--------------------------------------------------")
	}

	for _, domain := range localeProvider.Domains() {
		fmt.Printf("domain=%s\n", domain)
	}

	fmt.Println("--------------------------------------------------")

	// fr - not found
	lz = pool.Localizer(language.MustParse("fr"))
	fmt.Println(lz.Language()) // print default lang
}

type localizedError interface {
	Kind() kind.Enum
	Code() string
}
