package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
)

// FactoryError - пример пользовательской ошибки.
var FactoryError = mrerr.NewFactory(
	"errMyErrorWithParams",
	mrerr.ErrorTypeInternal,
	"my error with '{{ .param1 }}' and '{{ .param2 }}'",
)

func main() {
	err := createErr()

	fmt.Println(err)
	fmt.Println("Code = " + err.Code())
	fmt.Println("Kind = " + err.Kind().String()) // 0 - mrerr.ErrorKindInternal
	fmt.Println("InstanceID = " + err.InstanceID())

	err = createErrLevel2()
	fmt.Println(err)
}

func createErr() *mrerr.AppError {
	return FactoryError.New("my-param1", 123456)
}

func createErrLevel2() *mrerr.AppError {
	return createErr2()
}

func createErr2() *mrerr.AppError {
	return FactoryError.New("my-param2", "with Caller 2")
}
