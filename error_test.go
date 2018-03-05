package grace

import (
	"fmt"
	"strings"
	"testing"
)

func TestErr(t *testing.T) {
	err := Err()
	if err == nil {
		t.Fatal("Unexpected error")
	}
}

func TestErr_Stack(t *testing.T) {
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
}

func TestErr_Recover(t *testing.T) {
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
}

func TestErr_RecoverStack(t *testing.T) {
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
}
