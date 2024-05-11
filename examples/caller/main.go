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
		mrcaller.CallerDepth(3),
		// mrcaller.CallerUseShortPath(true),
		// mrcaller.CallerRootPath(curPath),
	)

	callStack := funcLevel1(caller)

	for iter := callStack.NewIterator(); ; {
		if n, item := iter.Next(); n > 0 {
			if n == 1 {
				fmt.Println("[CallStack] ")
			}

			fmt.Print(strconv.Itoa(n) + ". " + item.Name() + "(): " + item.File() + ":" + strconv.Itoa(item.Line()) + "\n")
		} else {
			break
		}
	}
}

func funcLevel1(caller *mrcaller.Caller) mrcaller.CallStack {
	return funcLevel2(caller)
}

func funcLevel2(caller *mrcaller.Caller) mrcaller.CallStack {
	return caller.CallStack(0)
}
