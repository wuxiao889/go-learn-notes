package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan int)
	go producer(ch)
	go consumer(ch, done)
	<-done
}

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i * 10
	}
	close(ch)
}

func consumer(ch <-chan int, done chan int) {
	for i := range ch {
		fmt.Println(i)
	}
	done <- 1
}
