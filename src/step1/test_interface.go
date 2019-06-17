package main

import (
	"fmt"
)

type Phone interface {
	call() bool
}

type _3GPhone struct {
	//Phone
}

type NokiaPhone struct {
	_3GPhone
	price int32
}

func (nokiaPhone NokiaPhone) call() bool{
	fmt.Println("I am Nokia, I can call you!")
	return true
}

type IPhone struct {
	_3GPhone
	price int32
}

func (iPhone IPhone) call() bool {
	fmt.Println("I am iPhone, I can call you!")
	return true
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	fmt.Println(phone)
	phone.call()

	phone = new(IPhone)
	fmt.Println(phone)
	phone.call()

	phone = IPhone{price:1500}
	fmt.Println(phone)
	phone.call()

	phone = &IPhone{}
	fmt.Println(phone)
	phone.call()
}
