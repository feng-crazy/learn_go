package main

import (
	. "./test_pack_1"
	"fmt"
)

//import . "./test_pack_1"

func main() {
	var test1 string
	//test1 = test_pack_1.ReturnStr()
	test1 = ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	//fmt.Printf("Integer from package1: %d\n", test_pack_1.Pack1Int)
	fmt.Printf("Integer from package1: %d\n", Pack1Int)
	PackFloat = ReturnFloat()
	fmt.Printf("Float from package1: %f\n", PackFloat)
}