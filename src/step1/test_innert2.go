package main

import (
	"fmt"
	"reflect"
)

type Log2 struct {
	msg string
}

type Customer2 struct {
	Name string
	*Log2
}

func (l *Log2) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log2) String() string {
	return l.msg
}

func (c *Customer2) log2() *Log2 {
	fmt.Println("log print",reflect.TypeOf(c.Log2))
	return c.Log2
}

func main() {
	c := new(Customer2)
	c.Name = "Barak Obama"
	//c.msg = "haha"  //error
	c.Log2 = new(Log2)
	c.Log2.msg = "1 - Yes we can!"
	c.msg = "1 - Yes we can!"
	// shorter
	c = &Customer2{"Barak Obama", &Log2{"1 - Yes we can!"}}
	//c = &Customer2{"Barak Obama", &{"1 - Yes we can!"}} // error
	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	c.Add("2 - After me the world will be a better place!")
	c.Log2.Add(" 2 b ?")
	//fmt.Println(c.Log2)
	fmt.Println(c.Log2.msg)
	fmt.Println(c.msg)
	fmt.Println(c.log2())

	var re_log *Log2 = c.log2()
	var tmp_str string = c.log2().String()
	fmt.Println("tmp_str:", tmp_str)
	fmt.Println("----------",re_log.String())
}

