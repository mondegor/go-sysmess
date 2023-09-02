package main

import (
    "fmt"

    "github.com/mondegor/go-sysmess/mrerr"
)

func main() {
    factory := mrerr.NewFactory(
        "errMyErrorWithParams",
        mrerr.ErrorKindInternal,
        "my error with '{{ .param1 }}' and '{{ .param2 }}'",
    )

    err := factory.New("my-param1", 123456)

    fmt.Println(err)
    fmt.Println(err.EventId())
    fmt.Println(err.Id())
    fmt.Println(err.Kind()) // 0 - mrerr.ErrorKindInternal
}
