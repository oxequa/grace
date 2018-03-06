package grace

import (
	"bytes"
	"github.com/go-siris/siris/core/errors"
	"log"
	"testing"
)

func TestPrinter_State(t *testing.T) {
	p := printer{
		enable: true,
		fprint: func(e error) {},
	}
	if !p.state() {
		t.Fatal("Unexpected errror", p.enable, "instead", p.state())
	}
}

func TestPrinter_Stamp(t *testing.T) {
	p := printer{
		enable: true,
		fprint: func(e error) {
			log.Println(e.Error())
		},
	}

	// set output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	p.stamp(errors.New("test"))
	if len(buf.Bytes()) <= 0 {
		t.Fatal("Unexpected error")
	}
}

func TestPrinter_Printf(t *testing.T) {
	p := printer{
		enable: true,
		fprint: func(e error) {},
	}
	if p.printf() == nil {
		t.Fatal("Unexpected errror")
	}
}
