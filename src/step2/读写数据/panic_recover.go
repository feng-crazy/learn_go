// panic_recover.go
package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func test() (tmp string, err2 error){
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
			err2 = fmt.Errorf("Panicing:%s\r\n", e)
		}
		tmp = "already recover"
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
	tmp = "After bad call"
	return tmp, err2
}

func main() {
	fmt.Printf("Calling test\r\n")
	tmp, err := test()
	fmt.Printf("Test completed :%s,%s\r\n", tmp, err)
}