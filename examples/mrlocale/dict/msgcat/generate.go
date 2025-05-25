package msgcat

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

//go:generate gotext -srclang=en-US update -out=../../internal/dict/msgcat/catalog.go -lang=en-US,ru-RU github.com/mondegor/go-sysmess/examples/mrlocale/dict/msgcat
//go:generate gotext-catalog-fix -src=../../internal/dict/msgcat/catalog.go -out=../../internal/dict/msgcat/catalog.go

func list() {
	p := message.NewPrinter(language.MustParse("en-US"))

	p.Sprintf("Message example")
	p.Sprintf("Message example with param %s", "param1")
	p.Sprintf("Message example with two params %[1]s and %[2]s", "param1", "param2")
	p.Sprintf("Total %d message(s)", "totalMessageCount")
}
