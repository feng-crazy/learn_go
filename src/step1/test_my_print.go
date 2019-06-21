package main

import (
	"fmt"
	"os"
	"strconv"
)

type Stringer interface {
	String() string
}

type Celsius float64

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c),'f', 1, 64) + " °C"
}

type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func myPrint(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {os.Stdout.WriteString(" ")}
		switch a := arg.(type) { // type switch
			//case Day: os.Stdout.WriteString(a.String()); fmt.Println("  Day")
			case Stringer:	os.Stdout.WriteString(a.String()); fmt.Println("  stringer")
			case int:		os.Stdout.WriteString(strconv.Itoa(a));fmt.Println("  int")
			case string:	os.Stdout.WriteString(a); fmt.Println("  string")
			// more types
			default:		os.Stdout.WriteString("???")
		}
	}
}

func main() {
	myPrint(Day(1), "was", Celsius(18.36), 5)  // Tuesday was 18.4 °C
}