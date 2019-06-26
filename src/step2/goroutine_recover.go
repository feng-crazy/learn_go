package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"time"
)

func main()  {
	fmt.Println("begin")
	go func() {

		defer func() {
			deferTmp:=100
			fmt.Println("tmp:",deferTmp)
		}()

		defer func() {
			if err := recover(); err!=nil{
				log.Printf("Work failed with %s ", err)
				runtime.Goexit()
			}
		}()

		defer func() {
			deferTmp:=200
			fmt.Println("tmp:",deferTmp)
		}()

		panic(errors.New("my panic "))
		tmp := 300
		fmt.Println("tmp:",tmp)

		defer func() {
			deferTmp:=400
			fmt.Println("tmp:",deferTmp)
		}()

	}()
	fmt.Println("end")
	for  {
		time.Sleep(1*time.Second)
	}
}
