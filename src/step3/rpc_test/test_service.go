package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main()  {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	for  {
		time.Sleep(1*time.Second)
	}
}

//type HelloService struct {}
//
//func (p *HelloService) Hello(request string, reply *string) error {
//	*reply = "hello:" + request
//	return nil
//}
//
//
//func main() {
//	err := rpc.RegisterName("HelloService", new(HelloService))
//	if err != nil {
//		log.Fatal("rpc.RegisterName error:", err)
//	}
//
//	listener, err := net.Listen("tcp", ":1234")
//	if err != nil {
//		log.Fatal("ListenTCP error:", err)
//	}
//
//	conn, err := listener.Accept()
//	if err != nil {
//		log.Fatal("Accept error:", err)
//	}
//
//	rpc.ServeConn(conn)
//}