package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	suck1(ch1, ch2)

	time.Sleep(1e9)
}

func pump1(ch chan int) {
	time.Sleep(1*time.Second)
	for i := 0; ; i++ {
		ch <- i * 2
		time.Sleep(1*time.Second)
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
		time.Sleep(999*time.Millisecond)
	}
}

func suck1(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}

		fmt.Println("**********")
	}
}