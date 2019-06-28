package main

import (
	"bytes"
	"fmt"
)

func main()  {
	const form = `ni ma`
	fmt.Printf("%T\n", form)

	c := "hello"
	fmt.Println(c)
	b := []byte(c)
	b[0] = 'c'
	c = string(b)
	fmt.Println(c)

	c_tmp := c[3]
	fmt.Println(c_tmp)

	var by = bytes.Buffer{}
	by.WriteString(c)
	by.WriteString("world")
	fmt.Println(by.String())

	map1 := map[string]int{"one": 1, "two": 2}

	val1, isPresent := map1["one"]
	fmt.Println(val1, isPresent)
}
