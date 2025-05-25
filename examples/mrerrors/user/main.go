package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerrors"
	"github.com/mondegor/go-sysmess/mrmsg"
)

// main - пример user ошибки с параметром в сообщении, без уникального ID и без stack trace.
func main() {
	proto := createErrorProto()

	err := createErrLevel1(proto)
	echo(err)

	if proto.Is(err) {
		fmt.Println("Yes, it is user error, proto:", proto.Error())
	}
}

func createErrorProto() *mrerrors.ProtoError {
	return mrerrors.NewProto(
		"my user error with param {{%param1%}} and param ###param2%%",
		mrerrors.WithProtoKind(mrerrors.ErrorKindUser),
		mrerrors.WithProtoArgsReplacer(func(message string) mrerrors.MessageReplacer {
			return mrmsg.NewMessageReplacer("{{%", "%}}", message)
		}),
		mrerrors.WithProtoCode("MyUserError"),
	)
}

func createErrLevel1(f *mrerrors.ProtoError) error {
	return f.New("MY-PARAM00001", 11112222)
}

func echo(err error) {
	fmt.Println(err)

	if e, ok := err.(*mrerrors.InstantError); ok {
		fmt.Println("- Kind = " + e.Kind().String())
		fmt.Println("- Code = " + e.Code())
		fmt.Println("- ID = " + e.ID()) // none
		fmt.Println("")
	}
}
