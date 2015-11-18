package backend

import (
  "os"
  "encoding/json"
  "reflect"
)

type FileBackend struct {
	Path string
}

// Write provides writing arguments to file
func (fb *FileBackend) Write(title string, data []reflect.Value) error{
	result := map[string][]reflect.Value{}
	result[title] = data
	ser, errmar := json.Marshal(result)
	if errmar != nil {
		return errmar
	}
	f, err := os.OpenFile(fb.Path, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	defer f.Close()
	if err != nil {
		return err
	}
	f.Write(ser)

	return nil

}

// read provides reading arguments from file
func (fb *FileBackend) Read(title string) []reflect.Value {
	return []reflect.Value{}
}

// Strring returns name of backend
func (fb *FileBackend) String() string{
	return "filebackend"
}

