package main

import (
	"fmt"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}
func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}
func main() {
	b()
}

func foo() int {
	var i int

	defer func() {
		i++
	}()

	return 1
}
