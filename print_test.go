package grace

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	p := Print()
	if p == nil {
		t.Fatal("Unexpected error")
	}
}

func TestPrinter_Err(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	err := errors.New("test")
	print := func(e error) {
		log.Println(e.Error())
	}
	Print().Err(print, err)
	if len(buf.Bytes()) <= 0 {
		t.Fatal("Unexpected error, buffer len", len(buf.Bytes()))
	}
}

func TestPrinter_Stack(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	print := func(e error) {
		log.Println(e.Error())
	}
	err := func() (e error) {
		defer Print().Stack(print, &e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length := len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 || len(buf.Bytes()) <= 0 || strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}
}

func TestPrinter_Recover(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	print := func(e error) {
		log.Println(e.Error())
	}
	err := func() (e error) {
		defer Print().Recover(print, &e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length := len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length != 1 || len(buf.Bytes()) <= 0 || !strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}
}

func TestPrinter_RecoverStack(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	print := func(e error) {
		log.Println(e.Error())
	}
	err := func() (e error) {
		defer Print().RecoverStack(print, &e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length := len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 || len(buf.Bytes()) <= 0 || !strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}
}
