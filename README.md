# goerror

Error contains call stack information.

> !Note that there will be some performance loss in obtaining the call stack information.
> Please using `goerror.SetRecordCaller(false)` in a production environment if you have high performance requirements.

[Docs](https://pkg.go.dev/github.com/mengdu/goerror#section-documentation)

```sh
go get github.com/mengdu/goerror
```

## Usage

```go
package main

import (
  "encoding/json"
  "fmt"

  "github.com/mengdu/goerror"
)

func main() {
  err := func() error {
    return goerror.New("Unknow Error!")
  }()
  fmt.Println(err)
  // Error(0): Unknow Error!
  //   at main.main.func1 (/path-to-project/main.go:12)
  //   at main.main (/path-to-project/main.go:13)
  //   at runtime.main (/usr/local/go/src/runtime/proc.go:250)
  //   at runtime.goexit (/usr/local/go/src/runtime/asm_amd64.s:1571)

  v, ok := err.(goerror.Error)
  fmt.Println(v.Message(), v.Code(), ok) // Unknow Error! 0 true

  bstr, _ := json.Marshal(map[string]interface{}{
    "err": err,
  })
  fmt.Println(string(bstr)) // {"err":{"code":0,"message":"Unknow Error!"}}
}
```
