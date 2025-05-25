package errcat

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//go:generate gotext -srclang=en-US update -out=../../internal/dict/errcat/catalog.go -lang=en-US,ru-RU github.com/mondegor/go-sysmess/examples/mrlocale/dict/errcat
//go:generate gotext-catalog-fix -src=../../internal/dict/errcat/catalog.go -out=../../internal/dict/errcat/catalog.go

func list() {
	p := message.NewPrinter(language.MustParse("en-US"))

	p.Sprintf("InternalError")
	p.Sprintf("SystemError")
	p.Sprintf("entity not found")
}
