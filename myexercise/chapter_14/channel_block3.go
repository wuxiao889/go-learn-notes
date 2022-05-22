package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go receive(ch)
	go send(ch)
	time.Sleep(10 * time.Second)
}

func send(ch chan int) {
	var i = 1
	for {
		ch <- i
		i++
	}
}

func receive(ch chan int) {
	for {
		fmt.Println(<-ch)
		time.Sleep(1 * time.Second)
	}
}
