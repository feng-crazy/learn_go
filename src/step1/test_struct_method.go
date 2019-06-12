package main

import (
	"fmt"
)

type call_func func(int) bool
type anwser_func func(int) bool

type mail func(int) bool

type phone struct {
	cb    call_func
	an    anwser_func
	name  string
	email mail
}

type iphone struct {
	phone
	iphone_feature string
}

type huawei struct {
	phone
	huawei_feature string
}

func (this *phone) answer() bool {
	fmt.Println(this.name, "answer ...")
	return true
}

func (this *phone) call() bool {
	fmt.Println(this.name, "call ...")
	this.name = "name is call"
	return true
}

func mymail(x int) bool {
	fmt.Println("mymail:", x)
	return false
}

func (this *phone) qq(x int) bool {
	fmt.Println(" qq :", this.name, x)
	return false
}

func main() {
	p1 := new(iphone)
	p1.name = "iphone"
	p1.iphone_feature = "laji"
	p2 := new(huawei)
	p2.name = "huawei"
	p2.huawei_feature = "niubi"
	p1.call()
	fmt.Println("p1 call name change:", p1.name)
	p2.call()
	fmt.Println("p2 call name change:", p2.name)
	p1.answer()
	p2.answer()

	p1.email = mymail
	p1.email(1)

	p2.email = func(i int) bool {
		fmt.Println("p2 email:", i)
		return false
	}
	p2.email(2)

	p1.qq(3)
	p2.qq(4)

	p1.cb = func(i int) bool {
		fmt.Println("p2 cb:", i)
		return false
	}

	p2.an = func(i int) bool {
		fmt.Println("p2 an:", i)
		return false
	}

	p1.cb(5)
	p2.an(6)

}
