package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(10)
	}
	fmt.Println(arr)
	fmt.Println(minSlice(arr))
	fmt.Println(maxSlice(arr))
}

func minSlice(s []int) (min int) {
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return
}

func maxSlice(s []int) (max int) {
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return
}
