package main

import (
	"fmt"
	"reflect"
)

type ToString  func() string

type T struct {
	A int
	B string
	//myString ToString
}

func impString() string{
	tmp := "haha sb"
	return tmp
}

func(t T) TString() string{
	tmp := "haha sb" + t.B
	return tmp
}

func main() {
	//var t interface{}= T{23, "skidoo", impString}
	//var t interface{}= T{23, "skidoo"}
	t := T{23, "skidoo"}
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

	fmt.Println("*************", s.NumMethod())

	ret := s.Method(0).Call(nil)
	fmt.Println("ret:", ret)

	//imp_t, ok := t.(T)
	//fmt.Println(imp_t, ok)
	//fmt.Println(imp_t.TString())

	//s = s.Elem()
	//s.Field(0).SetInt(77)
	//s.Field(1).SetString("Sunset Strip")
	//t.A = 77
	//t.B = "Sunset Strip"

	fmt.Println("t is now", t)

}