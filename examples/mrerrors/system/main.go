package main

import (
	"context"
	"fmt"

	"github.com/mondegor/go-sysmess/mrerrors"
	"github.com/mondegor/go-sysmess/mrlib/crypt"
	"github.com/mondegor/go-sysmess/mrmsg"
)

// main - пример system ошибки с параметром в сообщении, c уникальным ID, но без stack trace.
func main() {
	proto := createErrorProto()

	err := createErrLevel1(proto)
	echo(err)

	err = createErrLevel2(proto)
	echo(err)

	if proto.Is(err) {
		fmt.Println("Yes, it is system error, proto:", proto.Error())
	}
}

func createErrorProto() *mrerrors.ProtoError {
	return mrerrors.NewProto(
		"my system error with param %Param1% and param %Param2%",
		mrerrors.WithProtoKind(mrerrors.ErrorKindSystem),
		mrerrors.WithProtoCode("sys-code-001"),
		mrerrors.WithProtoArgsReplacer(func(message string) mrerrors.MessageReplacer {
			return mrmsg.NewMessageReplacer("%", "%", message)
		}),
		mrerrors.WithProtoOnCreated(func(ctx context.Context, err error) (instanceID string) {
			return crypt.GenerateInstanceID()
		}),
	)
}

func createErrLevel1(f *mrerrors.ProtoError) error {
	return f.New("MY-PARAM00001", 11111)
}

func createErrLevel2(f *mrerrors.ProtoError) error {
	return createErr2(f)
}

func createErr2(f *mrerrors.ProtoError) error {
	return f.New("MY-PARAM00002", 22222)
}

func echo(err error) {
	fmt.Println(err)

	if e, ok := err.(*mrerrors.InstantError); ok {
		fmt.Println("- Kind = " + e.Kind().String())
		fmt.Println("- Code = " + e.Code())
		fmt.Println("- ID = " + e.ID())
		fmt.Println("")
	}
}
