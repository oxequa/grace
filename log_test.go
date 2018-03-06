package grace

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLogger_Save(t *testing.T) {
	d, err := ioutil.TempFile("", "test.log")
	if err != nil {
		t.Fatal(err)
	}
	l := logger{fname: d.Name()}
	err = l.save([]byte("test string"))
	if err != nil {
		t.Fatal(err)
	}
	fi, e := os.Stat(d.Name())
	if e != nil {
		t.Fatal(e)
	}
	if fi.Size() <= 0 {
		t.Fatal("Unexpected file size", fi.Size())
	}
}

func TestLogger_Name(t *testing.T) {
	l := logger{
		enable: true,
		fname:  fname,
	}
	if l.name() != l.fname {
		t.Fatal("Unexpected errror", l.fname, "instead", l.name())
	}
}

func TestLogger_State(t *testing.T) {
	l := logger{
		enable: true,
		fname:  fname,
	}
	if !l.state() {
		t.Fatal("Unexpected errror", l.enable, "instead", l.state())
	}
}
