package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
)

func main() {
	list := mrerr.NewListWith("fieldEmail", fmt.Errorf("error in email"))
	fmt.Println(list)

	factory := mrerr.NewFactory(
		"errMyErrorWithParams",
		mrerr.ErrorKindUser,
		"my error with '{{ .param1 }}' and '{{ .param2 }}'",
	)

	list = mrerr.NewList(mrerr.FieldError{ID: "fieldPhone", Err: factory.New("123", "567")})

	if list == nil {
		fmt.Println("list is nil")
		return
	}

	tmp := *list

	if len(tmp) == 0 {
		fmt.Println("list is empty")
		return
	}

	fmt.Println(tmp[0].ID)		 // field id
	fmt.Println(tmp[0].Err.ID())   // error id for field
	fmt.Println(tmp[0].Err.Kind()) // 3 - mrerr.ErrorKindUser
	fmt.Println(tmp[0].Err)
}
