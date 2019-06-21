package main

import (
	"fmt"
)

type Phone1 interface {
	call() bool
}

type _3GPhone1 struct {
	Phone
}

type NokiaPhone1 struct {
	_3GPhone
	price int32
}

func (NokiaPhone1 NokiaPhone1) call() bool{
	fmt.Println("I am Nokia, I can call you!")
	return true
}

type IPhone1 struct {
	_3GPhone
	price int32
}

func (IPhone1 IPhone1) call() bool {
	fmt.Println("I am IPhone1, I can call you!")
	return true
}

func main() {
	var phone Phone1

	phone = new(NokiaPhone1)
	fmt.Println(phone)
	phone.call()

	phone = new(IPhone1)
	fmt.Println(phone)
	phone.call()

	phone = IPhone1{price:1500}
	fmt.Println(phone)
	phone.call()

	phone = &IPhone1{}
	fmt.Println(phone)
	phone.call()
}
