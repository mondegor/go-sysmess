package main

import (
	"errors"
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/features"
	"github.com/mondegor/go-sysmess/mrmsg"
)

func main() {
	mrerr.InitDefaultOptions(features.DefaultOptionsHandler())

	tr := Translator{}
	customError := mrerr.NewCustomError("formFieldEmail", errors.New("error in email"))
	fmt.Println(customError)

	proto := mrerr.NewProto(
		"errMyErrorWithParams",
		mrerr.ErrorKindUser,
		"my error with '{{ .param1 }}' and '{{ .param2 }}'",
	)

	list := mrerr.CustomErrors{
		customError,
		mrerr.NewCustomError("formFieldPhone1", proto.New("p1-11", "p2-11").WithAttr("my-attr1", "attr-value1")),
	}

	list = append(list, mrerr.NewCustomError("formField2", nil))
	list = append(list, mrerr.NewCustomError("formField3", errors.New("simple error")))
	list = append(list, mrerr.NewCustomError("formField4", proto.New("p1-44", "p2-44")))

	addSomeItems(&list)

	for i, item := range list {
		fmt.Printf("\nerror #%d:\n", i)
		echo(item, &tr)
	}
}

func addSomeItems(list *mrerr.CustomErrors) {
	proto := mrerr.NewProto(
		"errSomeItems",
		mrerr.ErrorKindSystem,
		"my error with '{{ .param1 }}'",
	)

	*list = append(*list, mrerr.NewCustomError("formField5", nil))
	*list = append(*list, mrerr.NewCustomError("formField6", proto.New("p1-66")))
}

func echo(err error, tr *Translator) {
	if e, ok := err.(*mrerr.CustomError); ok {
		fmt.Println("- CustomCode = " + e.CustomCode())
		fmt.Println("- ErrCode = " + e.Err().Code())
		fmt.Println("- Kind = " + e.Err().Kind().String())
		fmt.Println("- InstanceID = " + e.Err().InstanceID())
		fmt.Println("- MessageForUser = " + e.Err().Translate(tr).Reason)
		fmt.Println("- MessageForLog = " + e.Error())
	}
}

type (
	// Translator - Объект для работы с переводом ошибок на различные языки.
	Translator struct{}
)

func (t *Translator) HasErrorCode(code string) bool {
	return true
}

func (t *Translator) TranslateError(code, defaultMessage string, args ...mrmsg.NamedArg) mrmsg.ErrorMessage {
	argsAny := make([]any, 0, len(args))

	for i := range args {
		argsAny = append(argsAny, args[i])
	}

	return mrmsg.ErrorMessage{
		Reason: fmt.Sprintf(code+" replaced error for user", argsAny...),
	}
}
