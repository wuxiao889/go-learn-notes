package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 10)
	for i := range slice {
		slice[i] = rand.Intn(10)
	}
	fmt.Println(slice)
	factor := 10
	s, len := make([]int, factor*len(slice)), len(slice)
	copy(s, slice)
	slice = s
	fmt.Println(slice)
	slice = slice[0:len]
	fmt.Println(slice)
}
