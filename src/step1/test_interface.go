package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type _3GPhone struct {
	Phone
}

type NokiaPhone struct {
	_3GPhone
	price int
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
	_3GPhone
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()
	//fmt.Print(phone.price)

	nokiaPhone := NokiaPhone{price: 1000}
	fmt.Println(nokiaPhone.price)

	//phone = nokiaPhone
	//	//fmt.Print(phone.price)

	phone = new(IPhone)
	phone.call()

	phone = IPhone{}
	phone.call()

}
