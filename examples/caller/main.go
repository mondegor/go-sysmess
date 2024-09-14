package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mondegor/go-sysmess/mrcaller"
)

func main() {
	curPath, _ := os.Getwd()
	fmt.Println("RootPath: " + curPath)

	caller := mrcaller.New(
		mrcaller.WithDepth(16),
		mrcaller.WithShowFuncName(true),
		mrcaller.WithFilterStackTrace(
			mrcaller.FilterStackTraceTrimUpper(
				[]string{
					"main.funcLevel2",
					// "main.funcLevel1",
				},
			),
		),
	)

	stackTrace := funcLevel1(caller)

	for i := 0; i < stackTrace.Count(); i++ {
		name, file, line := stackTrace.Item(i)

		if i == 0 {
			fmt.Println("[CallStack] ")
		}

		fmt.Print(strconv.Itoa(i+1) + ". " + name + "(): " + file + ":" + strconv.Itoa(line) + "\n")
	}
}

func funcLevel1(caller *mrcaller.Caller) *mrcaller.StackTrace {
	return funcLevel2(caller)
}

func funcLevel2(caller *mrcaller.Caller) *mrcaller.StackTrace {
	return caller.StackTrace()
}
