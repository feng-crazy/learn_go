package main

import (
	"fmt"
	"math"
)

// 定义类型
type ErrNegativeSqrt float64

// 重写Error()
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number:  %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", result)
}
