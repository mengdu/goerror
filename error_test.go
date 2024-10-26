package goerror

import (
	"errors"
	"testing"
)

func BenchmarkStdError(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			func() error {
				return errors.New("Hello !")
			}()
		}
	})
}

func BenchmarkGoError(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			func() error {
				return New("Hello !")
			}()
		}
	})
}

func BenchmarkGoErrorNoCaller(b *testing.B) {
	SetRecordCaller(false) // close record call stack
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			func() error {
				return New("Hello !")
			}()
		}
	})
}
