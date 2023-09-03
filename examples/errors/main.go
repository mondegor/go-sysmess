package main

import (
    "fmt"

    "github.com/mondegor/go-sysmess/mrerr"
)

var FactoryError = mrerr.NewFactory(
    "errMyErrorWithParams",
    mrerr.ErrorKindInternal,
    "my error with '{{ .param1 }}' and '{{ .param2 }}'",
)

func main() {
    err := createErr()

    fmt.Println(err)
    fmt.Println(err.TraceId())
    fmt.Println(err.Id())
    fmt.Println(err.Kind()) // 0 - mrerr.ErrorKindInternal
}

func createErr() *mrerr.AppError {
    return FactoryError.Caller(1).New("my-param1", 123456)
}
