package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mondegor/go-sysmess/mrerr"
	"github.com/mondegor/go-sysmess/mrerrors"
	"github.com/mondegor/go-sysmess/mrerrors/mapping"
)

// main - пример user ошибки с вложенной internal ошибкой.
func main() {
	mrerr.InitDefaultOptions(nil)

	errInternal := createErrorInternalProto()
	errUser := createErrorUserProto()

	err := createErrInternal(errInternal)
	echo(err)

	err = createErrUser(errUser, err)
	echo(err)

	if errors.Is(err, errInternal) {
		fmt.Println("Yes, it has wrapped internal error, proto:", errInternal.Error())
	}

	if errUser.Is(err) {
		fmt.Println("Yes, it is user error, proto:", errUser.Error())
	}
}

func createErrorInternalProto() *mrerrors.ProtoError {
	return mrerr.NewKindInternal(
		"my internal error !!!",
		mrerr.WithDefaultCaller(),
		mrerr.WithDefaultOnCreated(),
	)
}

func createErrorUserProto() *mrerrors.ProtoError {
	return mrerr.NewKindUser(
		"MyUserError",
		"my user error with param '{Param1}' and param '{Param2}'",
		mrerr.WithDefaultArgsReplacer(),
	)
}

func createErrUser(f *mrerrors.ProtoError, err error) error {
	return f.Wrap(err, "MY-PARAM-FOR-USER-ERROR-00001", 11112222)
}

func createErrInternal(f *mrerrors.ProtoError) error {
	return createErrInternal2(f)
}

func createErrInternal2(f *mrerrors.ProtoError) error {
	return f.Wrap(errors.New("WRAPPED ERROR"))
}

func echo(err error) {
	fmt.Println(err)

	if e, ok := err.(*mrerrors.InstantError); ok {
		fmt.Println("- Kind = " + e.Kind().String())
		fmt.Println("- Code = " + e.Code())
		fmt.Println("- ID = " + e.ID())
		fmt.Println("- StackTrace = \n  - " + strings.Join(mapping.StackTraceToStrings(e.StackTrace()), "\n  - "))
		fmt.Println("")
	}
}
