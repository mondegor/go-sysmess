package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/errors"
	"github.com/mondegor/go-sysmess/errors/helper"
)

// main - пример использования errors.CustomError и errors.CustomListError.
func main() {
	userErrorProto := errors.NewUserProto("MyUserErrorWithParams", "my user error with '{Param1}' and '{Param2}'")
	systemErrorProto := errors.NewSystemProto("my system error")
	internalErrorProto := errors.NewInternalProto("my internal error")
	unknownError := errors.New("my error in form field")
	nilError := (error)(nil)

	errorList := errors.CustomListError{
		errors.WithCustomCode(userErrorProto.New("p1-1", "p1-2"), "fieldCode1"),
		errors.WithCustomCode(userErrorProto.Wrap(internalErrorProto.New("attr2-1", "value2-1"), "p2-1", "p2-2"), "fieldCode2"),
		errors.WithCustomCode(systemErrorProto.New("attr3-1", "value3-1"), "fieldCode3"),
		errors.WithCustomCode(internalErrorProto.New("attr4-1", "value4-1", "attr4-2", "value4-2"), "fieldCode4"),
		errors.WithCustomCode(unknownError, "fieldCode5"),
		errors.WithCustomCode(nilError, "fieldCode6"),
	}

	for i, item := range errorList {
		echoErrorInfo(item, i+1)
		fmt.Println("")
	}
}

func echoErrorInfo(err error, number int) {
	if e := (errors.CustomError)(nil); errors.As(err, &e) {
		fmt.Printf("error #%d:\n", number)
		fmt.Println("- CustomCode = " + e.CustomCode())
		fmt.Printf("- IsKindUser = %v\n", e.IsKindUser())
		fmt.Println("- MessageForUser = " + translateError(e.Unwrap()))
		fmt.Println("- MessageForLog = " + e.Error())
	}
}

func translateError(err error) string {
	message, args := helper.ExtractMessageForLocalization(err)

	if len(args) > 0 {
		return fmt.Sprintf("translateError()->%s + args: %+v", message, args)
	}

	return fmt.Sprintf("translateError()->%s", message)
}
