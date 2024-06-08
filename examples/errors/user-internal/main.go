package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/factory"
)

// main - пример user ошибки с вложенной internal ошибкой.
func main() {
	errInternal := createErrorInternalProto()
	errUser := createErrorUserProto()

	err := createErrInternal(errInternal)
	echo(err)

	err = createErrUser(errUser, err)
	echo(err)

	if errInternal.Is(err) {
		fmt.Println("Yes, internal error with code:", errInternal.Code())
	}

	if errUser.Is(err) {
		fmt.Println("Yes, user error with code:", errUser.Code())
	}
}

func createErrorInternalProto() *mrerr.AppErrorProto {
	return factory.NewAppErrorProto(
		factory.Options{
			Code:            "errMyInternalError",
			Kind:            mrerr.ErrorKindInternal,
			Message:         "my internal error with param {{ .param1 }}",
			WithIDGenerator: true,
			WithCaller:      true,
		},
	)
}

func createErrorUserProto() *mrerr.AppErrorProto {
	return factory.NewAppErrorProto(
		factory.Options{
			Code:            "errMyUserError",
			Kind:            mrerr.ErrorKindUser,
			Message:         "my user error with param {{ .param1 }} and param {{ .param2 }}",
			WithIDGenerator: false,
			WithCaller:      false,
		},
	)
}

func createErrUser(f *mrerr.AppErrorProto, err error) error {
	return f.Wrap(err, "My-param-for-user-error-00001", 11111)
}

func createErrInternal(f *mrerr.AppErrorProto) error {
	return createErrInternal2(f)
}

func createErrInternal2(f *mrerr.AppErrorProto) error {
	return f.New("My-param-for-Internal-error")
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
