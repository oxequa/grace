package grace

import (
	"errors"
	"fmt"
)

type grace struct {
	errors   []error
	handlers []HandlerFunc
}

// Return error value
func (g *grace) Error() string {
	if len(g.errors) > 0 {
		return g.errors[len(g.errors)-1].Error()
	}
	return ""
}

// Empty check if error message is an empty string.
func Empty(e error) bool {
	if e != nil {
		if e.Error() == "" {
			return true
		}
		return false
	}
	return true
}

// New return an error created on the base of a given interface.
func New(v ...interface{}) error {
	g := &grace{}
	g.errors = append(g.errors, errors.New(fmt.Sprint(v...)))
	return g
}

// Equal check if two errors are equal.
func Equal(e1 error, e2 error) bool {
	if e1.Error() == e2.Error() {
		return true
	}
	return false
}

// Def if error is nil return a new error with the interface as error value.
func Def(e error, v ...interface{}) error {
	if Empty(e) {
		return New(v...)
	}
	return e
}

// Set if error is not nil return a new error with the interface as error value.
func Set(e error, v ...interface{}) error {
	if !Empty(e) {
		return New(v...)
	}
	return e
}

// Prefix an error with one or a list of custom value.
func Prefix(e error, p ...interface{}) error {
	switch err := e.(type) {
	case *grace:
		for i, elm := range err.errors {
			err.errors[i] = New(fmt.Sprint(p...), elm.Error())
		}
		return e
	}
	return New(fmt.Sprint(p...), e.Error())
}
