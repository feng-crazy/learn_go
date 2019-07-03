package main

import (
	"fmt"
)

type Quotient struct {
	Quo, Rem int
}

func myFunc(arg interface{}){
	reply, ok := arg.(*Quotient)
	fmt.Println(reply, ok)
	reply.Quo = 100
	reply.Rem = 200
}

func main()  {
	quotient := new(Quotient)
	var reply interface{}
	fmt.Println("reply:", reply)
	reply = quotient
	fmt.Println("reply:", reply)
	myFunc(reply)
	fmt.Println("quotient:", quotient)
	fmt.Println("reply:", reply)
}
