package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerr/mrerrfactory"
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

func createErrorInternalProto() *mrerr.ProtoAppError {
	return mrerrfactory.NewProtoAppErrorByDefault(
		"errMyInternalError",
		mrerr.ErrorKindInternal,
		"my internal error with param {{ .param1 }}",
	)
}

func createErrorUserProto() *mrerr.ProtoAppError {
	return mrerrfactory.NewProtoAppErrorByDefault(
		"errMyUserError",
		mrerr.ErrorKindUser,
		"my user error with param {{ .param1 }} and param {{ .param2 }}",
	)
}

func createErrUser(f *mrerr.ProtoAppError, err error) error {
	return f.Wrap(err, "My-param-for-user-error-00001", 11111)
}

func createErrInternal(f *mrerr.ProtoAppError) error {
	return createErrInternal2(f)
}

func createErrInternal2(f *mrerr.ProtoAppError) error {
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
