package main

import "fmt"

func main() {
	ch := make(chan int)
	go pro(ch)
	ok := true
	for ok {
		i := <-ch
		fmt.Println(ok, i)
	}
}

func pro(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
}
