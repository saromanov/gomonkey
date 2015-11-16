package gomonkey

import (
	"fmt"
	"reflect"
)

// Generate test for funcs

type FuncItem struct {
	Data  interface{}
	Value reflect.Value
	Count int
}

func (fi *FuncItem) getArgumentTypes() ([]reflect.Kind, error) {
	typevalue := reflect.TypeOf(fi.Data)
	if typevalue.NumOut() == 0 {
		return nil, fmt.Errorf("Function must return ay least one value")
	}

	args := []reflect.Kind{}
	for i := 0; i < typevalue.NumIn(); i++ {
		args = append(args, typevalue.In(i).Kind())
	}

	return args, nil
}

func (fi *FuncItem) generate(args []reflect.Kind) []reflect.Value {
	for i := 0; i < fi.Count; i++ {

	}

	return nil
}
