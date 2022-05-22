package main

import (
	"fmt"
	"time"
)

func main() {
	pro := producer()
	go consumer(pro)
	time.Sleep(1e9)
}

func producer() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func consumer(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
