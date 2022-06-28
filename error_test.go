package goerror

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err := func() error {
		return NewWithCode("Unknow Error !", -1)
	}()
	fmt.Println(err)
	v, ok := err.(Error)
	fmt.Println(v.Message(), ok)
	bstr, e := json.Marshal(map[string]interface{}{
		"err": err,
	})
	fmt.Println(string(bstr), e)
}

func TestPlainError(t *testing.T) {
	err := func() error {
		return NewPlain("Unknow Error !")
	}()
	fmt.Println(err)
	v, ok := err.(Error)
	fmt.Println(v.Message(), ok)
	bstr, e := json.Marshal(map[string]interface{}{
		"err": err,
	})
	fmt.Println(string(bstr), e)
}

func TestGetCaller(t *testing.T) {
	frames := GetCaller()
	fmt.Println(frames)
}

func BenchmarkError1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() error {
			return errors.New("Hello !")
		}()
	}
}

func BenchmarkError2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() error {
			return New("Hello !")
		}()
	}
}

func BenchmarkError3(b *testing.B) {
	SetRecordCaller(false) // close record call stack
	for i := 0; i < b.N; i++ {
		func() error {
			return New("Hello !")
		}()
	}
}

func BenchmarkError4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() error {
			return NewPlain("Hello !")
		}()
	}
}
