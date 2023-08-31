package data

import (
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type FakeIt struct {
	*gofakeit.Faker
}

func NewFakeIt() *FakeIt {
	return &FakeIt{
		gofakeit.New(0),
	}
}

func (f *FakeIt) invoke(method string, args ...interface{}) (res any, err error) {
	result := invoke(f, method, args...)
	first := result[0].Interface()
	switch first.(type) {
	case string:
		res = first.(string)
	case []byte:
		res = first.([]byte)
	case int:
		res = first.(int)
	case time.Time:
		res = first.(time.Time)
	case float32:
		res = first.(float32)
	case float64:
		res = first.(float64)
	default:
		res = ""
		err = errors.New("unknown type")
	}
	return
}
