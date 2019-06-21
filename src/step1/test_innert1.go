package main

import (
	"fmt"
	"reflect"
)

type Log1 struct {
	msg string
}

type Customer1 struct {
	Name string
	Log1
}

func (l *Log1) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log1) String() string {
	return l.msg
}

func (c *Customer1) log1() Log1 {
	fmt.Println("log print",reflect.TypeOf(c.Log1))
	return c.Log1
}

func main() {
	c := new(Customer1)
	c.Name = "Barak Obama"
	c.Log1 = Log1{"ni ma"}
	c.msg = "1 - Yes we can!"
	c = &Customer1{"Barak Obama", Log1{"1 - Yes we can!"}}
	//c = &Customer1{"Barak Obama", {"1 - Yes we can!"}}  // error
	//c = &Customer1{"Barak Obama", "1 - Yes we can!"}  // error
	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	c.Add("2 - After me the world will be a better place!")
	c.Add(" 2 b ?")
	//fmt.Println(c.log)
	fmt.Println(c.msg)
	var re_log Log1 = c.log1()
	var tmp_str string = c.log1().String()
	fmt.Println("tmp_str:", tmp_str)
	fmt.Println("----------",re_log.String())

}

