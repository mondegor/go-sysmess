package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mondegor/go-sysmess/mrerr/stacktrace"
)

// main - пример формирования стека вызовов с использованием опций.
func main() {
	curPath, _ := os.Getwd()
	fmt.Println("RootPath: " + curPath)

	caller := stacktrace.New(
		stacktrace.WithDepth(16),
		stacktrace.WithStackTraceFilter(
			stacktrace.TrimUpperFilter(
				[]string{
					"main.funcLevel2",
					// "main.funcLevel1",
				},
			),
		),
	)

	stackTrace := funcLevel1(caller)

	for i := 0; i < stackTrace.Count(); i++ {
		name, file, line := stackTrace.Source(i)

		if i == 0 {
			fmt.Println("[StackTrace] ")
		}

		fmt.Print(strconv.Itoa(i+1) + ". " + name + "(): " + file + ":" + strconv.Itoa(line) + "\n")
	}
}

func funcLevel1(caller *stacktrace.Caller) *stacktrace.StackTrace {
	return funcLevel2(caller)
}

func funcLevel2(caller *stacktrace.Caller) *stacktrace.StackTrace {
	return caller.StackTrace()
}
