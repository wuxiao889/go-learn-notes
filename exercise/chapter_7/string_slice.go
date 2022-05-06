package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "纪红宇"

	fmt.Println("prinln:", str)
	fmt.Printf("s:%s\nv:%v\n", str, str)
	fmt.Println("strlen:", len(str))
	println("utf8.RuneCountInString(str):", utf8.RuneCountInString(str))
	fmt.Println()

	for i, v := range str {
		fmt.Printf("%d %c %U\n", i, v, v)
	}
	fmt.Println()

	str1 := str[0:6]
	for _, v := range str1 {
		fmt.Printf("%c %U %v\n", v, v, v)
	}
	fmt.Println()

	str2 := str[0:4]
	for _, v := range str2 {
		fmt.Printf("%c %U %v\n", v, v, v)
	}
	fmt.Println()

	// r := []rune(str)
	r := []int32(str)
	for i, v := range r {
		fmt.Printf("%d %c %U %v\n", i, v, v, v)
	}

}
