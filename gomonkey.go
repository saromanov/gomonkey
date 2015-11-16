package gomonkey

import (
	"fmt"
	//"os"
	"reflect"
)

type GoMonkey struct {
	// path can be path to the dir or path to the file
	Path string
}

func GenTestsByName(fu string) {

}

func (mon *GoMonkey) GenTests(item interface{}) error {
	value := reflect.ValueOf(item)
	if !value.IsValid() {
		return fmt.Errorf("Item is invalid")
	}

	if value.Kind() == reflect.Func {
		funcitem := FuncItem{Data: item, Value: value, Count: 10}
		args, err := funcitem.getArgumentTypes()
		if err != nil {
			return err
		}

		items := funcitem.generate(args)
		items.call(item, items)
	}
	return nil
}


// call function with arguments
func (mon *GoMonkey) call(item interface{}, args []reflect.Value) {

}
