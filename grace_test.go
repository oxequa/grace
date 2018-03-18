package grace

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

type test struct {
	s string
	i int
	t time.Duration
}

func TestDef(t *testing.T) {
	err := error(nil)
	result := Def(err, "test")
	if result.Error() != "test" {
		t.Fatal("Unexpected error")
	}
	err = New("")
	result = Def(err, "test")
	if result.Error() != "test" {
		t.Fatal("Unexpected error")
	}
	err = New("ok")
	result = Def(err, "test")
	if result.Error() == "test" {
		t.Fatal("Unexpected error")
	}
}

func TestSet(t *testing.T) {
	err := error(nil)
	result := Set(err, "test")
	if result != nil {
		t.Fatal("Unexpected error")
	}
	err = New("")
	result = Set(err, "test")
	if result.Error() == "test" {
		t.Fatal("Unexpected error")
	}
	err = New("test1")
	result = Set(err, "test")
	if result.Error() != "test" {
		t.Fatal("Unexpected error")
	}
}

func TestNew(t *testing.T) {
	// string
	s := "test"
	err := New(s)
	if err.Error() != s {
		t.Fatal("Unexpected error", err.Error(), "instead", s)
	}
	// int
	i := 10
	err = New(i)
	if err.Error() != strconv.Itoa(i) {
		t.Fatal("Unexpected error", err.Error(), "instead", strconv.Itoa(i))
	}
	// interface
	str := test{
		s: "test",
		i: 10,
		t: time.Duration(1),
	}
	int := 100
	err = New(int, str)
	if !strings.Contains(err.Error(), strconv.Itoa(i)) {
		t.Fatal("Unexpected error", err.Error(), "should contains", strconv.Itoa(i))
	}
	if !strings.Contains(err.Error(), s) {
		t.Fatal("Unexpected error", err.Error(), "should contains", strconv.Itoa(i))
	}
}

func TestEmpty(t *testing.T) {
	err := error(nil)
	result := Empty(err)
	if !result {
		t.Fatal("Unexpected error")
	}
	err = New("")
	if !result {
		t.Fatal("Unexpected error")
	}
}

func TestEqual(t *testing.T) {
	err1 := New("test")
	err2 := New("test1")
	if Equal(err1, err2) {
		t.Fatal("Unexpected error", err1, err2)
	}
	err2 = New("test")
	if !Equal(err1, err2) {
		t.Fatal("Unexpected error", err1, err2)
	}
}

func TestPrefix(t *testing.T) {
	prefix := "pre"
	error := New("test")
	result := Prefix(error, prefix)
	if !strings.Contains(result.Error(), "pre") {
		t.Fatal("Unexpected error")
	}

	error = New("test")
	error = Multi(error).Add(New("test1"))
	result = Prefix(error, prefix)
	list := Multi(result).Get()
	if len(list) <= 1 {
		t.Fatal("Unexpected len erors list", list)
	}
	for _, e := range list {
		if !strings.Contains(e.Error(), "pre") {
			t.Fatal("Unexpected error", e.Error())
		}
	}
}

func TestGrace_Error(t *testing.T) {
	err := New("test")
	if err.Error() != "test" {
		t.Fatal("Unexpected error", err)
	}
	err = Multi(err).Add(New("test1"))
	if err.Error() != "test1" {
		t.Fatal("Unexpected error", err)
	}

}
