package fruitcat

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//go:generate gotext -srclang=en-US update -out=../../internal/dict/fruitcat/catalog.go -lang=en-US,ru-RU github.com/mondegor/go-sysmess/examples/mrlocale/dict/fruitcat
//go:generate gotext-catalog-fix -src=../../internal/dict/fruitcat/catalog.go -out=../../internal/dict/fruitcat/catalog.go

func list() {
	p := message.NewPrinter(language.MustParse("en-US"))

	p.Sprintf("id7391")  // 7391|Pineapple|Description
	p.Sprintf("id97104") // 97104|Plum|Description
}
