package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := range fibo() {
		fmt.Println(i)
	}
	fmt.Println(time.Since(start))
}

func fibo() chan int {
	ans := []int{1, 1, 0}
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 1
		for i := 2; i < 50; i++ {
			ans[2] = (ans[1] + ans[0]) % 1000000007
			ch <- ans[2]
			ans[0], ans[1] = ans[1], ans[2]
		}
		close(ch)
	}()
	return ch
}
