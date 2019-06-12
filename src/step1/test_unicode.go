package main

import (
	"fmt"
	"strings"
)

//func main() {
//	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
//		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
//	}))
//
//}

func main() {
	s := "Hello World!"
	// 创建 Reader
	r := strings.NewReader(s)
	// 创建长度为 5 个字节的缓冲区
	b := make([]byte, 5)

	var err error
	// 循环读取 r 中的字符串
	for n, _ := r.Read(b); n > 0; n, err = r.Read(b) {
		fmt.Printf("%q, %d , %p", b[:n], n, err) // "Hello", " Worl", "d!"
	}
}
