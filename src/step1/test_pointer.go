package main

import "fmt"

const MAX int = 3

func main() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}

	tmp_str := "you are sb"
	var ptr1 *string
	ptr1 = &tmp_str
	println(ptr1)

	for i = 0; i < len(tmp_str); i++ {
		//ptr1 =ptr1 + 1
		fmt.Printf("tmp_str[%d] = %c\n", i, tmp_str[i])
	}

	arr := []int{0, 1, 2, 3}
	var arr_ptr = &arr
	//arr_ptr += 1
	fmt.Printf("%d\n", (*arr_ptr)[1])
	fmt.Printf("%d\n", arr[1])

	var tmp_func = func() {
		i := 10
		fmt.Println("tmp_func %d", i)
	}
	var tmp_func_ptr = &tmp_func
	(*tmp_func_ptr)()
}
