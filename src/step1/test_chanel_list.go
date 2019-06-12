package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("c <- x", x)
		c <- x
		x, y = y, x+y
	}
	//close(c)  // 如果没有关闭通道 协程结束，通道另一端还在接收将会报错
	//for true{
	//      time.Sleep(1000*time.Millisecond)
	//}
	fmt.Println("goroutine finish")
}

func main() {
	c := make(chan int, 20)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	count := 0
	for i := range c {
		fmt.Println(i)
		count++
		//if count ==10{
		//       fmt.Println("close(c)", i)
		//       close(c)
		//}
		//if count ==20{
		//        close(c)
		//}
	}
	for true {
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("main finish")
}
