package main

import "fmt"

//为练习 7.3 写一个新的版本，主函数调用一个使用序列个数作为参数的函数，
//该函数返回一个大小为序列个数的 Fibonacci 切片。

var num = 50

func main() {
	result := fi(num)
	for i, v := range result {
		fmt.Printf("the %d fi is %d\n", i, v)
	}
}

func fi(num int) []int {
	arr := make([]int, num)
	arr[0], arr[1] = 0, 1
	for i := 2; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}
