package grace

type Multiple struct {
	err *error
}

func Multi(e error) *Multiple {
	m := Multiple{&e}
	if e == nil {
		*m.err = &grace{}
	}
	return &m
}

// Errors return a list of errors if exist.
func (m *Multiple) Get() []error {
	if m.err != nil {
		switch err := (*m.err).(type) {
		case *grace:
			return err.errors
		default:
			return []error{err}
		}
	}
	return []error{}
}

// Append multiple errors to only one.
func (m *Multiple) Add(elems ...error) error {
	if m.err != nil {
		switch err := (*m.err).(type) {
		case *grace:
			err.errors = append(err.errors, elems...)
			return *m.err
		default:
			new := grace{}
			new.errors = append(new.errors, elems...)
			return &new
		}
	}
	return nil
}
