package gomonkey

import (
	"fmt"
	//"os"
	"reflect"
	"sync"
	//"time"
	"./backend"
)

type GoMonkey struct {
	// Path can be path to the dir or path to the file
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

		items, err := funcitem.generate(args)
		if err != nil {
			return err
		}

		backendstore := backend.FileBackend{Path: "current"}
		var wg sync.WaitGroup
		wg.Add(len(items))
		// TODO: If function is fails, write to storage list of arguments
		for _, arg := range items {
			backendstore.Write("1", arg)
			go func() {
				value.Call(arg)
				wg.Done()
			}()
		}
		wg.Wait()
	}
	return nil
}

