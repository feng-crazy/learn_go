package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var t interface{}= T{23, "skidoo"}
	//t := T{23, "skidoo"}
	//s := reflect.ValueOf(&t).Elem()
	s := reflect.ValueOf(t)
	fmt.Println(s)
	typeOfT := s.Type()
	fmt.Println("typefot:", typeOfT)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	//s = s.Elem()
	//s.Field(0).SetInt(77)
	//s.Field(1).SetString("Sunset Strip")

	fmt.Println("t is now", t)

}