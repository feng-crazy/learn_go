package main

import "fmt"

// 声明一个函数类型
//type cb1 func(int) int
type cb func(x int) (y int)

func main() {
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		x += 1
		return x
	})

	var call func(int) int
	call = callBack
	call(3)
}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) (y int) {
	fmt.Printf("我是回调，x：%d\n", x)
	y = x + 1
	return
}

func callBack1(x int) (y int) {
	fmt.Printf("我是回调，x：%d\n", x)
	y = x + 1
	return
}
