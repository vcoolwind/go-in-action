package function

import (
	"fmt"
	"reflect"
	"testing"
)

func Apply(f interface{}, args []interface{}) []reflect.Value {
	fun := reflect.ValueOf(f)
	inargs := make([]reflect.Value, len(args))
	for k, param := range args {
		inargs[k] = reflect.ValueOf(param)
	}
	r := fun.Call(inargs)
	return r

}

// 变参
func Format(args ...interface{}) string {
	fmt.Println(args...)
	return "format return"
}

func AddHello(a int, b int) (int, string) {
	return a + b, "hello"
}

func TestFunc(t *testing.T) {

	ret1 := Apply(AddHello, []interface{}{1, 2})
	for _, v := range ret1 {
		fmt.Println(v)
	}

}
