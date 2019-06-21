package main

import "fmt"

type Shaper interface {
	Area() float32
	Perimeter() float32
	//Area(int, int) int
}

type Geometry struct {

}

type Square struct {
	Geometry
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 {
	return 4 * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.length + r.width)
}


func main() {
	//sq1 := new(Square)
	//sq1.side = 5

	var sq1 = new(Square)
	sq1.side = 5

	sq2 := Square{side:6.0}

	var areaIntf Shaper
	areaIntf = sq1

	// error because 接口实现是指针
	//areaIntf = sq2
	//areaIntf = Square{side:6.0}

	//correct
	//areaIntf = Rectangle{length:6.0, width:5.0}

	//var shaper Shaper
	//shaper = Square{6.0}
	//fmt.Println(shaper)
	// shorter,without separate declaration:
	//areaIntf := Shaper(sq1)
	//areaIntf2 := Shaper(&sq2)
	//var areaIntf2 Shaper
	areaIntf2 := sq2
	// or even:
	//areaIntf := sq1
	fmt.Printf("The square has area: %f, %f\n", areaIntf.Area(), areaIntf2.Area())

	r := Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{side:5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	areaIntf3 := r
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...", areaIntf3)
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	elem1 := Rectangle{5, 3}
	elem2 := &Rectangle{6, 7}
	test_arr := []Shaper{elem1, elem2}
	fmt.Println(test_arr)
}