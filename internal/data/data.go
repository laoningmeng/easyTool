package data

import (
	"reflect"
)

type dataSource interface {
	invoke(method string, args ...interface{}) (res any, err error)
}

func invoke(source interface{}, method string, args ...interface{}) []reflect.Value {
	var input []reflect.Value
	for _, e := range args {
		input = append(input, reflect.ValueOf(e))
	}
	return reflect.ValueOf(source).MethodByName(method).Call(input)
}
