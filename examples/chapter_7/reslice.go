package main

import (
	"fmt"
)

func main() {
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//切片的容量=相关数组长度-切片第一个元素在相关数组的索引
	var a = ar[4:7] // reference to subarray {5,6} - len(a) is 2 and cap(a) is 5
	fmt.Println(cap(a), len(a))
	a = a[1:3]
	fmt.Println(cap(a), len(a))
	a = a[0:6]
	fmt.Println(cap(a), len(a))
	var s []byte
	s = append(s, "hello"...)
	s = append(s, []byte("world")...)
	str := append([]byte("hello,"), "world"...)
	str = append(str)
	copy(s, []byte("hello"))
	copy(s, "hello")
}
