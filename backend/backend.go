package backend

import (
   "reflect"
)

// Backend provides storing data to physics backend (file, db, etc...)
type Backend interface {

	// Write provides writing arguments to backend
	Write(title string, data []reflect.Value) error

	// Get returns arguments for testing func by title
	Get(title string)[]reflect.Value

	// String returns type of backend
	String() string
}