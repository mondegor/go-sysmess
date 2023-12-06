package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
)

func main() {
	fieldError := mrerr.NewFieldErrorMessage("fieldEmail", "error in email")
	fmt.Println(fieldError)

	factory := mrerr.NewFactory(
		"errMyErrorWithParams",
		mrerr.ErrorKindUser,
		"my error with '{{ .param1 }}' and '{{ .param2 }}'",
	)

	list := mrerr.FieldErrorList{
		fieldError,
		mrerr.NewFieldErrorAppError("fieldPhone", factory.New("123", "567")),
	}

	list = append(list, mrerr.NewFieldErrorAppError("field2", factory.New("p1-22", "p2-22")))
	list = append(list, mrerr.NewFieldError("field3", factory.New("p1-33", "p2-33")))
	list = append(list, mrerr.NewFieldError("field4", factory.New("p1-44", "p2-44")))
	list = append(list, mrerr.NewFieldError("field5", nil))

	addSomeItems(&list)

	for i, item := range list {
		fmt.Printf("\nitem %d:\n", i)
		fmt.Printf(
			"FieldID=%s, ErrID=%s, ErrKind=%v, Err={%s}\n",
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

	*list = append(*list, mrerr.NewFieldErrorAppError("field6", factory.New("p6-56")))
	*list = append(*list, mrerr.NewFieldError("field7", factory.New("p7-77")))
	*list = append(*list, mrerr.NewFieldError("field8", factory.New("p8-88")))
	*list = append(*list, mrerr.NewFieldErrorAppError("field9", nil))
}
