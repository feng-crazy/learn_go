package main

import (
	"fmt"
	"runtime"
)

type TestRun struct {
	author  string
	haha    string
}

func my_test(){
	test := new(TestRun)
	//var test TestRun
	test.haha = "sb"
	test.author = "hedengfeng"
	fmt.Println(test)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)

	runtime.SetFinalizer(test, func(test *TestRun) {print("free test")})
}

func main() {
	my_test()
	runtime.GC()
	//time.Sleep(10*time.Second)
}

