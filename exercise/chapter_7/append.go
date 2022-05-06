package main

import "fmt"

func main() {
	arr := make([]int, 10, 20)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(arr)
	// arr = append(arr[:3], arr[4:]...)
	// fmt.Println(arr)
	// arr = append(arr[:3], arr[6:]...)
	// fmt.Println(arr)
	// arr = append(arr[:3], append([]int{10}, arr[3:]...)...)
	// fmt.Println(arr)
}
