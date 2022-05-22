package main

import "fmt"

func main() {
	ch := make(chan int)
	go sum(1, 2, ch)
	fmt.Println(<-ch)
}

func sum(a, b int, ch chan int) {
	ch <- a + b
}
