package grace

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestErr(t *testing.T) {
	err := Err()
	if err == nil {
		t.Fatal("Unexpected error")
	}
}

func TestLog(t *testing.T) {
	i := Log()
	if i == nil {
		t.Fatal("Unexpected error")
	}
	l := i.getLog()
	if !l.enable {
		t.Fatal("Unexpected error")
	}
}

func TestPrint(t *testing.T) {
	i := Print()
	if i == nil {
		t.Fatal("Unexpected error")
	}
	l := i.getPrint()
	if !l.enable {
		t.Fatal("Unexpected error")
	}
}

func TestErr_Log(t *testing.T) {
	i := Err().Log("name.log")
	if i == nil {
		t.Fatal("Unexpected error")
	}
	l := i.getLog()
	if !l.enable {
		t.Fatal("Unexpected error")
	}
	if l.fname != "name.log" {
		t.Fatal("Unexpected error")
	}
}

func TestErr_Print(t *testing.T) {
	f := func(e error) {
		log.Println(e)
	}
	i := Err().Print(f)
	if i == nil {
		t.Fatal("Unexpected error")
	}
	l := i.getPrint()
	if !l.enable {
		t.Fatal("Unexpected error")
	}
}

func TestErr_New(t *testing.T) {
	var a error
	var b int
	b = 10
	err := Err().New(a, b)
	if err.Error() != strconv.Itoa(b) {
		t.Fatal("Unexpected value", err.Error(), strconv.Itoa(b))
	}
}

func TestErr_Set(t *testing.T) {
	var a error
	var b int
	b = 10
	Err().Set(&a, b)
	if a.Error() != strconv.Itoa(b) {
		t.Fatal("Unexpected value", a.Error(), strconv.Itoa(b))
	}
}

func TestErr_Stack(t *testing.T) {
	f := func(e error) {
		log.Println(e.Error())
	}
	err := func() (e error) {
		defer Err().Stack(&e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length := len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 || strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}

	// set output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// print
	err = func() (e error) {
		defer Err().Print(f).Stack(&e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length = len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 || len(buf.Bytes()) <= 0 || strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}

	// log
}

func TestErr_Recover(t *testing.T) {
	f := func(e error) {
		log.Println(e.Error())
	}
	err := func() (e error) {
		defer Err().Recover(&e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length := len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length != 1 || !strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}

	// set output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// print
	err = func() (e error) {
		defer Err().Print(f).Recover(&e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length = len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length != 1 || len(buf.Bytes()) <= 0 || !strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}

	// log
}

func TestErr_RecoverStack(t *testing.T) {
	f := func(e error) {
		log.Println(e.Error())
	}
	err := func() (e error) {
		defer Err().RecoverStack(&e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length := len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 || !strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}

	// set output
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// print
	err = func() (e error) {
		defer Err().Print(f).RecoverStack(&e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length = len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 || len(buf.Bytes()) <= 0 || !strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}

	// log
	err = func() (e error) {
		defer Log().RecoverStack(&e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length = len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 || !strings.Contains(err.Error(), "index out of range") {
		t.Fatal("Unexpected", err, err.Error(), length)
	}
}
