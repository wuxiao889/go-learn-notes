package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 5)
	go add(ch)
	go supply(ch)
	time.Sleep(1 * time.Second)
}

func add(ch chan int) {
	time.Sleep(2 * time.Second)
	for {
		fmt.Println(<-ch)
	}
}

func supply(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("send", i)
	}
}
