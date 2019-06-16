package main

import "fmt"

type Element2 interface{}

type Vector struct {
	a []Element2
}

func (p *Vector) At(i int) Element2 {
	return p.a[i]
}

func (p *Vector) Set(i int, e Element2) {
	p.a[i] = e
}

func main(){
	var vector Vector
	vector.a = make([]Element2, 2, 2)
	vector.Set(0,"haha")
	fmt.Println(vector.At(0))
	v, ok := vector.At(0).(string)
	fmt.Println(v, ok)
	tmpstr := "you are sb" + v
	fmt.Println(tmpstr)
}