package main

import "fmt"

func main() {
	arr1 := []float32{1.4, 5.4, 1.0, 2.0}
	fmt.Println(sum(arr1))
}

func sum(arr1 []float32) (sum float32) {
	for _, v := range arr1 {
		sum += v
	}
	return
}
