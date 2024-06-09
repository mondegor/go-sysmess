package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/factory"
)

// main - пример system ошибки с параметром в сообщении, c уникальным ID, но без stack trace.
func main() {
	proto := createErrorProto()

	err := createErrLevel1(proto)
	echo(err)

	err = createErrLevel2(proto)
	echo(err)

	if proto.Is(err) {
		fmt.Println("Yes, error with code:", proto.Code())
	}
}

func createErrorProto() *mrerr.ProtoAppError {
	return factory.NewProtoAppError(
		factory.Options{
			Code:            "errMySystemError",
			Kind:            mrerr.ErrorKindSystem,
			Message:         "my system error with param {{ .param1 }} and param {{ .param2 }}",
			WithIDGenerator: true,
			WithCaller:      false,
		},
	)
}

func createErrLevel1(f *mrerr.ProtoAppError) error {
	return f.New("MY-PARAM00001", 11111)
}

func createErrLevel2(f *mrerr.ProtoAppError) error {
	return createErr2(f)
}

func createErr2(f *mrerr.ProtoAppError) error {
	return f.New("MY-PARAM00002", 22222)
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
