package grace

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type (
	err struct {
		*logger
		*printer
		history []error
	}
	Error interface {
		Stack(e *error)
		Recover(e *error)
		Log(string) Error
		Print(PrintFunc) Error
		RecoverStack(e *error)
		Set(*error, interface{})
		New(error, interface{}) error
		getLog() *logger
		getPrint() *printer
	}
)

func Err() Error {
	return &err{}
}

func Log() Error {
	r := err{}
	r.logger = &logger{enable: true, err: &r, fname: fname}
	return &r
}

func Print() Error {
	r := err{}
	r.printer = &printer{
		enable: true,
		Err:    &r,
		fprint: func(e error) {
			fmt.Println(e.Error())
		},
	}
	return &r
}

func (err *err) Stack(e *error) {
	if r := recover(); r != nil {
		*e = errors.New(string(debug.Stack()))
		err.stamp(*e)
	}
}

func (err *err) Recover(e *error) {
	if r := recover(); r != nil {
		*e = errors.New(fmt.Sprintln(r))
		err.stamp(*e)
	}
}

func (err *err) RecoverStack(e *error) {
	if r := recover(); r != nil {
		*e = errors.New(fmt.Sprintln(r, string(debug.Stack())))
		err.stamp(*e)
	}
}

func (err *err) Log(file string) Error {
	err.logger = &logger{
		enable: true,
		err:    err,
		fname:  file,
	}
	return err
}

func (err *err) Print(fn PrintFunc) Error {
	err.printer = &printer{
		enable: true,
		Err:    err,
		fprint: fn,
	}
	return err
}

func (err *err) Set(e *error, val interface{}) {
	if e != nil {
		*e = errors.New(fmt.Sprint(val))
	}
}

func (err *err) New(e error, val interface{}) error {
	if e == nil {
		e = errors.New(fmt.Sprint(val))
	}
	return e
}

func (err *err) getLog() *logger {
	return err.logger
}

func (err *err) getPrint() *printer {
	return err.printer
}
