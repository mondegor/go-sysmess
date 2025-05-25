package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mondegor/go-sysmess/mrargs"
	"github.com/mondegor/go-sysmess/mrerr/generate"
	"github.com/mondegor/go-sysmess/mrerr/stacktrace"
	"github.com/mondegor/go-sysmess/mrerrors"
	"github.com/mondegor/go-sysmess/mrerrors/mapping"
	"github.com/mondegor/go-sysmess/mrmsg"
)

// main - пример internal ошибки c уникальным ID и со stack trace.
func main() {
	proto := createErrorProto()

	err := createErrLevel1(proto)
	echo(err)

	err = createErrLevel2(proto)
	echo(err)

	if proto.Is(err) {
		fmt.Println("Yes, it is internal error, proto:", proto.Error())
	}
}

func createErrorProto() *mrerrors.ProtoError {
	caller := stacktrace.New(
		stacktrace.WithDepth(5),
	)

	return mrerrors.NewProto(
		"my internal error with param '{param1}'",
		mrerrors.WithProtoKind(mrerrors.ErrorKindInternal),
		mrerrors.WithProtoCode("int-code-001"),
		mrerrors.WithProtoArgsReplacer(func(message string) mrerrors.MessageReplacer {
			return mrmsg.NewMessageReplacer("{", "}", message)
		}),
		mrerrors.WithProtoCaller(func() mrerrors.StackTracer {
			return caller.StackTrace()
		}),
		mrerrors.WithProtoOnCreated(func(ctx context.Context, err error) (instanceID string) {
			return generate.InstanceID()
		}),
	)
}

func createErrLevel1(f *mrerrors.ProtoError) error {
	return f.New(
		"MY-PARAM00001",
		"attr1", "attr-value1",
		12345,
		"attr3", mrargs.Group{
			"subattr3-1": "subval3-1",
			"subattr3-2": "subval3-2",
		},
		"attr4",
	).WithAttrs("attr-value4", "attr5-without-value")
}

func createErrLevel2(f *mrerrors.ProtoError) error {
	return createErr2(f)
}

func createErr2(f *mrerrors.ProtoError) error {
	return f.New("MY-PARAM00002")
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
