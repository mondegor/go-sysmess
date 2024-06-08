package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/factory"
)

// main - пример internal ошибки c уникальным ID и со stack trace.
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

func createErrorProto() *mrerr.AppErrorProto {
	return factory.NewAppErrorProto(
		factory.Options{
			Code:            "errMyInternalError",
			Kind:            mrerr.ErrorKindInternal,
			Message:         "my internal error",
			WithIDGenerator: true,
			WithCaller:      true,
		},
	)
}

func createErrLevel1(f *mrerr.AppErrorProto) error {
	return f.New()
}

func createErrLevel2(f *mrerr.AppErrorProto) error {
	return createErr2(f)
}

func createErr2(f *mrerr.AppErrorProto) error {
	return f.New()
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
