package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time //anonymous field
}

type testTime time.Time

func (t myTime) first3Chars() string {
	return t.String()[0:3]
	//return t.Time.String()[0:3]
}

func (t testTime) first3Chars() string {
	return time.Time(t).String()[0:3]
	//return t.Time.String()[0:3]
}

func main() {
	m := myTime{time.Now()}
	// 调用匿名Time上的String方法
	fmt.Println("Full time now:", m.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())

	tm := testTime(time.Now())
	fmt.Println("First 3 chars:", tm.first3Chars())
}

/* Output:
Full time now: Mon Oct 24 15:34:54 Romance Daylight Time 2011
First 3 chars: Mon
*/