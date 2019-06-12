package main

import (
	"fmt"
	"time"
)

func cal(a int, b int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
}

func main() {
	for i := 0; i < 10; i++ {
		go cal(i, i+1) //启动10个goroutine 来计算
	}
	time.Sleep(time.Second * 2) // sleep作用是为了等待所有任务完成
}

//结果
//8 + 9 = 17
//9 + 10 = 19
//4 + 5 = 9
//5 + 6 = 11
//0 + 1 = 1
//1 + 2 = 3
//2 + 3 = 5
//3 + 4 = 7
//7 + 8 = 15
//6 + 7 = 13
