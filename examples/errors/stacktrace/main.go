package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mondegor/go-core/errors/runtime/stacktrace"
)

// main - пример формирования стека вызовов с использованием опций.
func main() {
	curPath, _ := os.Getwd()
	fmt.Println("RootPath: " + curPath)

	caller := stacktrace.NewCaller(
		stacktrace.WithDepth(16),
		stacktrace.WithFindBottomBoundFunc(
			stacktrace.FindBottomBound(
				[]string{
					"main.funcLevel2",
					// "main.funcLevel1",
				},
			),
		),
	)

	stackTrace := funcLevel1(caller)
	stackTraceIterator := stackTrace.Iterator()

	for {
		index, name, file, line := stackTraceIterator()
		if index < 0 {
			break
		}

		if index == 0 {
			fmt.Println("[StackTrace] ")
		}

		fmt.Print(strconv.Itoa(index+1) + ". " + name + "(): " + file + ":" + strconv.Itoa(line) + "\n")
	}
}

func funcLevel1(caller stacktrace.Caller) stacktrace.StackTrace {
	return funcLevel2(caller)
}

func funcLevel2(caller stacktrace.Caller) stacktrace.StackTrace {
	return caller.Call()
}
