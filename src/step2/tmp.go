package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error")

func myfunc(){

}

func main() {
	//fmt.Printf("error: %v, %v", errNotFound, errNotFound.Error())
	//for i,j:=0,1;i<10 ;i++ {
	//	fmt.Println(i ,j)
	//	j++
	//	i++
	//}
	err :=shadow()
	fmt.Println(err)
}
func check1()(x int, err error)  {
	x = 100
	err = errors.New("check2 return err")
	return
}

func check2(x int)(y int, err error)  {
	err = errors.New("check2 return err")
	y = 200
	return
}

func shadow()(err error){
	x, err := check1() // x是新创建变量，err是被赋值
	//err:= errors.New("haha")
	//y, err := check2(x)
	fmt.Println("x:", x)
	if err != nil {
		return // 正确返回err
	}
	//if y, err := check2(x); err != nil { // y和if语句中err被创建
	//	return // if语句中的err覆盖外面的err，所以错误的返回nil！
	//} else {
	//	fmt.Println(y)
	//}

	if y, err1 := check2(x); err1 != nil { // y和if语句中err被创建
		return err1// if语句中的err覆盖外面的err，所以错误的返回nil！
	} else {
		fmt.Println(y)
	}
	fmt.Println(y)
	//return err1
	return
}