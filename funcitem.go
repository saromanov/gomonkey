package gomonkey

import (
	"./gen"
	"fmt"
	"reflect"
)

// Generate test for funcs

type FuncItem struct {
	Data  interface{}
	Value reflect.Value
	Count uint
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

// generate provides generation of lists of arguments
func (fi *FuncItem) generate(args []reflect.Kind) ([]reflect.Value, error) {
	if fi.Count == 0 {
		return nil, fmt.Errorf("Count must be > 0")
	}
	result := []reflect.Value{}
	var i uint
	typevalue := reflect.TypeOf(fi.Data)
	for i = 0; i < fi.Count; i++ {
		newargs := []reflect.Value{}
		for j, arg := range args {
			switch arg {
			case reflect.Slice:
				fmt.Println(typevalue.In(int(j)).Elem())
				break
			case reflect.Int:
				genint := gen.GenInt{}
				value, _ := genint.GenerateOne()
				newargs = append(newargs, reflect.Value(value))
				break
			case reflect.String:
				fmt.Println("STRING")
				break
			}
		}
	}

	return result, nil
}
