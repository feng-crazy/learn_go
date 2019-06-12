package main

import (
	"log"
	"runtime"
)

func where_is() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("where_is %s:%d", file, line)
}

func test_where() func(){
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("where %s:%d", file, line)
	}

	where()
	return where
}

func main(){
	where := test_where()


	where_is()

	where()
}