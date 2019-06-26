package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan string, 3)
	go pump(ch1, ch2)
	go suck(ch1, ch2)
	for i := 0; i< 100; i++{
		fmt.Println("*************************:", i)
		time.Sleep(1e9)
	}
}

func pump(ch1 chan int, ch2 chan string) {
	for i := 0; i< 100; i++ {
		ch1 <- i
		fmt.Println(<-ch2)
	}
}

func suck(ch1 chan int, ch2 chan string) {
	for i := 0; i< 100 ; i++{
		ch2<- string(i)
		fmt.Println(<-ch1)

	}
}