package main

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	const a_size = 4
	var balance = [6]float32{1000.0, 2, 3.4, 7.0, 50.0}
	//var balance2 []float32 = {1000.0, 2, 3.4, 7.0, 50.0}
	balance2 := [6]float32{1000.0, 2, 3.4, 7.0, 50.0}
	var tmp_array = [a_size]int{99}
	balance = balance2
	for i := range balance {
		print(i, " ")
		print(balance[i])
		print("\n")
	}

	for i := range tmp_array {
		print(i, " ")
		print(tmp_array[i])
		print("\n")
	}

	var tmp_balance [a_size]int
	tmp_balance = tmp_array

	for i := range tmp_balance {
		print(i, " ")
		print(tmp_array[i])
		print("\n")
	}

	var a = [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	for key, value := range a {
		println(key)
		var tmp [len(value)]int
		tmp = value
		for key, value := range tmp {
			println(key, value)
		}
	}
}
