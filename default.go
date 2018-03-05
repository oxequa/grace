package grace

import (
	"errors"
	"reflect"
)

type (
	def     struct{}
	Default interface{}
)

func Def() Default {
	return def{}
}

// Zero check if a value is empty
func (d def) Zero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && d.Zero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && d.Zero(v.Field(i))
		}
		return z
	case reflect.Ptr:
		if !v.IsNil() {
			return d.Zero(reflect.Indirect(v))
		}
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}

// Set to default value a given interface
func (d def) Set(init interface{}, def interface{}) error {
	if def == nil {
		return errors.New("invalid default param")
	}
	if reflect.Indirect(reflect.ValueOf(init)).Kind() != reflect.TypeOf(def).Kind() {
		return errors.New("different types")
	}
	reflect.Indirect(reflect.ValueOf(init)).Set(reflect.ValueOf(def))
	return nil
}

// New return an interface value equal to init or to def if init is zero
func (d def) New(init interface{}, def interface{}) interface{} {
	return nil
}
