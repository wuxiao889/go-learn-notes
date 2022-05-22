package main

import "fmt"

func main() {
	ch := make(chan int)
	go pro(ch)
	ok := true
	var i int
	for ok {
		if i, ok = <-ch; ok {
			fmt.Println(ok, i)
		}
	}
}

func pro(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	close(ch)
}
