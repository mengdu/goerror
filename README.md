# goerror

Error contains call stack information.

> !! Developing...

```go
func main(t *testing.T) {
  err := New("Unknow Error!", -1)
  fmt.Println(err)
  // Error(-1): Unknow Error!
  //   at github.com/mengdu/goerror.TestError (/path-to-pkg/error_test.go:10)
  //   at testing.tRunner (/usr/local/go/src/testing/testing.go:1439)
  //   at runtime.goexit (/usr/local/go/src/runtime/asm_amd64.s:1571)

  bstr, e := json.Marshal(map[string]interface{}{
    "err": err,
  })
  fmt.Println(string(bstr), e)
  // {"err":{"code":-1,"message":"Unknow Error!"}} <nil>
}
```
