package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x)

	x = min(3)
	fmt.Printf("The minimum is: %d\n", x)

	typecheck(1, "2")
}

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func typecheck(values ... interface{}) {
	for _, value := range values {
		switch v := value.(type) {
			case int:
				fmt.Println("int")
			case float32:
				fmt.Println("float")
			case string:
				fmt.Println("string")
			case bool:
				fmt.Println("bool")
			default:
				fmt.Println("default", v)
		}
	}
}