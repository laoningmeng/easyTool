package data

import (
	"os"
	"reflect"
	"strings"
)

type dataSource interface {
	invoke(method string, args ...interface{}) (res any, err error)
}

func invoke(source interface{}, method string, args ...interface{}) []reflect.Value {
	var input []reflect.Value
	if len(method) < 1 {
		os.Exit(1)
	}
	method = strings.ToUpper(method[:1]) + method[1:]
	for _, e := range args {
		input = append(input, reflect.ValueOf(e))
	}
	return reflect.ValueOf(source).MethodByName(method).Call(input)
}
