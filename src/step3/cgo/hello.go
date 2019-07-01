// hello.go

package main

//void SayHello(_GoString_ s);
import "C"

import (
	"fmt"
)

func main() {
	C.SayHello("Hello, World\n")
}

//export SayHello
func SayHello(s string) {
	fmt.Print(s)
}

//package main
//
//////void SayHello(const char* s);
//
////#include "hello.h"
//import "C"
//
//func main() {
//	C.SayHello(C.CString("Hello, World\n"))
//	//C.SayHello("Hello, World\n")  x
//	//C.SayHello(C.string("Hello, World\n"))  x
//}