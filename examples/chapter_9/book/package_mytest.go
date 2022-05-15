package main

import (
	"fmt"

	"./pack1"
)

// book/package_mytest.go:4:2: "./pack1/pack1" is relative, but relative import paths are not supported in module mode
func main() {
	var test1 string
	test1 = pack1.ReturnStr()
	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer from package1: %d\n", pack1.Pack1Int)
	// fmt.Printf(“Float from package1: %f\n”, pack1.pack1Float)
}
