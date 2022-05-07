package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)
	arr1 := append(arr[:3], arr[4:]...)
	fmt.Println(arr1)
	fmt.Println(arr)
	arr2 := append(arr[:3], arr[6:]...)
	fmt.Println(arr2)
	fmt.Println(arr)
	arr3 := append(arr[:3], append([]int{10}, arr[3:]...)...)
	fmt.Println(arr3)
	fmt.Println(arr)
}
