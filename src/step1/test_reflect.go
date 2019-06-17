package main

import (
	"fmt"
	"reflect"
)

type myFloat64 float64

func main() {
	var x myFloat64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	if v.Kind() == reflect.Float64{
		fmt.Println("haha")
	}
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(myFloat64)
	fmt.Println(y)
}