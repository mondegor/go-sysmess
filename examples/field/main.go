package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
)

func main() {
	fieldError := mrerr.NewFieldMessage("fieldEmail", "error in email")
	fmt.Println(fieldError)

	factory := mrerr.NewFactory(
		"errMyErrorWithParams",
		mrerr.ErrorKindUser,
		"my error with '{{ .param1 }}' and '{{ .param2 }}'",
	)

	list := mrerr.FieldErrorList{mrerr.NewFieldErrorAppErr("fieldPhone", factory.New("123", "567"))}

	list.Add("field1", nil)
	list.AddAppErr("field2", factory.New("p1-22", "p2-22"))
	list.Add("field3", factory.New("p1-33", "p2-33"))
	list.Add("field4", factory.New("p1-44", "p2-44"))

	addSomeItems(&list)

	for i, item := range list {
		fmt.Printf("\nitem %d:\n", i)
		fmt.Printf(
			"FieldID=%s, ErrID=%s, ErrKind=%v, Err=%s\n",
			item.ID(),              // field id
			item.AppError().ID(),   // error id for field
			item.AppError().Kind(), // 3 - mrerr.ErrorKindUser
			item.AppError(),
		)
	}
}

func addSomeItems(list *mrerr.FieldErrorList) {
	factory := mrerr.NewFactory(
		"errSomeItems",
		mrerr.ErrorKindSystem,
		"my error with '{{ .param1 }}'",
	)

	list.AddAppErr("field5", nil)
	list.AddAppErr("field6", factory.New("p6-56"))
	list.Add("field7", factory.New("p7-77"))
	list.Add("field8", factory.New("p8-88"))
}
