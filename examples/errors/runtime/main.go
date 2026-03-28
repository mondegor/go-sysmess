package main

import (
	"fmt"
	"strings"

	"github.com/mondegor/go-sysmess/errors"
	"github.com/mondegor/go-sysmess/errors/kind"
	"github.com/mondegor/go-sysmess/errors/runtime/stacktrace"
	"github.com/mondegor/go-sysmess/util/conv"
	"github.com/mondegor/go-sysmess/wire"
)

type (
	runtimeError interface {
		error

		Kind() kind.Enum
		Hint() any
	}

	errorHint interface {
		ErrorID() string
		StackTraceIterator() func() (index int, name, file string, line int)
	}
)

// main - пример internal ошибки c уникальным ID и со stack trace.
func main() {
	wire.InitErrors(
		wire.ErrorConfig{
			HasCaller:      true,
			CallerDepth:    3,
			CallerShowFunc: true,
			CallerUpperBounds: []string{
				"github.com/mondegor/go-sysmess/errors/runtime.(*proto).New",
			},
		},
	)

	internalErrorProto := errors.NewInternalProto("my internal error")
	systemErrorProto := errors.NewSystemProto("my system error")

	execute(internalErrorProto, 1)
	execute(systemErrorProto, 3)
}

func execute(err errors.RuntimeProtoError, numberStart int) {
	err1 := createErrLevel1(err)
	echoErrorInfo(err1, numberStart)

	err2 := createErrLevel2_step1(err)
	echoErrorInfo(err2, numberStart+1)

	if errors.Is(err1, err) && errors.Is(err2, err) {
		fmt.Printf("Yes, #%d and #%d are errors, original proto error: %s\n", numberStart, numberStart+1, err.Error())
	}

	fmt.Println("")
}

func createErrLevel1(err errors.RuntimeProtoError) error {
	return err.New(
		"attr1", "value1",
		"attr2", conv.Group{
			"subattr2-1": "subval2-1",
			"subattr2-2": "subval2-2",
		},
		"attr3", 3,
	)
}

func createErrLevel2_step1(err errors.RuntimeProtoError) error {
	return createErrLevel2_step2(err)
}

func createErrLevel2_step2(err errors.RuntimeProtoError) error {
	return err.New("attr4", "value4")
}

func echoErrorInfo(err error, number int) {
	if e := (runtimeError)(nil); errors.As(err, &e) {
		fmt.Printf("error #%d:\n", number)
		fmt.Println("- Kind = " + e.Kind().String())
		fmt.Println("- MessageForLog = " + e.Error())

		if bag, ok := e.Hint().(errorHint); ok {
			fmt.Println("- ErrorID = " + bag.ErrorID())

			stackList := stacktrace.ToStrings(bag.StackTraceIterator())
			if len(stackList) > 0 {
				fmt.Println("- StackTrace = \n  - " + strings.Join(stackList, "\n  - "))
			}
		}

		fmt.Println("")
	}
}
