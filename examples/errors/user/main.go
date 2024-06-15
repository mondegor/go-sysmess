package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/mrerrfactory"
)

// main - пример user ошибки с параметром в сообщении, без уникального ID и без stack trace.
func main() {
	proto := createErrorProto()

	err := createErrLevel1(proto)
	echo(err)

	if proto.Is(err) {
		fmt.Println("Yes, error with code:", proto.Code())
	}
}

func createErrorProto() *mrerr.ProtoAppError {
	return mrerrfactory.NewProtoAppErrorByDefault(
		"errMyUserError",
		mrerr.ErrorKindUser,
		"my user error with param {{ .param1 }} and param {{ .param2 }}",
	)
}

func createErrLevel1(f *mrerr.ProtoAppError) error {
	return f.New("MY-PARAM00001", 11111)
}

func echo(err error) {
	fmt.Println(err)

	if e, ok := err.(*mrerr.AppError); ok {
		fmt.Println("- Code = " + e.Code())
		fmt.Println("- Kind = " + e.Kind().String())
		fmt.Println("- InstanceID = " + e.InstanceID())
		fmt.Println("")
	}
}
