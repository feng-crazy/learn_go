package main

import (
	"fmt"
	"time"
)

var gloablI = 0

func say(s string, time_t time.Duration) {
	var interval time.Duration
	interval = time_t * time.Millisecond
	for i := 0; i < 5; i++ {
		fmt.Println(s, gloablI)
		time.Sleep(interval)
		gloablI += 1

	}
}

func main() {
	go say("world", 100)
	//time.Sleep(10*time.Millisecond)
	say("hello", 200)
}
