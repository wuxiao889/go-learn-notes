package main

import "fmt"

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func main() {
	var arr = [5]int{0, 1, 2, 3, 4}
	sum(arr[:])
	arr_change(arr[:])
	fmt.Println(arr) //[0 2 4 6 8]
}

func arr_change(arr []int) {
	for i := range arr {
		arr[i] *= 2
	}
}
