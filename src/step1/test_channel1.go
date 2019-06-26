package main

import (
	"fmt"
	"time"
)

func main()  {
	c := make(chan int, 3)
	go func(c chan int) {
		for i:=0;i<8 ;i++  {
			select {
			case c <- i:
			default:
			}
		}
		fmt.Println("go done")
		for  {
			time.Sleep(1*time.Second)
		}
	}(c)

	time.Sleep(1*time.Second)

	for j:=0;j<4 ;j++  {
		fmt.Println(<-c)
	}

	for true{
		time.Sleep(1*time.Second)
	}
}
