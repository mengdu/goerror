package main

import (
	"fmt"

	"github.com/mengdu/goerror"
)

func main() {
	err := goerror.New("Hello Error")
	fmt.Println(err)
}
