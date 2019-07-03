package main

import (
	"fmt"
	"reflect"
)

type NotKnownTypeTest struct {
	f1, f2, f3 float64
}

type NotknownType struct {
	s1, s2, s3 string
	integer1   int
	n1 NotKnownTypeTest
}

func (n NotknownType) String2() string{
	return n.s2 + " - " + n.s3 + " - " + n.s1
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon", 100,NotKnownTypeTest{1.0, 20, 30}}

func main() {
	var value reflect.Value
	value = reflect.ValueOf(secret) // <main.NotknownType Value>
	//typ := reflect.TypeOf(secret)    // main.NotknownType
	typ := value.Type()    // main.NotknownType
	// alternative:
	//typ := value.Type()  // main.NotknownType
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		var valueField reflect.Value
		valueField = value.Field(i)
		fmt.Printf("Field %d: %v , %s \n", i, valueField, valueField.String())
		if i==3{
			fmt.Printf("Field %d: %v , %d \n", i, valueField, valueField.Int())
		}
		if i==4{
			fmt.Printf("Field %d: %v , %d \n", i, valueField, valueField.Bytes())
		}
		//b_arr := valueField.Bytes()
		//for i, _ := range b_arr {
		//	fmt.Printf("%02x ", i)
		//}
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		//value.Field(i).SetString("C#")
	}
	fmt.Println("********", value.NumMethod())
	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
	fmt.Println(value.Method(1).Call(nil))
}