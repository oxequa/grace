package grace

import (
	"errors"
	"fmt"
	"runtime/debug"
)

type (
	printer struct{}
	Printer interface {
		Stack(PrintFunc, *error)
		Recover(PrintFunc, *error)
		Err(PrintFunc, error) error
		RecoverStack(PrintFunc, *error)
		Set(PrintFunc, error, interface{}) error
		New(PrintFunc, error, interface{}) error
	}
	PrintFunc func(e error)
)

func Print() Printer {
	return printer{}
}

func stamp(f PrintFunc, e error) error {
	if f != nil && e != nil {
		f(e)
		return e
	}
	return nil
}

func (p printer) Stack(f PrintFunc, e *error) {
	if r := recover(); r != nil {
		*e = errors.New(string(debug.Stack()))
		stamp(f, *e)
	}
}

func (p printer) Recover(f PrintFunc, e *error) {
	if r := recover(); r != nil {
		*e = errors.New(fmt.Sprint(r))
		stamp(f, *e)
	}
}

func (p printer) Err(f PrintFunc, e error) error {
	return stamp(f, e)
}

func (p printer) RecoverStack(f PrintFunc, e *error) {
	if r := recover(); r != nil {
		*e = errors.New(fmt.Sprintln(r, string(debug.Stack())))
		stamp(f, *e)
	}
}

func (p printer) Set(f PrintFunc, e error, val interface{}) error {
	if e != nil {
		e = errors.New(fmt.Sprint(val))
		return stamp(f, e)
	}
	return nil
}

func (p printer) New(f PrintFunc, e error, val interface{}) error {
	if e == nil {
		e = errors.New(fmt.Sprint(val))
		return stamp(f, e)
	}
	return nil
}
