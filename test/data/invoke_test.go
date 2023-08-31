package data

import (
	"fmt"
	"github.com/laoningmeng/fusion/internal/data"
	"reflect"
	"strings"
	"testing"
)

func TestInvoke(t *testing.T) {
	m := data.NewMysqlFaker()
	d := m.Data
	res := reflect.ValueOf(d).MethodByName("Name").Call(nil)
	fmt.Println(res)
}

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3}
	fmt.Println(s[1:])
	fmt.Println(s[3:])
}

func TestUppercase(t *testing.T) {
	s := "word"
	a := strings.ToUpper(s[:1]) + s[1:]
	fmt.Println(a)
}
