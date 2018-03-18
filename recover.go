package grace

import (
	"runtime/debug"
)

type Recovery struct {
	err   *error
	value string
	stack []byte
}

// Error set the error pointer to error message only
func (rc *Recovery) Error() {
	if r := recover(); r != nil {
		switch err := (*rc.err).(type) {
		case *grace:
			g := &grace{
				handlers: err.handlers,
				errors:   []error{New(r)},
			}
			err.errors = append(err.errors, Handler(g).Loop())
		default:
			*rc.err = &grace{errors: []error{New(r)}}
		}
	}
}

// Stack set the error pointer to stack trace only
func (rc *Recovery) Stack() {
	if r := recover(); r != nil {
		switch err := (*rc.err).(type) {
		case *grace:
			g := &grace{
				handlers: err.handlers,
				errors:   []error{New(string(debug.Stack()))},
			}
			err.errors = append(err.errors, Handler(g).Loop())
		default:
			*rc.err = &grace{errors: []error{New(string(debug.Stack()))}}
		}
	}
}

// Recover set an error pointer to recover error with stack trace.
func Recover(e *error) *Recovery {
	if r := recover(); r != nil {
		switch err := (*e).(type) {
		case *grace:
			g := &grace{
				handlers: err.handlers,
				errors:   []error{New(r, string(debug.Stack()))},
			}
			err.errors = append(err.errors, Handler(g).Loop())
		default:
			*e = &grace{errors: []error{New(r, string(debug.Stack()))}}
		}
	}
	return &Recovery{
		err: e,
	}
}
