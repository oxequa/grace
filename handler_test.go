package grace

import (
	"bytes"
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	result := Handler(error(nil))
	if result == nil {
		t.Fatal("Unexpected error")
	}
	result = Handler(New("test"))
	if result == nil {
		t.Fatal("Unexpected error")
	}
}

func TestHandlers_Add(t *testing.T) {
	err := New("test")
	err = Handler(err).Add(func(e error) error {
		return e
	})
	if len(Handler(err).Get()) != 1 {
		t.Fatal("Unexpected num of handlers")
	}
	err = Handler(err).Add(func(e error) error {
		return e
	}, func(e error) error {
		return e
	})
	if len(Handler(err).Get()) != 3 {
		t.Fatal("Unexpected num of handlers")
	}
}

func TestHandlers_Get(t *testing.T) {
	err := &grace{
		handlers: []HandlerFunc{
			func(e error) error {
				return e
			},
		},
	}
	if len(Handler(err).Get()) != 1 {
		t.Fatal("Unexpected num of handlers")
	}
}

func TestHandlers_Loop(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	err := &grace{
		handlers: []HandlerFunc{
			func(e error) error {
				log.Println(e)
				return e
			},
		},
	}
	Handler(err).Loop()
	if len(buf.Bytes()) != 0 {
		t.Fatal("Unexpected error")
	}

	err = &grace{
		handlers: []HandlerFunc{
			func(e error) error {
				log.Println(e)
				return e
			},
		},
		errors: []error{
			New("test"),
		},
	}
	Handler(err).Loop()
	if len(buf.Bytes()) <= 0 {
		t.Fatal("Unexpected error")
	}
}
