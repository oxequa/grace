package grace

import "testing"

func TestMulti(t *testing.T) {
	result := Multi(error(nil))
	if result == nil {
		t.Fatal("Unexpected error")
	}
	result = Multi(New("test"))
	if result == nil {
		t.Fatal("Unexpected error")
	}
}

func TestMultiple_Add(t *testing.T) {
	err := New("test")
	err = Multi(err).Add(New("test"), New("test1"))
	if l := len(Multi(err).Get()); l != 3 {
		t.Fatal("Unexpected lenght", l)
	}
}

func TestMultiple_Get(t *testing.T) {
	err := &grace{
		errors: []error{New("test"), New("test")},
	}
	if l := len(Multi(err).Get()); l != 2 {
		t.Fatal("Unexpected lenght", l)
	}
}
