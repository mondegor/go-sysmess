package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
)

func main() {
	customError := mrerr.NewCustomErrorMessage("fieldEmail", "error in email")
	fmt.Println(customError)

	factory := mrerr.NewFactory(
		"errMyErrorWithParams",
		mrerr.ErrorKindUser,
		"my error with '{{ .param1 }}' and '{{ .param2 }}'",
	).WithAttr("my-attr1", "attr-value1").WithAttr("my-attr2", "attr-value2")

	list := mrerr.CustomErrorList{
		customError,
		mrerr.NewCustomErrorAppError("fieldPhone", factory.New("123", "567")),
	}

	list = append(list, mrerr.NewCustomErrorAppError("field2", factory.New("p1-22", "p2-22")))
	list = append(list, mrerr.NewCustomError("field3", factory.New("p1-33", "p2-33")))
	list = append(list, mrerr.NewCustomError("field4", factory.New("p1-44", "p2-44")))
	list = append(list, mrerr.NewCustomError("field5", nil))

	addSomeItems(&list)

	for i, item := range list {
		fmt.Printf("\nitem %d:\n", i)
		fmt.Printf(
			"CustomCode=%s, ErrCode=%s, ErrKind=%v, Err={%s}\n",
			item.Code(),            // custom error code
			item.AppError().Code(), // internal error code of custom error
			item.AppError().Kind(), // 3 - mrerr.ErrorKindUser
			item.AppError(),
		)
	}
}

func addSomeItems(list *mrerr.CustomErrorList) {
	factory := mrerr.NewFactory(
		"errSomeItems",
		mrerr.ErrorKindSystem,
		"my error with '{{ .param1 }}'",
	)

	*list = append(*list, mrerr.NewCustomErrorAppError("field6", factory.New("p6-56")))
	*list = append(*list, mrerr.NewCustomError("field7", factory.New("p7-77")))
	*list = append(*list, mrerr.NewCustomError("field8", factory.New("p8-88")))
	*list = append(*list, mrerr.NewCustomErrorAppError("field9", nil))
}
