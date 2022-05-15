package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	arr := make([]int, 10)
	for i := range arr {
		arr[i] = rand.Intn(10)
	}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println(arr)
	print(sort.SearchInts(arr, 1))
}
