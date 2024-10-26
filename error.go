package goerror

import (
	"encoding/json"
	"fmt"
	"strings"
)

var json_with_stack = true

func SetJsonWithStack(b bool) {
	json_with_stack = b
}

type Error struct {
	stack   Stack
	code    int
	message string
	prev    *Error
}

// Error implement error interface
func (e Error) Error() string {
	strs := []string{fmt.Sprintf("Error(%d): %s", e.code, e.message)}
	if enableRecordCaller && len(e.stack) > 0 {
		strs = append(strs, e.stack.String())
	}
	i := 1
	for p := e.prev; p != nil; p = p.prev {
		pad := strings.Repeat(" ", i*2)
		strs = append(strs, pad+fmt.Sprintf("Error(%d): %s", p.code, p.message))
		if enableRecordCaller && len(p.stack) > 0 {
			lines := strings.Split(p.stack.String(), "\n")
			for _, line := range lines {
				strs = append(strs, pad+line)
			}
		}
		i++
	}
	return strings.Join(strs, "\n")
}

// MarshalJSON implement json.Marshaler interface
func (e Error) MarshalJSON() ([]byte, error) {
	v := map[string]interface{}{
		"message": e.message,
		"code":    e.code,
	}
	if json_with_stack && enableRecordCaller {
		v["stack"] = fmt.Sprintf("Error(%d): %s", e.code, e.message) + "\n" + e.stack.String()
	}
	return json.Marshal(v)
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
func New(message string, code ...int) Error {
	e := Error{message: message}
	if len(code) > 0 {
		e.code = code[0]
	}
	if enableRecordCaller {
		e.stack = callers(3, maxCallerDepth)
	}
	return e
}

// Wrap error to include call stack information
func From(err error, code ...int) error {
	if err == nil {
		return nil
	}
	var e Error
	if err2, ok := err.(Error); ok {
		e = Error{code: err2.code, message: err2.message, prev: &err2}
	} else {
		e = Error{message: err.Error()}
	}
	if len(code) > 0 {
		e.code = code[0]
	}
	if enableRecordCaller {
		e.stack = callers(3, maxCallerDepth)
	}
	return e
}

func Errorf(format string, a ...interface{}) Error {
	e := Error{message: fmt.Errorf(format, a...).Error()}
	if enableRecordCaller {
		e.stack = callers(3, maxCallerDepth)
	}
	return e
}
