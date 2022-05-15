package main

import "fmt"

//练习 7.2：for_array.go: 写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。
func main() {
	var arr [16]int
	for i := range arr {
		arr[i] = i
	}
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(arr)
	//字符byte等同于int8,Println格式化输出%v，
	fmt.Println('a')
	fmt.Println('A')
	fmt.Printf("%c", 65)
}
