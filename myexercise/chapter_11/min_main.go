package main

import (
	"fmt"

	"./min"
)

func main() {
	arr := []int{1, 3, 4, 5, -1}
	fmt.Println(min.Min(min.IntArray(arr)))
}
