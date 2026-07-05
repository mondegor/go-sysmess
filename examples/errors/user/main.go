package main

import (
	"fmt"

	"github.com/mondegor/go-core/errors"
	"github.com/mondegor/go-core/errors/helper"
	"github.com/mondegor/go-core/errors/kind"
	"github.com/mondegor/go-core/mrmsg"
)

type (
	userError interface {
		error

		Kind() kind.Enum
		Code() string
	}
)

// main - пример user ошибки с параметром в сообщении, без уникального ID и без stack trace.
func main() {
	userErrorProto1 := errors.NewUserProto("MyUserError", "my user error")
	userErrorProto2 := errors.NewUserProto("MyUserErrorWithParams", "my user error with params: {Param1} and {Param2}")

	echoErrorInfo(userErrorProto1, 1)
	echoErrorInfo(userErrorProto1.New(), 2)
	echoErrorInfo(userErrorProto2, 3)
	echoErrorInfo(userErrorProto2.New("MyParam4-1", "MyParam4-2"), 4)
	echoErrorInfo(userErrorProto1.Wrap(userErrorProto2.New("MyParam5-1", "MyParam5-2")), 5)
	echoErrorInfo(userErrorProto2.Wrap(userErrorProto1.New(), "MyParam6-1", "MyParam6-2"), 6)
}

func echoErrorInfo(err error, number int) {
	if e := (userError)(nil); errors.As(err, &e) {
		fmt.Printf("error #%d:\n", number)
		fmt.Println("- Kind = " + e.Kind().String())
		fmt.Println("- Code = " + e.Code())
		fmt.Println("- MessageForUser = " + translateError(e))
		fmt.Println("- MessageForLog = " + e.Error())

		fmt.Println("")
	}
}

func translateError(err error) string {
	message, args := helper.ExtractMessageForLocalization(err)

	p := mrmsg.NewMessageReplacer("{", "}", message)
	message, _ = p.Replace(args)

	return "translateError()->" + message
}
