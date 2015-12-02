package gomonkey

import (
	"fmt"
	//"os"
	"reflect"
	"strconv"
	"sync"
	"errors"
	//"time"
	"./backend"
)

var (
	errEmptyFuncs = errors.New("Empty list of functions for testing")
)

type GoMonkey struct {
	// Path can be path to the dir or path to the file
	Path string

	// Backend for store arguments if session is crached
	Backendstore backend.Backend 

	items []interface{}
}

func GenTestsByName(fu string) {

}

func (mon *GoMonkey) Add(item interface {}) {
	mon.items = append(mon.items, item)
}

func (mon *GoMonkey) Start() error {
	if len(mon.items) == 0 {
		return errEmptyFuncs
	}
	var wait sync.WaitGroup
	wait.Add(len(mon.items))
	for _, item := range mon.items {
		go func(it interface{}){
			mon.genTests(it)
			wait.Done()
		}(item)
	}
	wait.Wait()
	return nil
}

// GeтTests provides generation and testing target method
func (mon *GoMonkey) genTests(item interface{}) error {
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

		mon.setBackend()
		var count int
		count = 0
		// TODO: If function is fails, write to storage list of arguments
		for i, arg := range items {
			mon.Backendstore.Write(strconv.Itoa(i), arg)
			count++
			value.Call(arg)
		}
		fmt.Println("All tests are passed")
		fmt.Println(fmt.Sprintf("Total number of tests: %d", count))
	}
	return nil
}

func (mon *GoMonkey) setBackend() {
	if mon.Backendstore == nil {
		mon.Backendstore = backend.FileBackend{Path: "current"}
	}
}

