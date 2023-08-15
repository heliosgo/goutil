package uruntime

import (
	"errors"
	"log"
	"reflect"
)

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recover: %s", err)
			}
		}()

		f()
	}()
}

func GoArgs(f any, args ...any) error {
	if reflect.TypeOf(f).Kind() != reflect.Func {
		return errors.New("f must be function")
	}

	in := make([]reflect.Value, len(args))
	for i, v := range args {
		in[i] = reflect.ValueOf(v)
	}

	Go(func() { reflect.ValueOf(f).Call(in) })

	return nil
}
