package main

import (
	"errors"
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerrors"
)

// main - пример использования mrerr.CustomError и mrerr.CustomErrors.
func main() {
	mrerr.InitDefaultOptions(mrerr.DefaultOptionsHandler())

	customError := mrerr.NewCustomError("formFieldEmail", errors.New("error in email"))
	fmt.Println(customError)

	kindUserProto := mrerr.NewKindUser(
		"MyErrorWithParams",
		"my error with '{{param1}}' and '{{param2}}'",
		mrerr.WithDefaultArgsReplacer(),
	)

	list := mrerr.CustomErrors{
		customError,
		mrerr.NewCustomError("formFieldPhone1", kindUserProto.New("p1-11", "p2-11").WithAttr("my-attr1", "attr-value1")),
	}

	list = addSomeItems(list)

	for i, item := range list {
		fmt.Printf("\nerror #%d:\n", i)
		echo(item)
	}
}

func addSomeItems(list mrerr.CustomErrors) mrerr.CustomErrors {
	kindSystemProto := mrerr.NewKindSystem("my system error with '{{param1}}'")
	kindInternalProto := mrerr.NewKindInternal("my internal error with '{{param1}}'")

	list = append(list, mrerr.NewCustomError("formField2", kindSystemProto.New("p1-222")))
	list = append(list, mrerr.NewCustomError("formField3", kindInternalProto.New("p1-333")))

	return list
}

func echo(err error) {
	if e, ok := err.(*mrerr.CustomError); ok {
		fmt.Println("- CustomCode = " + e.CustomCode())
		fmt.Println("- Kind = " + e.Err().Kind().String())
		fmt.Println("- Code = " + e.Err().Code())
		fmt.Println("- ID = " + e.Err().ID())
		fmt.Println("- MessageForUser = " + translateError(e.Err()))
		fmt.Println("- MessageForLog = " + e.Error())
	}
}

func translateError(e *mrerrors.InstantError) string {
	switch {
	case e.Kind() == mrerrors.ErrorKindUser:
		return fmt.Sprintf("%s + args: %+v", e.Message(), e.Args())

	case e.Kind() == mrerrors.ErrorKindSystem:
		return "system error"

	default:
		return "internal error"
	}
}
