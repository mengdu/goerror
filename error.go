package goerror

import (
	"encoding/json"
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

type Frame struct {
	File string
	Name string
	Line int
}

type Error struct {
	stack   Stack
	code    int
	message string
}

func (e Error) Error() string {
	strs := []string{fmt.Sprintf("Error(%d): %s", e.code, e.message)}
	if len(e.stack) > 0 {
		frames := runtime.CallersFrames(e.stack)
		for {
			frame, more := frames.Next()
			strs = append(strs, fmt.Sprintf("  at %s (%s:%d)", frame.Function, frame.File, frame.Line))
			if !more {
				break
			}
		}
	}
	return strings.Join(strs, "\n")
}

func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"message": e.message,
		"code":    e.code,
	})
}

func (e Error) Message() string {
	return e.message
}

func (e Error) Code() int {
	return e.code
}

func (e Error) Stack() []Frame {
	return e.stack.Frame()
}

// New Error include call stack information
func New(message string) Error {
	e := Error{message: message}
	if enableRecordCaller {
		e.stack = callers(3, maxCallerDepth)
	}
	return e
}

// New Error with code and include call stack information
func NewWithCode(message string, code int) Error {
	e := Error{message: message, code: code}
	if enableRecordCaller {
		e.stack = callers(3, maxCallerDepth)
	}
	return e
}

// New Error not include call stack information
func NewPlain(message string) Error {
	e := Error{message: message}
	return e
}

// New Error with code and not include call stack information
func NewPlainWithCode(message string, code int) Error {
	e := Error{message: message, code: code}
	return e
}

// Wrap error to include call stack information
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	e := Error{message: err.Error()}
	if enableRecordCaller {
		e.stack = callers(3, maxCallerDepth)
	}
	return e
}

// Wrap error with code and include call stack information
func WrapWithCode(err error, code int) error {
	e := Error{message: err.Error(), code: code}
	if enableRecordCaller {
		e.stack = callers(3, maxCallerDepth)
	}
	return e
}

var maxCallerDepth = 10
var enableRecordCaller = true

func callers(skip int, depth int) []uintptr {
	pcs := make([]uintptr, depth)
	n := runtime.Callers(skip, pcs)
	return pcs[0:n]
}

// Get current func call stack information
func GetCaller() []Frame {
	var psc Stack = callers(3, maxCallerDepth)
	return psc.Frame()
}

// Set global enable call stack record, default `true`; (Does not affect the `GetCaller` function)
func SetRecordCaller(enable bool) {
	enableRecordCaller = enable
}

// Set call stack length
func SetMaxCallerDepth(n int) int {
	maxCallerDepth = n
	return maxCallerDepth
}
