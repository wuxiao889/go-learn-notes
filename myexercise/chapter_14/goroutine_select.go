package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan int)
	go pro(ch, done)
	ok := true
	for ok {
		select {
		case i := <-ch:
			{
				fmt.Println(i)
			}
		case <-done:
			ok = false
		}
	}
}

func pro(ch chan int, done chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}
	done <- 1
}
