package main

import (
	"fmt"
	"math/rand"
)

func main() {
	des := make([]int, 3)
	for i := range des {
		des[i] = rand.Intn(5)
	}
	fmt.Println("des:", printSlice(des))

	ins := make([]int, 4)
	for i := range ins {
		ins[i] = rand.Intn(15)
	}
	fmt.Println("src:", printSlice(ins))
	fmt.Println("------------")
	result := insertSlice(des, ins, 1)
	fmt.Println("after insrt:", printSlice(result))
	fmt.Println("------------")
	result = insertSlice(des, ins, 6)
	fmt.Println("after insrt:", printSlice(result))
	fmt.Println("------------")
	result = removeSlice(result, 1, 5)
	fmt.Println(printSlice(result))
}

// func insertSlice(des []int, src []int, p int) []int {
// 	var len int

// 	if p >= cap(des) {
// 		len = p + cap(src)
// 		des = expand(des, len)
// 		copy(des[p:], src)
// 	} else {
// 		len = cap(des) + cap(src)
// 		temp := make([]int, cap(des)-p)
// 		copy(temp, des[p:])
// 		fmt.Println("temp:", printSlice(temp))

// 		des = expand(des, len)

// 		copy(des[p:], src)
// 		copy(des[p+cap(src):], temp)
// 	}

// 	println("len", len)
// 	return des[:len]
// }

// func expand(des []int, size int) []int {
// 	slice := make([]int, size)
// 	copy(slice, des)
// 	des = slice
// 	fmt.Println("after expand:", printSlice(des))
// 	return des
// }

func insertSlice(des, ins []int, index int) []int {
	var result []int
	if index > cap(des) {
		result = make([]int, index+cap(ins))
		copy(result, des)
		copy(result[index:], ins)
	} else {
		result = make([]int, cap(des)+cap(ins))
		at := copy(result, des[:index])
		at += copy(result[at:], ins)
		copy(result[at:], des[index:])
	}
	return result
}
func printSlice(slice []int) string {
	return fmt.Sprintf("%v\tcap: %d\tlen: %d", slice, cap(slice), len(slice))
}

func removeSlice(slice []int, start, end int) []int {
	result := make([]int, cap(slice)-(end-start+1))
	at := copy(result, slice[:start])
	at += copy(result[at:], slice[end+1:])
	return result
}
