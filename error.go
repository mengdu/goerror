package goerror

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"sync"
)

type Stack struct {
	File string
	Line int
	Name string
}

type Error struct {
	stack   []Stack
	code    int
	message string
}

func (e Error) Error() string {
	strs := []string{fmt.Sprintf("Error(%d): %s", e.code, e.message)}
	for _, v := range e.stack {
		strs = append(strs, fmt.Sprintf("  at %s (%s:%d)", v.Name, v.File, v.Line))
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

func (e Error) Stack() []Stack {
	return e.stack
}

// New Error include call stack information
func New(message string) Error {
	e := Error{message: message}
	if enableRecordCaller {
		e.stack = GetCaller()
	}
	return e
}

// New Error with code and include call stack information
func NewWithCode(message string, code int) Error {
	e := Error{message: message, code: code}
	if enableRecordCaller {
		e.stack = GetCaller()
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
	e := Error{message: err.Error()}
	if enableRecordCaller {
		e.stack = GetCaller()
	}
	return e
}

// Wrap error with code and include call stack information
func WrapWithCode(err error, code int) error {
	e := Error{message: err.Error(), code: code}
	if enableRecordCaller {
		e.stack = GetCaller()
	}
	return e
}

var currentPackageName = ""
var maxCallerDepth = 10
var once sync.Once
var enableRecordCaller = true

// Get current func call stack information
func GetCaller() []Stack {
	once.Do(func() {
		_, file, _, _ := runtime.Caller(0)
		currentPackageName = file
	})

	arr := []Stack{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(1, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for {
		frame, more := frames.Next()
		if frame.File == currentPackageName {
			continue
		}
		arr = append(arr, Stack{
			File: frame.File,
			Line: frame.Line,
			Name: frame.Function,
		})
		if !more {
			break
		}
	}
	return arr
}

// Set global enable call stack record, default `true`; (Does not affect the `GetCaller` function)
func SetRecordCaller(enable bool) {
	enableRecordCaller = enable
}
