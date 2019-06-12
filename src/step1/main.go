package main

import "fmt"

func getSequence(a, b int) func() int {

	say := func() {
		fmt.Println("Hello")
	}
	say()
	value := a + b
	fmt.Println(value)
	i := 0

	//return_func := func() int
	return func() int {
		i += 1
		return i
	}
	//return return_func

}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	var a = 10
	var b = 25
	var balance = []float32{1000.0, 2, 3.4, 7.0, 50.0}
	var tmp_array = [4]int{99}

	nextNumber := getSequence(a, b)

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence(a, b)
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	for i := range balance {
		print(i, " ")
		print(balance[i])
		print("\n")
	}

	for i := range tmp_array {
		print(i, " ")
		print(balance[i])
		print("\n")
	}
}
