package grace

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type (
	err   struct{}
	Error interface {
		Stack(e *error)
		Recover(e *error)
		RecoverStack(e *error)
		Set(error, interface{}) error
		New(error, interface{}) error
	}
)

func Err() Error {
	return err{}
}

func (r err) Stack(e *error) {
	if r := recover(); r != nil {
		*e = errors.New(string(debug.Stack()))
	}
}

func (r err) Recover(e *error) {
	if r := recover(); r != nil {
		*e = errors.New(fmt.Sprintln(r))
	}
}

func (r err) RecoverStack(e *error) {
	if r := recover(); r != nil {
		*e = errors.New(fmt.Sprintln(r, string(debug.Stack())))
	}
}

func (r err) Set(e error, val interface{}) error {
	if e != nil {
		e = errors.New(fmt.Sprint(val))
	}
	return e
}

func (r err) New(e error, val interface{}) error {
	if e == nil {
		e = errors.New(fmt.Sprint(val))
	}
	return e
}
