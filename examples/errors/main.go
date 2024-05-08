package main

import (
	"fmt"

	"github.com/mondegor/go-sysmess/mrerr"
)

var FactoryError = mrerr.NewFactoryWithCaller(
	"errMyErrorWithParams",
	mrerr.ErrorKindInternal,
	"my error with '{{ .param1 }}' and '{{ .param2 }}'",
)

func main() {
	err := createErr()

	fmt.Println(err)
	fmt.Println("trace ID = " + err.TraceID())
	fmt.Println("Code = " + err.Code())
	fmt.Println("Kind = " + err.Kind().String()) // 0 - mrerr.ErrorKindInternal

	err = createErrLevel2()
	fmt.Println(err)

	err = createErrDisabledCaller()
	fmt.Println(err)
}

func createErr() *mrerr.AppError {
	return FactoryError.New("my-param1", 123456)
}

func createErrLevel2() *mrerr.AppError {
	return createErr2()
}

func createErr2() *mrerr.AppError {
	return FactoryError.WithCaller(2).New("my-param2", "with Caller 2")
}

func createErrDisabledCaller() *mrerr.AppError {
	return FactoryError.DisableCaller().New("my-param3", "with DisabledCaller")
}
