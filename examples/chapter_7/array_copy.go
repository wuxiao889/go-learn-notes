package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4: 5}
	fmt.Println(arr)
	arr_change(arr)
	fmt.Printf("%d\n", arr)
}

//传递数组的函数声明的形参大小必须和数组一样
//因为[]int和[len]int是两种类型，数组长度是数组类型的一部分
func arr_change(arr [5]int) {
	for i := range arr {
		arr[i] *= 2
	}
	fmt.Println(arr)
}
