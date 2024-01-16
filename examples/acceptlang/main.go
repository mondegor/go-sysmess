package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrlang"
)

func main() {
	langs := mrlang.ParseAcceptLanguage("ru;q=0.9, fr-CH, fr-EN, fr;q=0.8, en-REGion01, en;q=0.7, *;q=0.5")

	fmt.Printf("langs: %v", langs)
}
