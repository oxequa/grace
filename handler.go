package grace

type Handlers struct {
	err error
}

type HandlerFunc func(error) error

func Handler(e error) *Handlers {
	h := Handlers{e}
	if e == nil {
		h.err = &grace{}
	}
	return &h
}

// Loop a list of handlers on an error or multiple error.
func (h *Handlers) Loop() error {
	switch err := h.err.(type) {
	case *grace:
		for _, elmE := range err.errors {
			for _, elmH := range err.handlers {
				elmE = elmH(elmE)
			}
		}
	}
	return h.err
}

// Get a list of existing handler func
func (h *Handlers) Get() []HandlerFunc {
	switch err := h.err.(type) {
	case *grace:
		return err.handlers
	}
	return nil
}

// Add a list of handlers to a new error.
func (h *Handlers) Add(f ...HandlerFunc) error {
	switch err := h.err.(type) {
	case *grace:
		err.handlers = append(err.handlers, f...)
		return h.err
	default:
		g := grace{}
		if h.err != nil {
			g.errors = append(g.errors, h.err)
		}
		g.handlers = append(g.handlers, f...)
		return &g
	}
}
