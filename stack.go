package goerror

import (
	"fmt"
	"runtime"
	"strings"
)

type Stack []uintptr

func (s Stack) Frame() []Frame {
	f := []Frame{}
	if len(s) == 0 {
		return f
	}
	frames := runtime.CallersFrames(s)
	for {
		frame, more := frames.Next()
		f = append(f, Frame{Name: frame.Function, File: frame.File, Line: frame.Line})
		if !more {
			break
		}
	}
	return f
}

func (s Stack) String() string {
	strs := []string{}
	frames := runtime.CallersFrames(s)
	for {
		frame, more := frames.Next()
		strs = append(strs, fmt.Sprintf("  at %s (%s:%d)", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}
	return strings.Join(strs, "\n")
}

type Frame struct {
	File string
	Name string
	Line int
}

var maxCallerDepth = 10
var enableRecordCaller = true

func callers(skip int, depth int) []uintptr {
	pcs := make([]uintptr, depth)
	n := runtime.Callers(skip, pcs)
	return pcs[0:n]
}

// Set global enable call stack record, default `true`; (Does not affect the `GetCaller` function)
func SetRecordCaller(enable bool) {
	enableRecordCaller = enable
}

// Set max call stack length, default `10`
func SetMaxCallerDepth(n int) int {
	maxCallerDepth = n
	return maxCallerDepth
}
