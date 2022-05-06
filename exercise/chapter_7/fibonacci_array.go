package main

import "fmt"

//我们看到了一个递归计算 Fibonacci 数值的方法。
//是通过数组我们可以更快的计算出 Fibonacci 数。完成该方法并打印出前 50 个 Fibonacci 数字。
func main() {
	arr := [49]int{0, 1}
	for i := 2; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	fmt.Println(arr)
}
