package grace

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestRecover(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	err := func() (e error) {
		defer Recover(&e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return e
	}()
	length := len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 {
		t.Fatal("Unexpected error")
	}

	// add a custom function
	err = func() (e error) {
		e = Handler(e).Add(func(e error) error {
			log.Println(e)
			return e
		})
		defer Recover(&e)
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length = len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 || len(buf.Bytes()) <= 0 {
		t.Fatal("Unexpected error", length, len(buf.Bytes()))
	}
}

func TestRecovery_Error(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	// get only error message
	err := func() (e error) {
		e = Handler(e).Add(func(e error) error {
			log.Println(e)
			return e
		})
		defer Recover(&e).Error()
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length := len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length != 1 || len(buf.Bytes()) <= 0 {
		t.Fatal("Unexpected error", err, length)
	}
}

func TestRecovery_Stack(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	// get only error message
	err := func() (e error) {
		e = Handler(e).Add(func(e error) error {
			log.Println(e)
			return e
		})
		defer Recover(&e).Stack()
		numbers := []int{1, 2}
		fmt.Println(numbers[3])
		return
	}()
	length := len(strings.Split(strings.TrimSuffix(err.Error(), "\n"), "\n"))
	if err == nil || length <= 1 || len(buf.Bytes()) <= 0 {
		t.Fatal("Unexpected error", err, length)
	}
}
