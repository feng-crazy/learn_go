package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main()  {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1" + ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := &Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith Multiply: %d*%d=%d\n", args.A, args.B, reply)

	// Asynchronous call
	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	fmt.Printf("Arith Divide: %d*%d=%d\n", args.A, args.B, quotient)
	replyCall := <-divCall.Done	// will be equal to divCall
	// check errors, print, etc.
	fmt.Printf("Arith Divide: %d*%d=%d\n", args.A, args.B, quotient)
	fmt.Println("replyCall:", replyCall)
}

//func main() {
//	client, err := rpc.Dial("tcp", "localhost:1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}
//
//	var reply string
//	err = client.Call("HelloService.Hello", "hello", &reply)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(reply)
//}
