package main

import (
	"errors"
	"fmt"

	"github.com/mengdu/goerror"
)

var (
	ERR_UNKNOWN  = goerror.New("Unknown", -1)
	ERR_PARAMS   = goerror.New("Params error", 1)
	ERR_NO_LOGIN = goerror.New("No Login", 403)
)

func main() {
	// goerror.SetJsonWithStack(false)
	// goerror.SetRecordCaller(false)
	err := goerror.New("Hello Error")
	fmt.Println(err)
	fmt.Println(ERR_UNKNOWN)
	buf, err2 := ERR_UNKNOWN.MarshalJSON()
	fmt.Println(err2, string(buf))
	fmt.Println(goerror.From(err, 1))
	fmt.Println(goerror.From(goerror.From(err, 1), 2))
	fmt.Println(goerror.Errorf("new error: %s", err.Message()))
	fmt.Println(goerror.Errorf("new error: %s", errors.New("test error")))
}
