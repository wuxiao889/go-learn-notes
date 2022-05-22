package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int)
	go add(ch)
	go supply(ch)
	time.Sleep(1 * time.Second)
}

func add(ch chan int) {
	var sum int
	for {
		sum += <-ch
		fmt.Println(sum)
	}
}

func supply(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
