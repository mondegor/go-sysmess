package main

import (
    "fmt"

    "github.com/mondegor/go-sysmess/mrerr"
)

var (
    FactoryError = mrerr.NewFactory(
        "errMyErrorWithParams",
        mrerr.ErrorKindInternal,
        "my error with '{{ .param1 }}' and '{{ .param2 }}'",
    )
)

func main() {
    err := createErr()

    fmt.Println(err)
    fmt.Println(err.TraceID())
    fmt.Println(err.ID())
    fmt.Println(err.Kind()) // 0 - mrerr.ErrorKindInternal

    err = createErrLevel2()
    fmt.Println(err)
}

func createErr() *mrerr.AppError {
    return FactoryError.New("my-param1", 123456)
}

func createErr2() *mrerr.AppError {
    return FactoryError.Caller(1).New("my-param2", 555555)
}

func createErrLevel2() *mrerr.AppError {
    return createErr2()
}
