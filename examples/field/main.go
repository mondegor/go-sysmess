package main

import (
    "fmt"

    "github.com/mondegor/go-sysmess/mrerr"
)

func main() {
    list := mrerr.NewListWith("fieldEmail", fmt.Errorf("error in email") )
    fmt.Println(list)

    factory := mrerr.NewFactory(
        "errMyErrorWithParams",
        mrerr.ErrorKindUser,
        "my error with '{{ .param1 }}' and '{{ .param2 }}'",
    )

    list = mrerr.NewListWith("fieldPhone", factory.New("123", "567") )

    fmt.Println(list[0].Id) // field id
    fmt.Println(list[0].Err.Id()) // error id for field
    fmt.Println(list[0].Err.Kind()) // 3 - mrerr.ErrorKindUser
    fmt.Println(list[0].Err)
}
